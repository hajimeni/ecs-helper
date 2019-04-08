package cmd

import (
    "ecs-helper/option"
    "github.com/spf13/cobra"
    cmdService "ecs-helper/cmd/service"
)

func NewCmdService() *cobra.Command {
    o := &option.ServiceCmdOptions{}

    cmd := &cobra.Command{
        Use: "service",
        Short: "",
        Run: runHelp,
    }

    cmd.PersistentFlags().BoolVar(&o.Verbose, "verbose", false, "")
    cmd.PersistentFlags().StringSliceVarP(&o.ComposeFiles, "file", "f", []string{"docker-compose.yml", "docker-compose.override.yml"}, "")
    cmd.PersistentFlags().IntVar(&o.Timeout, "timeout", 10, "")
    cmd.PersistentFlags().StringVarP(&o.EcsHelperConfigPath, "config", "c","ecs-config.yml", "")

    cmd.AddCommand(cmdService.NewCmdServiceCreate(o))
    cmd.AddCommand(cmdService.NewCmdServiceDeploy(o))
    cmd.AddCommand(cmdService.NewCmdServiceScale(o))
    cmd.AddCommand(cmdService.NewCmdServiceDelete(o))
    cmd.AddCommand(cmdService.NewCmdServiceList(o))
    return cmd
}
