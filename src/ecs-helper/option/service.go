package option

type ServiceCmdOptions struct {
    Timeout int
    Verbose bool
    ComposeFiles []string
    EcsHelperConfigPath string
}

type ServiceCreateCmdOptions struct {
    *ServiceCmdOptions
}

type ServiceDeployCmdOptions struct {
    *ServiceCmdOptions
    ForceDeployment bool
}

type ServiceScaleCmdOptions struct {
    *ServiceCmdOptions
    Count int
}
type ServiceDeleteCmdOptions struct {
    *ServiceCmdOptions
}

type ServiceListCmdOptions struct {
    *ServiceCmdOptions
}

