package cmd

import (
    "github.com/spf13/cobra"
    "ecs-helper/option"
    "ecs-helper/service"
)

func NewCmdServiceScale(srvCmd *option.ServiceCmdOptions) *cobra.Command {
    o := &option.ServiceScaleCmdOptions{ServiceCmdOptions: srvCmd}
    cmd := &cobra.Command{
        Use: "scale",
        Short: "",
        RunE: func(c *cobra.Command, args []string) error {
            _, err := service.ServiceScale(*o)
            if err != nil {
                return err
            }
            return nil
        },
    }
    cmd.Flags().IntVar(&o.Count, "count", -1, "")

    return cmd
}
