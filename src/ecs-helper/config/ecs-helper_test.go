package config

import (
    "testing"
    "os"
    "path/filepath"
)

func TestLoadEcsHelperConfig(t *testing.T) {
    wd, _ := os.Getwd()
    println(filepath.Base(wd))
    example_config := filepath.Join(wd, "/ecs-helper-example.yml")
    config, err := LoadEcsHelperConfig(example_config)
    if err != nil {
        t.Errorf("Cannot load config file: %s", wd)
    }
    println(config.ToString())
    if  len(config.TaskDefinition.Services) != 2 {
        t.Errorf("Service definition error")
    }
    if config.Version != 1 {
        t.Errorf("Version error")
    }
    // TODO struct の DeepEqual
}
