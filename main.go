package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

func main() {
	regions := []string{
		"us-east-1",
		"us-east-2",
		"us-west-1",
		"us-west-2",
		"ap-northeast-1",
		"ap-northeast-2",
		"ap-northeast-3",
		"ap-south-1",
		"ap-southeast-1",
		"ap-southeast-2",
		"ca-central-1",
		"eu-central-1",
		"eu-west-1",
		"eu-west-2",
		"eu-west-3",
		"eu-north-1",
		"sa-east-1",
	}

	for i := range regions {

		sess, err := session.NewSession(&aws.Config{Region: aws.String(regions[i])})
		if err != nil {
			fmt.Println(err)
			return
		}

		svc := secretsmanager.New(sess)

		max := int64(100)
		listinput := &secretsmanager.ListSecretsInput{
			MaxResults: &max,
		}

		result, err := svc.ListSecrets(listinput)
		if err != nil {

			fmt.Printf("error inlisting secrets: %+v", err.Error())
			return
		}
		delete := true
		for i := range result.SecretList {
			input := &secretsmanager.DeleteSecretInput{
				SecretId:                   result.SecretList[i].ARN,
				ForceDeleteWithoutRecovery: &delete,
			}

			_, err := svc.DeleteSecret(input)
			if err != nil {

				fmt.Printf("error in deleting secret: %+v", err.Error())
				return
			}
		}
	}

}
