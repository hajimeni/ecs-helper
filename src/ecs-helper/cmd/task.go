package cmd

import (
    "ecs-helper/option"
    "github.com/spf13/cobra"
    taskService "ecs-helper/cmd/task"
)

func NewCmdTask() *cobra.Command {
    o := &option.TaskCmdOptions{}

    cmd := &cobra.Command{
        Use: "task",
        Short: "",
        Run: runHelp,
    }

    cmd.PersistentFlags().BoolVar(&o.Verbose, "verbose", false, "")
    cmd.PersistentFlags().StringSliceVarP(&o.ComposeFiles, "file", "f", []string{"docker-compose.yml", "docker-compose.override.yml"}, "")
    cmd.PersistentFlags().IntVar(&o.Timeout, "timeout", 10, "")
    cmd.PersistentFlags().StringVarP(&o.EcsHelperConfigPath, "config", "c","ecs-config.yml", "")

    cmd.AddCommand(taskService.NewCmdTaskCreate(o))

    return cmd
}
