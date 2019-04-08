package cmd

import (
    "ecs-helper/option"
    "ecs-helper/service"
    "github.com/spf13/cobra"
)

func NewCmdRoleList() *cobra.Command {
    o := option.RoleListCmdOptions{}
    cmd := &cobra.Command{
        Use: "list",
        Short: "",
        RunE: func(c *cobra.Command, args []string) error {
            err := service.ListTaskRoles(o)
            if err != nil {
                return err
            }
            return nil
        },
    }
    cmd.PersistentFlags().StringVarP(&o.Name, "name", "n", "", "")

    return cmd
}
