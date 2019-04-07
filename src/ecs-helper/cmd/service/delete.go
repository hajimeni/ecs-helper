package cmd

import (
    "github.com/spf13/cobra"
    "ecs-helper/option"
    "ecs-helper/service"
)

func NewCmdServiceDelete(srvOpt *option.ServiceCmdOptions) *cobra.Command {
    o := &option.ServiceDeleteCmdOptions{ srvOpt }
    cmd := &cobra.Command{
        Use: "delete",
        Short: "",
        RunE: func(c *cobra.Command, args []string) error {
            _, err := service.ServiceDelete(*o)
            if err != nil {
                return err
            }
            return nil
        },
    }

    return cmd
}
