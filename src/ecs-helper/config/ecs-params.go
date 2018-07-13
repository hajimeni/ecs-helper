package config

type EcsParamsWrapper struct {
    Version int
    TaskDefinition TaskDefinition `yaml:"task_definition"`
    RunParams RunParams `yaml:"run_params,omitempty"`
}

type RunParams struct {
    NetworkConfiguration NetworkConfiguration `yaml:"network_configuration"`
}

type NetworkConfiguration struct {
    AwsVpcConfiguration AwsVpcConfiguration `yaml:"awsvpc_configuration"`
}

type AwsVpcConfiguration struct {
    Subnets []string
    SecurityGroups []string `yaml:"security_groups"`
    AssignPublicIp string `yaml:"assign_public_ip"`
}

type TaskDefinition struct {
    EcsNetworkMode string `yaml:"ecs_network_mode"`
    TaskRoleArn *string `yaml:"task_role_arn,omitempty"`
    TaskExecutionRole *string `yaml:"task_execution_role,omitempty"`
    TaskSize TaskSize `yaml:"task_size,omitempty"`
    Services map[string]ServiceConfig
}

type TaskSize struct {
    CpuLimit string `yaml:"cpu_limit,omitempty"`
    MemLimit string `yaml:"mem_limit,omitempty"`
}

type ServiceConfig struct {
    Essential bool
    CpuShares int `yaml:"cpu_shares,omitempty"`
    MemLimit int `yaml:"mem_limit,omitempty"`
    MemReservation int `yaml:"mem_reservation,omitempty"`
}

