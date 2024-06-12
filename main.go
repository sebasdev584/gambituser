package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	".gambituser/awsgo"
	".gambituser/bd"
	".gambituser/models"
	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(EjecutoLambda)
}

func EjecutoLambda(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {
	awsgo.InicializoAWS()

	if !ValidoParametros() {
		fmt.Println("Error en los parámetros. debe enviar 'SecretName'")
		err := errors.New("error en parámetros debe enviar SecretName")

		return event, err
	}

	var data models.SignUp

	for row, att := range event.Request.UserAttributes {
		if row == "email" {
			data.UserEmail = att
			fmt.Println("Email", data.UserEmail)
		}

		if row == "sub" {
			data.UserUUID = att
			fmt.Println("Sub = " + data.UserUUID)
		}
	}

	err := bd.ReadSecret()
	if err != nil {
		fmt.Println("Error al leer secret" + err.Error())
		return event, err
	}

	err = bd.SignUp(data)

	return event, err

}

func ValidoParametros() bool {
	var getParam bool

	_, getParam = os.LookupEnv("SecretName")

	return getParam
}