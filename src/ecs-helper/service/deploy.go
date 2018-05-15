package service

import (
    "ecs-helper/config"
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "path/filepath"
    "fmt"
    "log"
    "os"
    "ecs-helper/wrapper"
    "strconv"
)

func Deploy(cmdOpts *config.DeployCmdOptions, globalOpts *config.CmdOptions) (int, error) {
    tmpPath, _ := ioutil.TempDir(".", ".tmp-")
    log.Println(strconv.Itoa(cmdOpts.Timeout))
    defer os.RemoveAll(tmpPath)

    // load from ecs-helper
    helperConfig, err := config.LoadEcsHelperConfig(cmdOpts.EcsHelperConfigPath)
    if err != nil {
        fmt.Println(err)
        return -1, err
    }

    // docker-compose.yml path
    ecsParamPath := generateEcsParamYaml(helperConfig, tmpPath)

    args := []string{
        "compose",
    }
    for _, f := range cmdOpts.ComposeFiles {
        args = append(args, "-f", f)
    }
    if globalOpts.Verbose {
        args = append(args, "--debug")
    }

    if helperConfig.ProjectName != "" {
        args = append(args, "-p", helperConfig.ProjectName)
    }
    if helperConfig.ServiceDefinition.Region != "" {
        args = append(args, "-r", helperConfig.ServiceDefinition.Region)
    }
    args = append(args, "--ecs-params", ecsParamPath)
    if helperConfig.ServiceDefinition.ClusterName != "" {
        args = append(args, "-c", helperConfig.ServiceDefinition.ClusterName)
    }

    // append service
    args = append(args, "service", "up")
    args = append(args, "--timeout", strconv.Itoa(cmdOpts.Timeout))
    args = append(args, "--deployment-max-percent", strconv.Itoa(helperConfig.ServiceDefinition.DeploymentConfiguration.MaximumPercent))
    args = append(args, "--deployment-min-healthy-percent", strconv.Itoa(helperConfig.ServiceDefinition.DeploymentConfiguration.MinimumHealthyPercent))
    for _, lb := range helperConfig.ServiceDefinition.LoadBalancers {
        if lb.LoadBalancerName != nil {
            args = append(args, "--load-balancer-name", *lb.LoadBalancerName)
        }
        if lb.TargetGroupArn != nil {
            args = append(args, "--target-group-arn", *lb.TargetGroupArn)
        }
        args = append(args, "--container-name", *lb.ContainerName)
        args = append(args, "--container-port", strconv.Itoa(*lb.ContainerPort))
    }
    //args = append(args, "--create-log-groups")
    args = append(args, "--force-deployment")

    return wrapper.ExecuteEcsCli(args)
}

func generateEcsParamYaml(helperConfig *config.EcsHeplerConfig, dir string) string {
    configBytes, _ := yaml.Marshal(helperConfig)
    ecsParams := EcsParamsWrapper{}
    yaml.Unmarshal(configBytes, &ecsParams)

    ecsParamsBytes, _ := yaml.Marshal(ecsParams)
    ecsParamsFilePath := filepath.Join(dir, "ecs-params.yml")

    ioutil.WriteFile(ecsParamsFilePath, ecsParamsBytes, 0644)

    log.Printf("ecs-params.yml output\n---\n%s\n", string(ecsParamsBytes))

    return ecsParamsFilePath
}
