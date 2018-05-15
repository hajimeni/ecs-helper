package cmd

import (
    "github.com/spf13/cobra"
    "ecs-helper/config"
    "ecs-helper/service"
    "os"
)

var (
    roleListCmdOptions = &config.RoleListCmdOptions{}
)

func init() {
    RootCmd.AddCommand(roleCmd)
    roleCmd.AddCommand(roleLsCmd)

    roleLsCmd.PersistentFlags().StringVarP(&roleListCmdOptions.Name, "name", "n", "", "")

    roleCmd.SetOutput(os.Stdout)
    roleLsCmd.SetOutput(os.Stdout)
}


var roleCmd = &cobra.Command{
    Use: "role",
    RunE: func(cmd *cobra.Command, args []string) error {
        return cmd.Execute()
    },
}

var roleLsCmd = &cobra.Command{
    Use: "list",
    RunE: func(cmd *cobra.Command, args []string) error {
        return service.ListTaskRoles(*roleListCmdOptions)
    },
}
