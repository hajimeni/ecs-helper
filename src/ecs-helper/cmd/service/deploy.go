package cmd

import (
    "github.com/spf13/cobra"
    "ecs-helper/option"
    "ecs-helper/service"
)

func NewCmdServiceDeploy(srvOpt *option.ServiceCmdOptions) *cobra.Command {
    o := &option.ServiceDeployCmdOptions{ServiceCmdOptions: srvOpt}
    cmd := &cobra.Command{
        Use: "deploy",
        Short: "",
        RunE: func(c *cobra.Command, args []string) error {
            _, err := service.ServiceDeploy(*o)
            if err != nil {
                return err
            }
            return nil
        },
    }
    cmd.Flags().BoolVar(&o.ForceDeployment, "force-deployment", false, "")

    return cmd
}
