package service

import (
    "ecs-helper/config"
    "io/ioutil"
    "fmt"
    "log"
    "os"
    "ecs-helper/wrapper"
    "strconv"
)

func Stop(cmdOpts *config.StopCmdOptions, globalOpts *config.CmdOptions) (int, error) {
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
    args = append(args, "service", "stop")
    args = append(args, "--timeout", strconv.Itoa(cmdOpts.Timeout))
    //args = append(args, "--create-log-groups")

    return wrapper.ExecuteEcsCli(args)
}
