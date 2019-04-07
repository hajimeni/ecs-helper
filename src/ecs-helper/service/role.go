package service

import (
    "ecs-helper/option"
    "encoding/json"
    "fmt"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/iam"
    "net/url"
    "os"
    "strings"
    "text/tabwriter"
)

func ListTaskRoles(opt option.RoleListCmdOptions) error {
    awsConfig := aws.Config{}
    sess := session.Must(session.NewSession(&awsConfig))
    svc := iam.New(sess)

    input := &iam.ListRolesInput{}
    roles := []*iam.Role{}
    err := svc.ListRolesPages(input, func(page *iam.ListRolesOutput, lastPage bool) bool {
        roles = append(roles, page.Roles...)
        return page.Marker != nil
    })
    if err != nil {
        fmt.Println(err)
    }
    w := tabwriter.NewWriter(os.Stdout, 20, 2, 8, ' ', 0)
    w.Write([]byte("name\tarn\n"))

    for _, v := range roles {
        if isEcsTaskRole(*v.AssumeRolePolicyDocument) && (opt.Name == "" || strings.Contains(*v.RoleName, opt.Name)) {
            w.Write([]byte(fmt.Sprintf("%s\t%s\n", *v.RoleName, *v.Arn)))
        }
    }
    w.Flush()
    return nil
}


func isEcsTaskRole(assumeRole string) bool {
    data, _ := url.QueryUnescape(assumeRole)
    o := map[string]interface{}{}
    json.Unmarshal([]byte(data), &o)
    for _, s := range o["Statement"].([]interface{}) {
        p, ok := s.(map[string]interface{})["Principal"]
        if ok {
            s, ok := p.(map[string]interface{})["Service"]
            if ok {
                switch t := s.(type){
                case string:
                    if t == "ecs-tasks.amazonaws.com" {
                        return true
                    }
                case []string:
                    for _, ts := range t {
                        if ts == "ecs-tasks.amazonaws.com" {
                            return true
                        }
                    }
                }
            }
        }
    }
    return false
}
