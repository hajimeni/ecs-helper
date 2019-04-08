package cmd

import (
    "github.com/spf13/cobra"
    cmdAlb "ecs-helper/cmd/alb"
)

func NewCmdAlb() *cobra.Command {
    cmd := &cobra.Command{
        Use: "alb",
        Short: "",
        Run: runHelp,
    }

    cmd.AddCommand(cmdAlb.NewCmdAlbTglist())
    return cmd
}
