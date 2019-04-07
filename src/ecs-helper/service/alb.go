package service

import (
    "ecs-helper/option"
    "fmt"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/elbv2"
    "os"
    "strings"
    "text/tabwriter"
)

func AlbTglist(opt option.AlbTglistCmdOptions) error {
    awsConfig := aws.Config{}
    if opt.Region != "" {
        awsConfig.WithRegion(opt.Region)
    }
    sess := session.Must(session.NewSession(&awsConfig))
    svc := elbv2.New(sess)

    input := &elbv2.DescribeTargetGroupsInput{}
    targetGroups := []*elbv2.TargetGroup{}
    err := svc.DescribeTargetGroupsPages(input, func(page *elbv2.DescribeTargetGroupsOutput, lastPage bool) bool {
        targetGroups = append(targetGroups, page.TargetGroups...)
        return page.NextMarker != nil
    })
    if err != nil {
        fmt.Println(err)
    }
    w := tabwriter.NewWriter(os.Stdout, 20, 2, 8, ' ', 0)

    w.Write([]byte("name\tarn\n"))
    for _, v := range targetGroups {
        if opt.Name == "" || strings.Contains(*v.TargetGroupName, opt.Name){
            w.Write([]byte(fmt.Sprintf("%s\t%s\n", *v.TargetGroupName, *v.TargetGroupArn)))
        }
    }
    w.Flush()
    return nil
}
