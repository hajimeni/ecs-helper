FROM golang:latest AS builder
ENV GOOS=linux \
    CGO_ENABLED=0 \
    GOARCH=amd64
WORKDIR /go/src/ecs-helper
COPY . .
RUN make

FROM amazonlinux:1

COPY --from=builder /go/src/ecs-helper/ecs-helper /bin/ecs-helper

RUN mkdir -p /var/ecs-helper/.bin \
 && curl -sL -o /var/ecs-helper/.bin/ecs-cli https://s3.amazonaws.com/amazon-ecs-cli/ecs-cli-linux-amd64-latest \
 && chmod +x /var/ecs-helper/.bin/ecs-cli

WORKDIR /var/ecs-helper

ENTRYPOINT ["ecs-helper"]
CMD ["--help"]
