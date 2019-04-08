package config

import (
	// "ecs-helper/config"
	goyaml "gopkg.in/yaml.v2"
	"testing"
)
var c = `
version: 1
project_name: test_project
region: test_region

# like service definition
service_definition:
  cluster: test_cluster
  launch_type: FARGATE
  load_balancers:
    - target_group_arn: test_target_group_arn
      container_name: test_container_name
      container_port: 100
  deployment_configuration:
    maximum_percent: 100
    minimum_healthy_percent: 50
  placement_constraints:
    - type: hoge
      expression: "aa"
  placement_strategy:
    - type: fuga
      field: "bb"
  health_check_grace_period_seconds: 60
  scheduling_strategy: REPLICA
  deployment_controller:
    type: EC2

## equals to ecs_params
ecs_params:
  version: 1
  task_definition:
    ecs_network_mode: ip
    task_role_arn: string
    task_execution_role: string
    task_size:
      cpu_limit: string
      mem_limit: string
    pid_mode: string
    ipc_mode: string
    services:
      service_name:
        essential: boolean
        repository_credentials:
          credentials_parameter: string
        cpu_shares: integer
        mem_limit: string
        mem_reservation: string
        healthcheck:
          test: ["CMD", "curl -f http://localhost"]
          interval: string
          timeout: string
          retries: integer
          start_period: string
        secrets:
          - value_from: string
            name: string
    docker_volumes:
        - name: string
          scope: string
          autoprovision:
          driver: string
          driver_opts:
             string: string
          labels:
             string: string
  run_params:
    network_configuration:
      awsvpc_configuration:
        subnets:
          - subnet_id1
          - subnet_id2
        security_groups:
          - secgroup_id1
          - secgroup_id2
        assign_public_ip: ENABLED
    task_placement:
      strategy:
        - type: string
          field: string
      constraints:
        - type: string
          expression: string
    service_discovery:
      container_name: string
      container_port: integer
      private_dns_namespace:
        vpc: string
        id: string
        name: string
        description: string
      public_dns_namespace:
        id: string
        name: string
      service_discovery_service:
        name: string
        description: string
        dns_config:
          type: string
          ttl: integer
        healthcheck_custom_config:
          failure_threshold: integer
`

func Test_loadFromYaml(t *testing.T) {
	tt := EcsConfig{}
	err := goyaml.Unmarshal([]byte(c), &tt)
	if err != nil {
		t.Errorf("Yaml load error. error=%v", err)
	}
	if tt.Version != 1 {
		t.Errorf("Version load error")
	}
	if tt.ProjectName != "test_project" {
		t.Errorf("Projectname load error")
	}
	if tt.ServiceDefinition.Cluster != "test_cluster" {
		t.Errorf("ServiceDefinition Cluster load error")
	}
	if len(tt.ServiceDefinition.LoadBalancers) != 1 {
		t.Errorf("ServiceDefinition LoadBalancers load error")
	}
	if *tt.ServiceDefinition.LoadBalancers[0].TargetGroupArn != "test_target_group_arn" {
		t.Errorf("ServiceDefinition LoadBalancers TargetGroupArn load error")
	}
	if tt.ServiceDefinition.LoadBalancers[0].LoadBalancerName != nil {
		t.Errorf("ServiceDefinition LoadBalancers LoadBalancerName load error")
	}
	if tt.ServiceDefinition.LoadBalancers[0].ContainerName != "test_container_name" {
		t.Errorf("ServiceDefinition LoadBalancers ContainerName load error")
	}
	if tt.ServiceDefinition.LoadBalancers[0].ContainerPort != 100 {
		t.Errorf("ServiceDefinition LoadBalancers ContainerPort load error")
	}
	if tt.ServiceDefinition.Role != nil {
		t.Errorf("ServiceDefinition Role load error")
	}
	if tt.ServiceDefinition.DeploymentConfiguration == nil {
		t.Errorf("ServiceDefinition DeploymentConfiguration load error")
	}
	if tt.ServiceDefinition.DeploymentConfiguration.MaximumPercent != 100 {
		t.Errorf("ServiceDefinition DeploymentConfiguration MaximumPercent load error")
	}
	if tt.ServiceDefinition.DeploymentConfiguration.MinimumHealthyPercent != 50 {
		t.Errorf("ServiceDefinition DeploymentConfiguration MinimumHealthyPercent load error")
	}
	if len(tt.ServiceDefinition.PlacementConstraints) == 0 {
		t.Errorf("ServiceDefinition DeploymentConfiguration PlacementConstraints load error")
	}
	if tt.ServiceDefinition.PlacementConstraints[0].Type != "hoge" {
		t.Errorf("ServiceDefinition PlacementConstraints Type load error")
	}
	if tt.ServiceDefinition.PlacementConstraints[0].Expression != "aa" {
		t.Errorf("ServiceDefinition PlacementConstraints Expression load error")
	}
	if len(tt.ServiceDefinition.PlacementStrategy) == 0 {
		t.Errorf("ServiceDefinition DeploymentConfiguration PlacementStrategy load error")
	}
	if tt.ServiceDefinition.PlacementStrategy[0].Type != "fuga" {
		t.Errorf("ServiceDefinition PlacementStrategy Type load error")
	}
	if tt.ServiceDefinition.PlacementStrategy[0].Field != "bb" {
		t.Errorf("ServiceDefinition PlacementStrategy Field load error")
	}
	if *tt.ServiceDefinition.HealthCheckGracePeriodSeconds != 60 {
		t.Errorf("ServiceDefinition HealthCheckGracePeriodSeconds load error")
	}
	if *tt.ServiceDefinition.SchedulingStrategy != "REPLICA" {
		t.Errorf("ServiceDefinition SchedulingStrategy load error")
	}
	if tt.ServiceDefinition.DeploymentController == nil {
		t.Errorf("ServiceDefinition DeploymentController load error")
	}
	if tt.ServiceDefinition.DeploymentController.Type != "EC2" {
		t.Errorf("ServiceDefinition DeploymentController Type load error")
	}
	t.Logf("%v", tt.EcsParams)
}

