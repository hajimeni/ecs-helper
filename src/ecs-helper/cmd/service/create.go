package cmd

import (
    "github.com/spf13/cobra"
    "ecs-helper/option"
    "ecs-helper/service"
)

func NewCmdServiceCreate(srvOpt *option.ServiceCmdOptions) *cobra.Command {
    o := &option.ServiceCreateCmdOptions{ srvOpt }

    cmd := &cobra.Command{
        Use: "create",
        Short: "",
        RunE: func(c *cobra.Command, args []string) error {
            _, err := service.ServiceCreate(*o)
            if err != nil {
                return err
            }
            return nil
        },
    }

    return cmd
}
