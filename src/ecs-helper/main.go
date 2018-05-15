package main

import (
    "os"
    "log"
    "ecs-helper/cmd"
)

const (
    ExitCodeOk = 0
    ExitCodeEcsComposeError
    ExitCodeCommandError
)

func main () {
    // docker-compose.yml
    if err := cmd.RootCmd.Execute(); err != nil {
        log.Fatal(err)
        os.Exit(ExitCodeCommandError)
    }

}
