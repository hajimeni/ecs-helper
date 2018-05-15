package cmd

import (
    "github.com/spf13/cobra"
    "ecs-helper/service"
    "ecs-helper/config"
    "errors"
)

var (
    stopCmdOptions = &config.StopCmdOptions{
        Timeout: 10,
        EcsHelperConfigPath: "ecs-config.yml",
    }
)

func init() {
    RootCmd.AddCommand(stopCmd)
    stopCmd.PersistentFlags().StringSliceVarP(&stopCmdOptions.ComposeFiles, "file", "f", []string{"docker-compose.yml", "docker-compose.override.yml"}, "")
    stopCmd.PersistentFlags().IntVar(&stopCmdOptions.Timeout, "timeout", 10, "")
    stopCmd.PersistentFlags().StringVarP(&stopCmdOptions.EcsHelperConfigPath, "config", "c","ecs-config.yml", "")
}

var stopCmd = &cobra.Command{
    Use: "stop",
    Short: "",
    Long: "",
    RunE: func(cmd *cobra.Command, args []string) error {
        code, err := service.Stop(stopCmdOptions, cmdOpts)
        if err != nil {
            return err
        }
        if code != 0 {
            return errors.New("Command error")
        }
        return nil
    },
}
