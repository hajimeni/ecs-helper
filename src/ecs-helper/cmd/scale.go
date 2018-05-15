package cmd

import (
    "github.com/spf13/cobra"
    "ecs-helper/service"
    "ecs-helper/config"
    "errors"
)

var (
    scaleCmdOptions = &config.ScaleCmdOptions{
        Size: 1,
        ComposeFiles: []string{"docker-compose.yml", "docker-compose.override.yml"},
        Timeout: 10,
        EcsHelperConfigPath: "ecs-config.yml",
    }
)

func init() {
    RootCmd.AddCommand(scaleCmd)

    scaleCmd.PersistentFlags().IntVar(&scaleCmdOptions.Size, "size", 10, "")
    scaleCmd.PersistentFlags().StringSliceVarP(&scaleCmdOptions.ComposeFiles, "file", "f", []string{"docker-compose.yml", "docker-compose.override.yml"}, "")
    scaleCmd.PersistentFlags().IntVar(&scaleCmdOptions.Timeout, "timeout", 10, "")
    scaleCmd.PersistentFlags().StringVarP(&scaleCmdOptions.EcsHelperConfigPath, "config", "c","ecs-config.yml", "")
}

var scaleCmd = &cobra.Command{
    Use: "scale",
    Short: "",
    Long: "",
    RunE: func(cmd *cobra.Command, args []string) error {
        code, err := service.Scale(scaleCmdOptions, cmdOpts)
        if err != nil {
            return err
        }
        if code != 0 {
            return errors.New("Command error")
        }
        return nil
    },
}
