package cmd

import (
    "github.com/spf13/cobra"
    "ecs-helper/option"
    "ecs-helper/service"
)

func NewCmdServiceList(srvCmd *option.ServiceCmdOptions) *cobra.Command {
    o := &option.ServiceListCmdOptions{ServiceCmdOptions: srvCmd}
    cmd := &cobra.Command{
        Use: "list",
        Short: "",
        RunE: func(c *cobra.Command, args []string) error {
            _, err := service.ServiceList(*o)
            if err != nil {
                return err
            }
            return nil
        },
    }

    return cmd
}
