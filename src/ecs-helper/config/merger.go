package config


import (
    "github.com/imdario/mergo"
    "io/ioutil"
    "fmt"
    "gopkg.in/yaml.v2"
)

func MergeConfig() (map[interface{}]interface{}, error) {
    dockerComposeYaml, err := ioutil.ReadFile("docker-compose.yml")
    dockerComposeOverrideYaml, err := ioutil.ReadFile("docker-compose.override.yml")
    composeStruct := make(map[interface{}]interface{})
    yaml.Unmarshal(dockerComposeYaml, &composeStruct)
    overrideStruct := make(map[interface{}]interface{})
    yaml.Unmarshal(dockerComposeOverrideYaml, &overrideStruct)

    mergo.Merge(&overrideStruct, composeStruct)

    if err != nil {
        fmt.Println(err)
        return nil, err
    }

    return overrideStruct, nil
}
