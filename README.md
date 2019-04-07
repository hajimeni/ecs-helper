ecs-helper
---

## Motivation

- I want to configuration file because `ecs-cli compose service` command has so many many parameters. 
    ```
    ecs-cli compose service create [--deployment-max-percent n] [--deployment-min-healthy-percent n] [--load-balancer-name value|--target-group-arn value] [--container-name value] [--container-port value] [--role value] [--launch-type launch_type] [--health-check-grace-period integer] [--create-log-groups] [--enable-service-discovery] [--vpc value] [--private-dns-namespace value] [--private-dns-namespace-id value] [--public-dns-namespace value] [--public-dns-namespace-id value] [--sd-container-name value] [--sd-container-port value] [--dns-ttl value] [--dns-type value] [--healthcheck-custom-config-failure-threshold value] [--scheduling-strategy value] [--tags key1=value1,key2=value2] [--disable-ecs-managed-tags] [--help]
    ```
- I want to get `ALB/TG`, `iam`, `cloudwatchevents`, `route53` resources using one binary.
- I want to describe unsupported options in the configuration file.
- I want to 

## Latest version

__Please note that 0.3.0 and 0.2.0 are not compatible__

- `v0.3.0`
  - (Download (for macOS))[https://github.com/hajimeni/ecs-helper/releases/download/v0.3.0/ecs-helper-darwin-amd64.tar.gz]
  - (Download (for Linux))[https://github.com/hajimeni/ecs-helper/releases/download/v0.3.0/ecs-helper-linux-amd64.tar.gz]

- `v0.2.0`
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
- scale up ecs-service
    ```
    ecs-helper service scale --count 3 \
      --config ecs-config.yml
    ```
- update ecs-service and task definition
    ```
    ecs-helper service up --file docker-compose.yml --file docker-compose.override.yml \
      ---config ecs-config.yml
    ```
- delete ecs-service
    ```
    ecs-helper service delete --name service_name --file docker-compose.yml --file docker-compose.override.yml \
      --params ecs-params.yml --config ecs-config.yml
    ```

1. Create docker-compose.yml (+ docker-compose.override.yml and other `docker-compose` format yml files)
1. Create `ecs-config.yml` (Format is [here](#ecs-config.yml) )
1. Deploy ecs Service
    ```
    $ ecs-helper deploy -f docker-compose.yml -f docker-compose.override.yml -c ecs-config.yml
    ```
    `deploy` command is nearly like `ecs-cli compose service up`
      - Create ECS Task definition (or update Task definition version)
      - Create ECS Service (or update ECS Service)
      - Wait for ECS Service desired count equals to Task count.
      - Wait for ALB Health check if needed.
      - if new revision ECS Task is stopped, roll back old revision.

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

## Service commands

### deploy

Nearly like `ecs-cli compose service up`

### erase

Nearly like `ecs-cli service rm`

### scale

Nearly like `ecs-cli service scale`

### stop

Nearly like `ecs-cli service stop`

## Resource commands

### alb

### `alb tglist`

`aws elbv2 describe-taget-groups`
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

### `ecr create`

`aws ecr create-repository`

```
Usage:
  ecs-helper ecr create [flags]

Flags:
  -h, --help            help for create
  -n, --name string     
  -r, --region string
```

### `ecr list`

`aws ecr describe-repositories`

```
Usage:
  ecs-helper ecr list [flags]

Flags:
  -h, --help            help for list
  -n, --name string     
  -r, --region string
```

### `ecr images`

all repositories `aws ecr describe-images`

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

### get-login

```
Usage:
  ecs-helper get-login [flags]

Flags:
  -h, --help                 help for get-login
  -r, --region string        
  -i, --registry-id string
```

### role

### `role list`

list only Assumed `ecs-tasks.amazonaws.com` `aws iam list-roles`

```
Usage:
  ecs-helper role list [flags]

Flags:
  -h, --help          help for list
  -n, --name string
```


### cli

perfect wrapper for `ecs-cli`
command and flags are (here)[https://docs.aws.amazon.com/AmazonECS/latest/developerguide/ECS_CLI_reference.html]

```
$ ecs-helper cli [global options] command [command options] [arguments...]
```
