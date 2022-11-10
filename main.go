package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

func main() {

	sess, err := session.NewSession(&aws.Config{Region: aws.String("us-east-1")})
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
