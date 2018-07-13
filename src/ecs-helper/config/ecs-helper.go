package config

import (
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "log"
)

type EcsHeplerConfig struct {
    Version int
    ProjectName string `yaml:"project_name"`
    LaunchType string `yaml:"launch_type"`

    TaskDefinition TaskDefinition `yaml:"task_definition"`
    ServiceDefinition ServiceDefinition `yaml:"service_definition"`
    RunParams RunParams `yaml:"run_params,omitempty"`
}

type LoadbalancerConfig struct {
    LoadBalancerName *string `yaml:"load_balancer_name,omitempty"`
    TargetGroupArn *string `yaml:"target_group_arn,omitempty"`
    ContainerName *string `yaml:"container_name,omitempty"`
    ContainerPort *int `yaml:"container_port,omitempty"`
    DnsName *string `yaml:"dns_name,omitempty"`
    DnsZoneName *string `ymal:"dns_zone_name,omitempty"`
}

type DeploymentConfiguration struct {
    MaximumPercent int `yaml:"maximum_percent"`
    MinimumHealthyPercent int `yaml:"minimum_healthy_percent"`
}

type ServiceDefinition struct {
    Region string
    ClusterName string `yaml:"cluster_name"`
    LoadBalancers []LoadbalancerConfig `yaml:"load_balancers,omitempty"`
    DeploymentConfiguration DeploymentConfiguration `yaml:"deployment_configuration"`

    HealthCheckGracePeriod int `yaml:"health_check_grace_period"`
}

func (e *EcsHeplerConfig) ToString() string {
    out, _ := yaml.Marshal(e)
    return string(out)
}

func LoadEcsHelperConfig(path string) (*EcsHeplerConfig, error) {
    fileBytes, err := ioutil.ReadFile(path)
    if err != nil {
        return nil, err
    }
    out := EcsHeplerConfig{}
    yaml.Unmarshal(fileBytes, &out)
    log.Println("ecs helper config\n---", string(fileBytes))
    return &out, nil
}
