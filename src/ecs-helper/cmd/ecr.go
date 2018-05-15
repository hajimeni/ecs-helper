package cmd

import (
    "github.com/spf13/cobra"
    "ecs-helper/config"
    "ecs-helper/service"
    "os"
)

var (
    ecrListCmdOptions = &config.EcrListCmdOptions{}
    ecrImagesCmdOptions = &config.EcrImagesCmdOptions{}
    ecrCreateCmdOptions = &config.EcrCreteCmdOptions{}
)

func init() {
    ecrLsCmd.PersistentFlags().StringVarP(&ecrListCmdOptions.Name, "name", "n", "", "")
    ecrLsCmd.PersistentFlags().StringVarP(&ecrListCmdOptions.Region, "region", "r", "", "")

    ecrImagesCmd.PersistentFlags().StringVarP(&ecrImagesCmdOptions.Name, "name", "n", "", "")
    ecrImagesCmd.PersistentFlags().StringVarP(&ecrImagesCmdOptions.Region, "region", "r", "", "")
    ecrImagesCmd.PersistentFlags().BoolVar(&ecrImagesCmdOptions.Tagged, "tagged", false, "")
    ecrImagesCmd.PersistentFlags().BoolVar(&ecrImagesCmdOptions.UnTagged, "untagged", false, "")

    ecrCreateCmd.PersistentFlags().StringVarP(&ecrCreateCmdOptions.Name, "name", "n", "", "")
    ecrCreateCmd.PersistentFlags().StringVarP(&ecrCreateCmdOptions.Region, "region", "r", "", "")

    ecrCmd.SetOutput(os.Stdout)
    ecrLsCmd.SetOutput(os.Stdout)
    ecrImagesCmd.SetOutput(os.Stdout)
    ecrCreateCmd.SetOutput(os.Stdout)

    RootCmd.AddCommand(ecrCmd)
    ecrCmd.AddCommand(ecrLsCmd)
    ecrCmd.AddCommand(ecrImagesCmd)
    ecrCmd.AddCommand(ecrCreateCmd)
}


var ecrCmd = &cobra.Command{
    Use: "ecr",
    RunE: func(cmd *cobra.Command, args []string) error {
        return cmd.Execute()
    },
}

var ecrLsCmd = &cobra.Command{
    Use: "list",
    RunE: func(cmd *cobra.Command, args []string) error {
        return service.ListEcrRepositories(*ecrListCmdOptions)
    },
}

var ecrImagesCmd = &cobra.Command{
    Use: "images",
    RunE: func(cmd *cobra.Command, args []string) error {
        return service.ListEcrImages(*ecrImagesCmdOptions)
    },
}

var ecrCreateCmd = &cobra.Command{
    Use: "create",
    RunE: func(cmd *cobra.Command, args []string) error {
        return service.CreateEcrRepository(*ecrCreateCmdOptions)
    },
}
