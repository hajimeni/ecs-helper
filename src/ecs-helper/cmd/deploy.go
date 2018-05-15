package cmd

import (
    "github.com/spf13/cobra"
    "ecs-helper/service"
    "ecs-helper/config"
    "errors"
)

var (
    deployCmdOptions = &config.DeployCmdOptions{
        ComposeFiles: []string{"docker-compose.yml", "docker-compose.override.yml"},
        Timeout: 10,
        EcsHelperConfigPath: "ecs-config.yml",
    }
)

func init() {
    RootCmd.AddCommand(deployCmd)

    deployCmd.PersistentFlags().StringSliceVarP(&deployCmdOptions.ComposeFiles, "file", "f", []string{"docker-compose.yml", "docker-compose.override.yml"}, "")
    deployCmd.PersistentFlags().IntVar(&deployCmdOptions.Timeout, "timeout", 10, "")
    deployCmd.PersistentFlags().StringVarP(&deployCmdOptions.EcsHelperConfigPath, "config", "c","ecs-config.yml", "")
}

var deployCmd = &cobra.Command{
    Use: "deploy",
    Short: "",
    Long: "",
    RunE: func(cmd *cobra.Command, args []string) error {
        code, err := service.Deploy(deployCmdOptions, cmdOpts)
        if err != nil {
            return err
        }
        if code != 0 {
            return errors.New("Command error")
        }
        return nil
    },
}
