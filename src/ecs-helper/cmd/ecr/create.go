package cmd

import (
    "ecs-helper/option"
    "ecs-helper/service"
    "github.com/spf13/cobra"
)

func NewCmdEcrCreate() *cobra.Command {
    o := option.EcrCreteCmdOptions{}
    cmd := &cobra.Command{
        Use: "create",
        Short: "",
        RunE: func(c *cobra.Command, args []string) error {
            err := service.CreateEcrRepository(o)
            if err != nil {
                return err
            }
            return nil
        },
    }
    cmd.PersistentFlags().StringVarP(&o.Name, "name", "n", "", "")
    cmd.PersistentFlags().StringVarP(&o.Region, "region", "r", "", "")

    return cmd
}
