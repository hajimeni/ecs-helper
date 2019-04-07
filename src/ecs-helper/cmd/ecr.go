package cmd

import (
    "github.com/spf13/cobra"
    cmdEcr "ecs-helper/cmd/ecr"

)


func NewCmdEcr() *cobra.Command {
    cmd := &cobra.Command{
        Use: "ecr",
        Short: "",
        Run: runHelp,
    }

    cmd.AddCommand(cmdEcr.NewCmdEcrList())
    cmd.AddCommand(cmdEcr.NewCmdEcrImages())
    cmd.AddCommand(cmdEcr.NewCmdEcrCreate())

    return cmd
}
