package cmd

import (
    "github.com/spf13/cobra"
    "ecs-helper/option"
    "ecs-helper/service"
)

func NewCmdAlbTglist() *cobra.Command {
    o := option.AlbTglistCmdOptions{}
    cmd := &cobra.Command{
        Use: "tglist",
        Short: "",
        RunE: func(c *cobra.Command, args []string) error {
            err := service.AlbTglist(o)
            if err != nil {
                return err
            }
            return nil
        },
    }
    cmd.PersistentFlags().StringVarP(&o.Name, "name", "n", "", "")
    cmd.PersistentFlags().StringVarP(&o.Region, "region", "r", "", "")

    return cmd
}
