package config

import (
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "log"
)

type ServiceConfig struct {
    MonitoringType string `yaml:"monitoring_type"`
    LoggingType string `yaml:"logging_type"`
    Essential bool
}

type TaskDefinition struct {
    EcsNetworkMode string `yaml:"ecs_network_mode"`
    TaskRoleArn *string `yaml:"task_role_arn,omitempty"`
    TaskExecutionRole *string `yaml:"task_execution_role,omitempty"`
    Services map[string]ServiceConfig
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

type EcsHeplerConfig struct {
    Version int
    TemplatePath string `yaml:"template_path"`
    ProjectName string `yaml:"project_name"`
    LaunchType string `yaml:"launch_type"`

    TaskDefinition TaskDefinition `yaml:"task_definition"`
    ServiceDefinition ServiceDefinition `yaml:"service_definition"`

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
    out := EcsHeplerConfig{
        TemplatePath: "templage",
        LaunchType: "EC2",
        ServiceDefinition: ServiceDefinition{
            DeploymentConfiguration: DeploymentConfiguration{
                MaximumPercent: 200,
                MinimumHealthyPercent: 100,
            },
            HealthCheckGracePeriod: 10,
        },
        TaskDefinition: TaskDefinition{
            EcsNetworkMode: "bridge",
        },
    }
    yaml.Unmarshal(fileBytes, &out)
    log.Println("ecs helper config\n---", string(fileBytes))
    return &out, nil
}
