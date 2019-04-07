package service

import (
    "ecs-helper/option"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "fmt"
    "strings"
    "text/tabwriter"
    "os"
    "github.com/aws/aws-sdk-go/service/ecr"
    "github.com/docker/go-units"
    "time"
    "errors"
)

func ListEcrRepositories(opt option.EcrListCmdOptions) error {
    repositories, err := listRepositories(opt)
    if err != nil {
        fmt.Println(err)
        return err
    }
    w := tabwriter.NewWriter(os.Stdout, 20, 2, 8, ' ', 0)
    w.Write([]byte("ID\tURI\n"))
    for _, v := range repositories {
        w.Write([]byte(fmt.Sprintf("%s\t%s\n", *v.RepositoryName, *v.RepositoryUri)))
    }
    w.Flush()
    return nil
}

func listRepositories(opt option.EcrListCmdOptions) ([]*ecr.Repository, error) {
    awsConfig := aws.Config{}
    if opt.Region != "" {
        awsConfig.WithRegion(opt.Region)
    }
    sess := session.Must(session.NewSession(&awsConfig))
    svc := ecr.New(sess)
    input := &ecr.DescribeRepositoriesInput{}

    repositories := []*ecr.Repository{}
    err := svc.DescribeRepositoriesPages(input, func(page *ecr.DescribeRepositoriesOutput, lastPage bool) bool {
        repositories = append(repositories, page.Repositories...)
        return page.NextToken != nil
    })
    if err != nil {
        return nil, err
    }
    result := []*ecr.Repository{}
    for _, v := range repositories {
        if opt.Name == "" || strings.Contains(*v.RepositoryName, opt.Name){
            result = append(result, v)
        }
    }
    return result, nil
}


func ListEcrImages(opt option.EcrImagesCmdOptions) error {
    awsConfig := aws.Config{}
    if opt.Region != "" {
        awsConfig.WithRegion(opt.Region)
    }
    sess := session.Must(session.NewSession(&awsConfig))
    svc := ecr.New(sess)

    repositories, err := listRepositories(option.EcrListCmdOptions{
        Name: opt.Name,
        Region: opt.Region,
    })
    if err != nil {
        fmt.Println(err)
        return err
    }
    var imageFilter *ecr.DescribeImagesFilter
    if opt.Tagged || opt.UnTagged {
        imageFilter = &ecr.DescribeImagesFilter{}
        if opt.Tagged {
            imageFilter.SetTagStatus("TAGGED")
        } else if opt.UnTagged {
            imageFilter.SetTagStatus("UNTAGGED")
        }
    }

    images := []*ecr.ImageDetail{}
    for _, r := range repositories {
        input := &ecr.DescribeImagesInput{
            RepositoryName: r.RepositoryName,
        }
        if imageFilter != nil {
            input.SetFilter(imageFilter)
        }
        svc.DescribeImagesPages(input, func(page *ecr.DescribeImagesOutput, lastPage bool) bool {
            images = append(images, page.ImageDetails...)
            return page.NextToken != nil
        })
    }


    w := tabwriter.NewWriter(os.Stdout, 20, 2, 8, ' ', 0)

    cnt := 0
    for _, v := range images {
        if opt.Name == "" || strings.Contains(*v.RepositoryName, opt.Name){
            repo := repositoryUrl(*v.RepositoryName, *v.RegistryId, svc.SigningRegion)
            digest := strings.Split(*v.ImageDigest, ":")[1][0:12]
            pushedAt := units.HumanDuration(time.Now().UTC().Sub(time.Unix(v.ImagePushedAt.Unix(), 0)))
            size := units.HumanSizeWithPrecision(float64(aws.Int64Value(v.ImageSizeInBytes)), 3)

            if len(v.ImageTags) == 0 {
                tag := "<none>"
                printImageContent(w, &cnt, repo, tag, digest, pushedAt, size)
            } else {
                for _, tag := range v.ImageTags {
                    printImageContent(w, &cnt, repo, *tag, digest, pushedAt, size)
                }
            }
        }
    }
    w.Flush()
    return nil
}

func repositoryUrl(name string, id string, region string) string {
    return fmt.Sprintf("%s.dkr.ecr.%s.amazonaws.com/%s", id, region, name)
}


func printImageContent(w *tabwriter.Writer, cnt *int, args ...string) {
    if *cnt == 0 {
        w.Write([]byte("REPOSITORY URI\tTAG\tIMAGE ID\tCREATED\tSIZE\n"))
    }
    w.Write([]byte(fmt.Sprintf("%s\t%s\t%s\t%s\t%s\n", args[0], args[1], args[2], args[3], args[4])))
    *cnt++
    if *cnt >= 50 {
        w.Flush()
        *cnt = 0
    }
}


func CreateEcrRepository(opt option.EcrCreteCmdOptions) error {
    if opt.Name == "" {
        return errors.New("Repository name is required")
    }

    awsConfig := aws.Config{}
    if opt.Region != "" {
        awsConfig.WithRegion(opt.Region)
    }
    sess := session.Must(session.NewSession(&awsConfig))
    svc := ecr.New(sess)

    input := &ecr.CreateRepositoryInput{
        RepositoryName: &opt.Name,
    }
    v, err := svc.CreateRepository(input)

    if err != nil {
        fmt.Println(err)
        return err
    }
    repo := v.Repository
    out, err := svc.DescribeRepositories(&ecr.DescribeRepositoriesInput{
        RepositoryNames: []*string{repo.RepositoryName},
    })
    if err != nil {
        fmt.Println(err)
        return err
    }
    repo = out.Repositories[0]

    fmt.Printf("repository created: %s ", *repo.RepositoryUri)
    return nil
}
