package cmd

import (
    "ecs-helper/option"
    "ecs-helper/service"
    "github.com/spf13/cobra"
)

func NewCmdEcrList() *cobra.Command {
    o := option.EcrListCmdOptions{}
    cmd := &cobra.Command{
        Use: "list",
        Short: "",
        RunE: func(c *cobra.Command, args []string) error {
            err := service.ListEcrRepositories(o)
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
