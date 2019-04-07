package cmd

import (
    "ecs-helper/option"
    "ecs-helper/service"
    "github.com/spf13/cobra"
)

func NewCmdEcrImages() *cobra.Command {
    o := option.EcrImagesCmdOptions{}
    cmd := &cobra.Command{
        Use: "images",
        Short: "",
        RunE: func(c *cobra.Command, args []string) error {
            err := service.ListEcrImages(o)
            if err != nil {
                return err
            }
            return nil
        },
    }

    cmd.PersistentFlags().StringVarP(&o.Name, "name", "n", "", "")
    cmd.PersistentFlags().StringVarP(&o.Region, "region", "r", "", "")
    cmd.PersistentFlags().BoolVar(&o.Tagged, "tagged", false, "")
    cmd.PersistentFlags().BoolVar(&o.UnTagged, "untagged", false, "")

    return cmd
}
