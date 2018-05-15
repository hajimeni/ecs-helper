package cmd

import (
    "github.com/spf13/cobra"
    "ecs-helper/service"
    "ecs-helper/config"
    "errors"
)

var (
    eraseCmdOptions = &config.EraseCmdOptions{
        Timeout: 10,
        EcsHelperConfigPath: "ecs-config.yml",
    }
)

func init() {
    RootCmd.AddCommand(eraseCmd)
    eraseCmd.PersistentFlags().StringSliceVarP(&eraseCmdOptions.ComposeFiles, "file", "f", []string{"docker-compose.yml", "docker-compose.override.yml"}, "")
    eraseCmd.PersistentFlags().IntVar(&eraseCmdOptions.Timeout, "timeout", 10, "")
    eraseCmd.PersistentFlags().StringVarP(&eraseCmdOptions.EcsHelperConfigPath, "config", "c","ecs-config.yml", "")
}

var eraseCmd = &cobra.Command{
    Use: "erase",
    Short: "",
    Long: "",
    RunE: func(cmd *cobra.Command, args []string) error {
        code, err := service.Erase(eraseCmdOptions, cmdOpts)
        if err != nil {
            return err
        }
        if code != 0 {
            return errors.New("Command error")
        }
        return nil
    },
}
