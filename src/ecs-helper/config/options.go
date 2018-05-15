package config

type CmdOptions struct {
    Verbose bool
}


type DeployCmdOptions struct {
    ComposeFiles []string
    Timeout int
    EcsHelperConfigPath string
}

type StopCmdOptions struct {
    ComposeFiles []string
    Timeout int
    EcsHelperConfigPath string
}

type EraseCmdOptions struct {
    ComposeFiles []string
    Timeout int
    EcsHelperConfigPath string
}

type ScaleCmdOptions struct {
    Size int
    ComposeFiles []string
    Timeout int
    EcsHelperConfigPath string
}


type GetLoginCmdOptions struct {
    Region string
    RegistryId string
}


type AlbListCmdOptions struct {
    Name string
    Region string
}


type RoleListCmdOptions struct {
    Name string
}


type EcrListCmdOptions struct {
    Name string
    Region string
}

type EcrImagesCmdOptions struct {
    Name string
    Region string
    Tagged bool
    UnTagged bool
}

type EcrCreteCmdOptions struct {
    Name string
    Region string
}
