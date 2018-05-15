package cmd

import (
    "github.com/spf13/cobra"
    "ecs-helper/wrapper"
    "ecs-helper/config"
)

var (
    cmdOpts = &config.CmdOptions{}
)

var RootCmd = &cobra.Command{
     Use: "ecs-helper",
     Short: "",
     Long: "",
     Run: func(cmd *cobra.Command, args []string) {},
}

func init() {
    cobra.OnInitialize()

    RootCmd.PersistentFlags().BoolVar(&cmdOpts.Verbose, "verbose", false,"")
    RootCmd.PersistentFlags().BoolVar(&cmdOpts.Verbose, "debug", false,"")

    RootCmd.AddCommand(ecsCliCmd)
}

var ecsCliCmd = &cobra.Command{
    Use: "cli",
    DisableFlagParsing: true,
    Short: "",
    Long: "",
    Run: func(cmd *cobra.Command, args []string) {
        wrapper.ExecuteEcsCli(args)
    },
}
