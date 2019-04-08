package service

import (
    "ecs-helper/config"
    "ecs-helper/option"
    "ecs-helper/wrapper"
    "io/ioutil"
    "os"
    "strings"
)

func TaskCreate(cmdOption option.TaskCreateCmdOptions) (int, error) {
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

func taskCreateToEcsCliArgs(cfg *config.EcsConfig, composeFiles []string, ecsParamPath string, verbose bool) []string {
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


    args = append(args, "create")

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
