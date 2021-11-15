package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"

	awsSecretsManager "github.com/aws/aws-sdk-go/service/secretsmanager"
)

type secretsManager struct {
	SecretsManager *awsSecretsManager.SecretsManager
}

func secret() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(endpoints.ApNortheast1RegionID),
		Credentials: credentials.NewStaticCredentials(
			"AWSアクセスキー",
			"AWSシークレットアクセスキー",
			"",
		),
	})

	secretsManager := &secretsManager{awsSecretsManager.New(sess)}
	input := &awsSecretsManager.GetSecretValueInput{
		SecretId: aws.String("secrets managerのシークレットの名前"),
	}
	value, err := secretsManager.SecretsManager.GetSecretValue(input)
	if err != nil {
		panic(err)
	}
	fmt.Println(value)
}
