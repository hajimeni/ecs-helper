package service

import (
    "ecs-helper/option"
    "ecs-helper/config"
    "ecs-helper/wrapper"
    "gopkg.in/yaml.v2"
    "log"
    "path/filepath"
    "strconv"
    "os"
    "io/ioutil"
    "errors"
    "strings"
)

func ServiceCreate(cmdOption option.ServiceCreateCmdOptions) (int, error) {
    // optionで指定されたファイルのload
    helperConfig, err := config.LoadEcsHelperConfig(cmdOption.EcsHelperConfigPath); if err != nil {
        return -1, err
    }

    tmpPath, _ := ioutil.TempDir(".", ".tmp-")
    defer os.RemoveAll(tmpPath)
    ecsParamPath := generateEcsParamFile(helperConfig.EcsParams, tmpPath)


    // ファイルのパラメータをコマンドラインオプションに変換
    args := serviceCreateToEcsCliArgs(helperConfig, cmdOption.ComposeFiles, ecsParamPath, cmdOption.Verbose)

    // ecs-cliの実行
    return wrapper.ExecuteEcsCli(args)
}

func serviceCreateToEcsCliArgs(cfg *config.EcsConfig, composeFiles []string, ecsParamPath string, verbose bool) []string {
    args := []string{ "compose", }
    if verbose {
        args = append(args, "--verbose")
    }
    args = append(args, "--region",cfg.Region)
    args = append(args, "--project-name",cfg.ProjectName)
    args = append(args, "--ecs-params", ecsParamPath)
    for _, f := range composeFiles {
        args = append(args, "--file", f)
    }

    definition := cfg.ServiceDefinition
    args = append(args, "--cluster",definition.Cluster)


    args = append(args, "service", "create")
    if definition.DeploymentConfiguration != nil {
        args = append(args, "--deployment-min-healthy-percent", strconv.Itoa(definition.DeploymentConfiguration.MinimumHealthyPercent))
        args = append(args, "--deployment-max-percent", strconv.Itoa(definition.DeploymentConfiguration.MaximumPercent))
    }

    for _, lb := range definition.LoadBalancers {
        if lb.LoadBalancerName != nil {
            args = append(args, "--load-balancer-name", *lb.LoadBalancerName)
        }
        if lb.TargetGroupArn != nil {
            args = append(args, "--target-group-arn", *lb.TargetGroupArn)
        }
        args = append(args, "--container-name", lb.ContainerName)
        args = append(args, "--container-port", strconv.Itoa(lb.ContainerPort))
    }

    tagArgs := []string{}
    for _, tag := range definition.Tags {
        if tag.Key != "" && tag.Value != "" {
            tagArgs = append(tagArgs, tag.Key + "=" + tag.Value)
        }
    }
    if len(tagArgs) > 0 {
        args = append(args, "--tags", strings.Join(tagArgs, ","))
    }

    return args
}

func ServiceDeploy(cmdOption option.ServiceDeployCmdOptions) (int, error) {

    // optionで指定されたファイルのload
    helperConfig, err := config.LoadEcsHelperConfig(cmdOption.EcsHelperConfigPath); if err != nil {
        return -1, err
    }

    tmpPath, _ := ioutil.TempDir(".", ".tmp-")
    defer os.RemoveAll(tmpPath)
    ecsParamPath := generateEcsParamFile(helperConfig.EcsParams, tmpPath)


    // ファイルのパラメータをコマンドラインオプションに変換
    args := serviceDeployToEcsCliArgs(helperConfig, cmdOption.ComposeFiles, ecsParamPath, cmdOption.Timeout, cmdOption.ForceDeployment)

    // ecs-cliの実行
    return wrapper.ExecuteEcsCli(args)
}

func serviceDeployToEcsCliArgs(cfg *config.EcsConfig, composeFiles []string, ecsParamPath string, timeout int, forceDeployment bool) []string {
    args := []string{ "compose", }
    args = append(args, "--region",cfg.Region)
    args = append(args, "--project-name",cfg.ProjectName)
    args = append(args, "--ecs-params", ecsParamPath)
    for _, f := range composeFiles {
        args = append(args, "--file", f)
    }

    definition := cfg.ServiceDefinition
    args = append(args, "--cluster",definition.Cluster)

    args = append(args, "service", "up")
    if forceDeployment == true {
        args = append(args, "--force-deployment")
    }

    if definition.DeploymentConfiguration != nil {
        args = append(args, "--deployment-min-healthy-percent", strconv.Itoa(definition.DeploymentConfiguration.MinimumHealthyPercent))
        args = append(args, "--deployment-max-percent", strconv.Itoa(definition.DeploymentConfiguration.MaximumPercent))
    }

    for _, lb := range definition.LoadBalancers {
        if lb.LoadBalancerName != nil {
            args = append(args, "--load-balancer-name", *lb.LoadBalancerName)
        }
        if lb.TargetGroupArn != nil {
            args = append(args, "--target-group-arn", *lb.TargetGroupArn)
        }
        args = append(args, "--container-name", lb.ContainerName)
        args = append(args, "--container-port", strconv.Itoa(lb.ContainerPort))
    }

    tagArgs := []string{}
    for _, tag := range definition.Tags {
        if tag.Key != "" && tag.Value != "" {
            tagArgs = append(tagArgs, tag.Key + "=" + tag.Value)
        }
    }
    if len(tagArgs) > 0 {
        args = append(args, "--tags", strings.Join(tagArgs, ","))
    }

    args = append(args, "--timeout", strconv.Itoa(timeout))

    return args
}


func ServiceScale(cmdOption option.ServiceScaleCmdOptions) (int, error) {
    if cmdOption.Count < 0 {
        return -2, errors.New("count required >= 0")
    }

    // optionで指定されたファイルのload
    helperConfig, err := config.LoadEcsHelperConfig(cmdOption.EcsHelperConfigPath); if err != nil {
        return -1, err
    }

    tmpPath, _ := ioutil.TempDir(".", ".tmp-")
    defer os.RemoveAll(tmpPath)
    ecsParamPath := generateEcsParamFile(helperConfig.EcsParams, tmpPath)


    // ファイルのパラメータをコマンドラインオプションに変換
    args := serviceScaleToEcsCliArgs(helperConfig, cmdOption.ComposeFiles, ecsParamPath, cmdOption.Timeout, cmdOption.Count)

    // ecs-cliの実行
    return wrapper.ExecuteEcsCli(args)
}

func serviceScaleToEcsCliArgs(cfg *config.EcsConfig, composeFiles []string, ecsParamPath string, timeout int, count int) []string {
    args := []string{ "compose", }
    args = append(args, "--region",cfg.Region)
    args = append(args, "--project-name",cfg.ProjectName)
    args = append(args, "--ecs-params", ecsParamPath)
    for _, f := range composeFiles {
        args = append(args, "--file", f)
    }

    definition := cfg.ServiceDefinition
    args = append(args, "--cluster",definition.Cluster)

    args = append(args, "service", "scale")
    if definition.DeploymentConfiguration != nil {
        args = append(args, "--deployment-min-healthy-percent", strconv.Itoa(definition.DeploymentConfiguration.MinimumHealthyPercent))
        args = append(args, "--deployment-max-percent", strconv.Itoa(definition.DeploymentConfiguration.MaximumPercent))
    }
    args = append(args, "--timeout", strconv.Itoa(timeout))

    args = append(args, strconv.Itoa(count))

    return args
}

func ServiceList(cmdOption option.ServiceListCmdOptions) (int, error) {
    // optionで指定されたファイルのload
    helperConfig, err := config.LoadEcsHelperConfig(cmdOption.EcsHelperConfigPath); if err != nil {
        return -1, err
    }

    tmpPath, _ := ioutil.TempDir(".", ".tmp-")
    defer os.RemoveAll(tmpPath)
    ecsParamPath := generateEcsParamFile(helperConfig.EcsParams, tmpPath)


    // ファイルのパラメータをコマンドラインオプションに変換
    args := serviceListToEcsCliArgs(helperConfig, cmdOption.ComposeFiles, ecsParamPath)

    // ecs-cliの実行
    return wrapper.ExecuteEcsCli(args)
}

func serviceListToEcsCliArgs(cfg *config.EcsConfig, composeFiles []string, ecsParamPath string) []string {
    args := []string{ "compose", }
    args = append(args, "--region",cfg.Region)
    args = append(args, "--project-name",cfg.ProjectName)
    args = append(args, "--ecs-params", ecsParamPath)
    for _, f := range composeFiles {
        args = append(args, "--file", f)
    }

    definition := cfg.ServiceDefinition
    args = append(args, "--cluster",definition.Cluster)

    args = append(args, "service", "list")

    return args
}

func ServiceDelete(cmdOption option.ServiceDeleteCmdOptions) (int, error) {
    // optionで指定されたファイルのload
    helperConfig, err := config.LoadEcsHelperConfig(cmdOption.EcsHelperConfigPath); if err != nil {
        return -1, err
    }

    tmpPath, _ := ioutil.TempDir(".", ".tmp-")
    defer os.RemoveAll(tmpPath)
    ecsParamPath := generateEcsParamFile(helperConfig.EcsParams, tmpPath)


    // ファイルのパラメータをコマンドラインオプションに変換
    args := serviceDeleteToEcsCliArgs(helperConfig, cmdOption.ComposeFiles, ecsParamPath, cmdOption.Timeout)

    // ecs-cliの実行
    return wrapper.ExecuteEcsCli(args)
}

func serviceDeleteToEcsCliArgs(cfg *config.EcsConfig, composeFiles []string, ecsParamPath string, timeout int) []string {
    args := []string{ "compose", }
    args = append(args, "--region",cfg.Region)
    args = append(args, "--project-name",cfg.ProjectName)
    args = append(args, "--ecs-params", ecsParamPath)
    for _, f := range composeFiles {
        args = append(args, "--file", f)
    }

    definition := cfg.ServiceDefinition
    args = append(args, "--cluster",definition.Cluster)

    args = append(args, "service", "delete")
    args = append(args, "--timeout", strconv.Itoa(timeout))

    return args
}

func generateEcsParamFile(ecsParams map[interface{}]interface{}, dir string) string {
    ecsParamsFilePath := filepath.Join(dir, "ecs-params.yml")

    ecsParamsBytes := generateEcsParamYamlByte(ecsParams)
    log.Printf("ecs-params.yml output\n---\n%s\n", string(ecsParamsBytes))

    ioutil.WriteFile(ecsParamsFilePath, ecsParamsBytes, 0644)
    return ecsParamsFilePath
}

func generateEcsParamYamlByte(ecsParams map[interface{}]interface{}) []byte {
    ecsParamsBytes, _ := yaml.Marshal(ecsParams)
    return ecsParamsBytes
}
