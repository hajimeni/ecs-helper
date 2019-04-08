package cmd

import (
    "github.com/spf13/cobra"
    "ecs-helper/option"
    "ecs-helper/service"
)

func NewCmdTaskCreate(taskOpt *option.TaskCmdOptions) *cobra.Command {
    o := &option.TaskCreateCmdOptions{ taskOpt }

    cmd := &cobra.Command{
        Use: "create",
        Short: "",
        RunE: func(c *cobra.Command, args []string) error {
            _, err := service.TaskCreate(*o)
            if err != nil {
                return err
            }
            return nil
        },
    }

    return cmd
}
