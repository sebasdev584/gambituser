package secretm

import (
	"encoding/json"
	"fmt"

	".gambituser/awsgo"
	".gambituser/models"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

func GetSecret(secretName string) (models.SecretRDSJson, error) {
	var dataSecret models.SecretRDSJson
	fmt.Println("> Pido Secreto" + secretName)
	// Create a Secrets Manager client

	svc := secretsmanager.NewFromConfig(awsgo.Cfg)
	secret, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	})

	if err != nil {
		fmt.Println(err.Error())
		return dataSecret, err
	}

	json.Unmarshal([]byte(*secret.SecretString), &dataSecret)
	fmt.Println(" > Lectura secret OK" + secretName)

	return dataSecret, nil
}