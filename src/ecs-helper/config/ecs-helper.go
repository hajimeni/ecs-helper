package config

import (
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "log"
)

type EcsConfig struct {
    Version int
    ProjectName string `yaml:"project_name"`
    Region string
    ServiceDefinition `yaml:"service_definition"`
    EcsParams map[interface{}]interface{} `yaml:"ecs_params"`
}
type ServiceDefinition struct {
    Cluster string
    LaunchType string `yaml:"launch_type"`
    TaskDefinition *string `yaml:"task_definition,omitempty"`
    LoadBalancers []LoadBalancers `yaml:"load_balancers,omitempty"`
    Role *string
    DeploymentConfiguration *DeploymentConfiguration `yaml:"deployment_configuration,omitempty"`
    PlacementConstraints []PlacementConstraints `yaml:"placement_constraints"`
    PlacementStrategy []PlacementStrategy `yaml:"placement_strategy"`
    HealthCheckGracePeriodSeconds *int `yaml:"health_check_grace_period_seconds,omitempty"`
    SchedulingStrategy *string `yaml:"scheduling_strategy,omitempty"`
    DeploymentController *DeploymentController `yaml:"deployment_controller,omitempty"`
    Tags []Tags
    EnableECSManagedTags *bool `yaml:"enable_ecs_managed_tags,omitempty"`
    PropagateTags *string `yaml:"propagate_tags,omitempty"`
}
type LoadBalancers struct {
    TargetGroupArn *string `yaml:"target_group_arn,omitempty"`
    LoadBalancerName *string `yaml:"load_balancer_name,omitempty"`
    ContainerName string `yaml:"container_name"`
    ContainerPort int `yaml:"container_port"`
}
type DeploymentConfiguration struct {
    MaximumPercent int `yaml:"maximum_percent"`
    MinimumHealthyPercent int `yaml:"minimum_healthy_percent"`
}
type PlacementConstraints struct {
    Type string
    Expression string
}
type PlacementStrategy struct {
    Type string
    Field string
}
type DeploymentController struct {
    Type string
}
type Tags struct {
    Key string
    Value string
}

func (e *EcsConfig) ToString() string {
    out, _ := yaml.Marshal(e)
    return string(out)
}

func LoadEcsHelperConfig(path string) (*EcsConfig, error) {
    fileBytes, err := ioutil.ReadFile(path)
    if err != nil {
        return nil, err
    }
    out := EcsConfig{}
    yaml.Unmarshal(fileBytes, &out)
    log.Println("ecs helper config\n---\n", string(fileBytes))
    return &out, nil
}
