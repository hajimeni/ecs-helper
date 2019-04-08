package option

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
