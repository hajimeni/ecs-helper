package cmd

import (
    "ecs-helper/config"
    "github.com/spf13/cobra"
    "os"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/ecr"
    "fmt"
    "encoding/base64"
    "strings"
    "github.com/aws/aws-sdk-go/aws"
)

var (
    getLoginCmdOptions = &config.GetLoginCmdOptions{}
)

func init() {
    getLoginCmd.SetOutput(os.Stdout)
    getLoginCmd.PersistentFlags().StringVarP(&getLoginCmdOptions.Region, "region", "r", "", "")
    getLoginCmd.PersistentFlags().StringVarP(&getLoginCmdOptions.RegistryId, "registry-id", "i", "", "")

    RootCmd.AddCommand(getLoginCmd)
}

var getLoginCmd = &cobra.Command{
    Use: "get-login",
    Short: "",
    Long: "",
    RunE: func(cmd *cobra.Command, args []string) error {
        config := aws.Config{}
        if getLoginCmdOptions.Region != "" {
            config.WithRegion(getLoginCmdOptions.Region)
        }
        sess := session.Must(session.NewSession(&config))
        svc := ecr.New(sess)
        input := &ecr.GetAuthorizationTokenInput{}
        if getLoginCmdOptions.RegistryId != "" {
            input.SetRegistryIds([]*string{&getLoginCmdOptions.RegistryId})
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
