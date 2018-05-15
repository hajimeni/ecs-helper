package cmd

import (
    "github.com/spf13/cobra"
    "ecs-helper/config"
    "ecs-helper/service"
    "os"
)

var (
    albListCmdOptions = &config.AlbListCmdOptions{}
)

func init() {
    RootCmd.AddCommand(albCmd)
    albCmd.AddCommand(albTgLsCmd)

    albTgLsCmd.PersistentFlags().StringVarP(&albListCmdOptions.Name, "name", "n", "", "")
    albTgLsCmd.PersistentFlags().StringVarP(&albListCmdOptions.Region, "region", "r", "", "")

    albCmd.SetOutput(os.Stdout)
    albTgLsCmd.SetOutput(os.Stdout)
}


var albCmd = &cobra.Command{
    Use: "alb",
    RunE: func(cmd *cobra.Command, args []string) error {
        return cmd.Execute()
    },
}

var albTgLsCmd = &cobra.Command{
    Use: "tglist",
    RunE: func(cmd *cobra.Command, args []string) error {
        return service.ListTargetGroups(*albListCmdOptions)
    },
}
