package cmd

import (
    "github.com/spf13/cobra"
    "ecs-helper/wrapper"
)

func NewCmdEcsCli() *cobra.Command {
    cmd := &cobra.Command{
        Use: "cli",
        DisableFlagParsing: true,
        Short: "",
        Long: "",
        RunE: func(cmd *cobra.Command, args []string) error {
            _, err := wrapper.ExecuteEcsCli(args)
            if err != nil {
                return err
            }
            return nil
        },
    }

    return cmd
}
