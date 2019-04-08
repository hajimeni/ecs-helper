ecs-helper
---

## Motivation

- I want to configuration file because `ecs-cli compose service` command has so many many parameters. 
    ```
    ecs-cli compose service create [--deployment-max-percent n] [--deployment-min-healthy-percent n] [--load-balancer-name value|--target-group-arn value] [--container-name value] [--container-port value] [--role value] [--launch-type launch_type] [--health-check-grace-period integer] [--create-log-groups] [--enable-service-discovery] [--vpc value] [--private-dns-namespace value] [--private-dns-namespace-id value] [--public-dns-namespace value] [--public-dns-namespace-id value] [--sd-container-name value] [--sd-container-port value] [--dns-ttl value] [--dns-type value] [--healthcheck-custom-config-failure-threshold value] [--scheduling-strategy value] [--tags key1=value1,key2=value2] [--disable-ecs-managed-tags] [--help]
    ```
- I want to get `ALB/TG`, `iam`, `cloudwatchevents`, `route53` resources using one binary.
- I want to describe unsupported options in the configuration file.

## Latest version

__Please note that 0.3.0 and 0.2.0 are not compatible__

- `v0.3.0`
  - (Download (for macOS))[https://github.com/hajimeni/ecs-helper/releases/download/v0.3.0/ecs-helper-darwin-amd64.tar.gz]
  - (Download (for Linux))[https://github.com/hajimeni/ecs-helper/releases/download/v0.3.0/ecs-helper-linux-amd64.tar.gz]

- `v0.2.0`(depcated)
  - (Download (for macOS))[https://github.com/hajimeni/ecs-helper/releases/download/v0.2.0/ecs-helper-darwin-amd64.tar.gz]
  - (Download (for Linux))[https://github.com/hajimeni/ecs-helper/releases/download/v0.2.0/ecs-helper-linux-amd64.tar.gz]

## Usage

__Please be careful because it describes only how to use 0.3.0__

#### Simple example

- create ecs-service
    ```
    ecs-helper service create --file docker-compose.yml --file docker-compose.override.yml \
      --config ecs-config.yml
    ```
  - like `ecs-cli compose create`
- scale up ecs-service
    ```
    ecs-helper service scale --count 3 \
      --config ecs-config.yml
    ```
  - like `ecs-cli compose scale`
- update ecs-service and task definition
    ```
    ecs-helper service deploy --file docker-compose.yml --file docker-compose.override.yml \
      ---config ecs-config.yml
    ```
  - like `ecs-cli compose up`
- delete ecs-service
    ```
    ecs-helper service delete --name service_name --file docker-compose.yml --file docker-compose.override.yml \
      --params ecs-params.yml --config ecs-config.yml
    ```
  - like `ecs-cli compose delete`

## Required policies
```
{
    "Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "Sample",
            "Effect": "Allow",
            "Action": [
              "elasticloadbalancing:DescribeTargetGroups",
              "ecr:DescribeImages",
              "ecr:DescribeRepositories",
              "ecr:CreateRepository",
              "ecr:GetAuthorizationToken",
              "ecr:BatchGetImage",
              "ecr:PutImage",
              "iam:ListRoles",
              "ecs:*"
            ],
            "Resource": "*"
        }
    ]
}
```

### help

```
ecs-helper --help
```

## ecs-config file format

```
## for validation
version: 1
project_name: string
region: ap-northeast-1

# like service definition
## more information, see https://docs.aws.amazon.com/ja_jp/AmazonECS/latest/developerguide/service_definition_parameters.html
service_definition:
  cluster: string
  task_definition: string
  load_balancers:
    target_group_arn: string
    load_balancer_name: string
    container_name: string
    container_port: 100
  launch_type: string
  role: string
  deployment_configuration:
    maximum_percent: 100
    minimum_healthy_percent: 50
    
  # TODO: this parameter is unused
  placement_constraints:
    - type: string
      expression: string

  # TODO: this parameter is unused
  placement_strategy:
    - type: string
      field: string
  health_check_grace_period_seconds: integer
  scheduling_strategy: string

  # TODO: this parameter is unused
  deployment_controller:
    type: string
  tags:
    - key: hoge
      value: futa
    - key: foo
      value: bar
      
  # TODO: this parameter is unused
  enable_ecs_managed_tags: true
  
  # TODO: this parameter is unused
  propagate_tags: "TASK_DEFINITION"

## equals to ecs_params.yaml
## see https://docs.aws.amazon.com/AmazonECS/latest/developerguide/cmd-ecs-cli-compose-ecsparams.html
ecs_params:
  version: 1
  task_definition:
    ecs_network_mode: string
    task_role_arn: string
    task_execution_role: string
    task_size:
      cpu_limit: string
      mem_limit: string
    pid_mode: string
    ipc_mode: string
    services:
      <service_name>:
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
          driver_opts: boolean
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
```

## Service commands

`ecs-helper service <subcommand>`

### create

`ecs-helper service create`

- like `ecs-cli compose service create`

### deploy

`ecs-helper service deploy`

- like `ecs-cli compose service up`

### scale

`ecs-helper service scale`

- like `ecs-cli compose service scale`

### delete

`ecs-helper service deploy`

- like `ecs-cli compose service delete`

## Task commands

`ecs-helper task <subcommand>`

### create

`ecs-helper task create`

- like `ecs-cli compose create`

### alb

`ecs-helper alb <subcommand>`

### tglist

`ecs-helper alb tglist`

- like `aws elbv2 describe-taget-groups`
https://docs.aws.amazon.com/cli/latest/reference/elbv2/describe-target-groups.html

```
Usage:
  ecs-helper alb tglist [flags]

Flags:
  -h, --help            help for tglist
  -n, --name string     
  -r, --region string
```

### ecr

`ecs-helper ecr <subcommand>`

### create

`ecs-helper ecr create`

- like `aws ecr create-repository`

```
Usage:
  ecs-helper ecr create [flags]

Flags:
  -h, --help            help for create
  -n, --name string     
  -r, --region string
```

### list

`ecs-helper ecr list`

like `aws ecr describe-repositories`

```
Usage:
  ecs-helper ecr list [flags]

Flags:
  -h, --help            help for list
  -n, --name string     
  -r, --region string
```

### images

`ecs-helper ecr create`

like `aws ecr describe-images` in all repositories

```
Usage:
  ecs-helper ecr images [flags]

Flags:
  -h, --help            help for images
  -n, --name string     
  -r, --region string   
      --tagged          
      --untagged
```

## get-login

`ecs-helper ecr get-login`

```
Usage:
  ecs-helper get-login [flags]

Flags:
  -h, --help                 help for get-login
  -r, --region string        
  -i, --registry-id string
```

### role

`ecs-helper role <subcommand>`

### list

`ecs-helper role list`

like `aws iam list-roles`  
list only Assumed `ecs-tasks.amazonaws.com`

```
Usage:
  ecs-helper role list [flags]

Flags:
  -h, --help          help for list
  -n, --name string
```


## cli

perfect wrapper for `ecs-cli`
command and flags are (here)[https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ECS_CLI_reference.html]

```
$ ecs-helper cli [global options] command [command options] [arguments...]
```
