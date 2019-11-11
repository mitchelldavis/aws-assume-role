package main

import (
    "os"
    "flag"
    "fmt"
    "github.com/aws/aws-sdk-go/aws/awserr"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/sts"
)

var durationSecondsArg int64
var externalIdArg, roleArnArg, roleSessionNameArg string

func init() {
    flag.Int64Var(&durationSecondsArg, "duration", 3600, "The amount of time, in seconds, to keep the temporary credentials alive.")
    flag.StringVar(&externalIdArg, "externalId", "", "A unique identifier that may be required to assume the role.")
    flag.StringVar(&roleArnArg, "roleArn", "", "The Arn of the role to assume.")
    flag.StringVar(&roleSessionNameArg, "sessionName", "", "A unique identifier for the session so you can distinguish between to prinipals assuming the same role.")
    flag.Parse()
}

func main() {
    if durationSecondsArg < 3600 || durationSecondsArg > 43200 {
        fmt.Println("You must specify a duration between 1 and 12 hours.")
        os.Exit(1)
    }

    if roleArnArg == "" {
        fmt.Println("You must supply a role arn.")
        os.Exit(1)
    }

    if roleSessionNameArg == "" {
        fmt.Println("You must supply a role session name.")
        os.Exit(1)
    }

    var externalId *string = nil

    if externalIdArg != "" {
        externalId = &externalIdArg
    }

    svc := sts.New(session.New())
    input := &sts.AssumeRoleInput{
        DurationSeconds: &durationSecondsArg,
        ExternalId: externalId,
        RoleArn: &roleArnArg,
        RoleSessionName: &roleSessionNameArg,
    }

    result, err := svc.AssumeRole(input)
    if err != nil {
        if aerr, ok := err.(awserr.Error); ok {
            switch aerr.Code() {
            case sts.ErrCodeMalformedPolicyDocumentException:
                fmt.Println(sts.ErrCodeMalformedPolicyDocumentException, aerr.Error())
            case sts.ErrCodePackedPolicyTooLargeException:
                fmt.Println(sts.ErrCodePackedPolicyTooLargeException, aerr.Error())
            case sts.ErrCodeRegionDisabledException:
                fmt.Println(sts.ErrCodeRegionDisabledException, aerr.Error())
            default:
                fmt.Println(aerr.Error())
            }
        } else {
            // Print the error, cast err to awserr.Error to get the Code and
            // Message from an error.
            fmt.Println(err.Error())
        }
        os.Exit(1)
    }

    fmt.Printf("export AWS_ACCESS_KEY_ID=%s\n", *result.Credentials.AccessKeyId)
    fmt.Printf("export AWS_SECRET_ACCESS_KEY=%s\n", *result.Credentials.SecretAccessKey)
    fmt.Printf("export AWS_SESSION_TOKEN=%s\n", *result.Credentials.SessionToken)
}
