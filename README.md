ecs-helper
---

## Latest version

- `v0.1.0`
  - (Download (for macOS))[https://github.com/hajimeni/ecs-helper/releases/download/v0.1.0/ecs-helper-darwin-amd64.tar.gz]
  - (Download (for Linux))[https://github.com/hajimeni/ecs-helper/releases/download/v0.1.0/ecs-helper-linux-amd64.tar.gz]

## Usage

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
