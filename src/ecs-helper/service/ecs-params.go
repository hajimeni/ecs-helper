package service

type EcsParamsWrapper struct {
    Version int
    TaskDefinition TaskDefinition `yaml:"task_definition"`
}

type TaskDefinition struct {
    EcsNetworkMode string `yaml:"ecs_network_mode"`
    TaskRoleArn *string `yaml:"task_role_arn,omitempty"`
    TaskExecutionRole *string `yaml:"task_execution_role,omitempty"`
    Services map[string]ServiceConfig
}

type ServiceConfig struct {
    Essential bool
}

