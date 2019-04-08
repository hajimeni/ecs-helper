package option

type TaskCmdOptions struct {
    Timeout int
    Verbose bool
    ComposeFiles []string
    EcsHelperConfigPath string
}

type TaskCreateCmdOptions struct {
    *TaskCmdOptions
}
