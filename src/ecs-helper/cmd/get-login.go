package cmd

import (
    "ecs-helper/option"
    "encoding/base64"
    "fmt"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/ecr"
    "github.com/spf13/cobra"
    "strings"
)

func NewCmdGetLogin() *cobra.Command {
    o := option.GetLoginCmdOptions{}

    cmd := &cobra.Command{
        Use: "get-login",
        Short: "",
        RunE: func(cmd *cobra.Command, args []string) error {
            config := aws.Config{}
            if o.Region != "" {
                config.WithRegion(o.Region)
            }
            sess := session.Must(session.NewSession(&config))
            svc := ecr.New(sess)
            input := &ecr.GetAuthorizationTokenInput{}
            if o.RegistryId != "" {
                input.SetRegistryIds([]*string{&o.RegistryId})
            }

            result, err := svc.GetAuthorizationToken(input)
            if err != nil {
                return err
            }
            for _, v := range result.AuthorizationData {
                data, _ := base64.StdEncoding.DecodeString(*v.AuthorizationToken)
                endpoint := v.ProxyEndpoint
                splitted := strings.Split(string(data), ":")
                fmt.Printf("docker login -u %s -p %s %s", splitted[0], splitted[1], *endpoint)
                return nil
            }
            return nil
        },
    }
    cmd.PersistentFlags().StringVarP(&o.Region, "region", "r", "", "")
    cmd.PersistentFlags().StringVarP(&o.RegistryId, "registry-id", "i", "", "")

    return cmd
}
