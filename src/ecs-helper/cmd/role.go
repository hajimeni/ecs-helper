package cmd

import (
    "github.com/spf13/cobra"
    roleAlb "ecs-helper/cmd/role"
)

func NewCmdRole() *cobra.Command {
    cmd := &cobra.Command{
        Use: "role",
        Short: "",
        Run: runHelp,
    }

    cmd.AddCommand(roleAlb.NewCmdRoleList())
    return cmd
}
