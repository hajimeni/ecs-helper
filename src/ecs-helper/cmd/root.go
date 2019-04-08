package cmd

import (
    "github.com/spf13/cobra"
    "os"
)

func NewCmdRoot() *cobra.Command {
    cmd := &cobra.Command{
        Use: "ecs-helper",
        Short: "ecs-cli simple command wrapper",
        Run: runHelp,
    }
    cobra.OnInitialize(initConfig)

    cmd.AddCommand(NewCmdService())
    cmd.AddCommand(NewCmdEcsCli())
    cmd.AddCommand(NewCmdTask())
    cmd.AddCommand(NewCmdEcr())
    cmd.AddCommand(NewCmdAlb())
    cmd.AddCommand(NewCmdGetLogin())
    cmd.AddCommand(NewCmdRole())

    return cmd
}

func Execute() {
    cmd := NewCmdRoot()
    cmd.SetOutput(os.Stdout)
    if err := cmd.Execute(); err != nil {
        cmd.SetOutput(os.Stderr)
        cmd.Println(err)
        os.Exit(1)
    }

}

func initConfig() {
    // pass
}


func runHelp(cmd *cobra.Command, args []string) {
    _ = cmd.Help()
}
