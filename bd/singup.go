package bd

import (
	"fmt"

	".gambituser/models"
	".gambituser/tools"
	_ "github.com/go-sql-driver/mysql"
)

func SignUp(sig models.SignUp) error {
	fmt.Println("Comienza Registro")

	err := DbConnect()
	if err != nil {
		return err
	}

	defer Db.Close()

	sentencia := "INSERT INTO users (User_Email, User_UUID, User_DateAdd) VALUES ('"+sig.UserEmail+"', '"+sig.UserUUID+"','"+tools.DateMySQL()+"')"
	fmt.Println("Sentencia", sentencia)

	_, err = Db.Exec(sentencia)
	if err != nil {
		fmt.Println("Error al registrar usuario", err)
		return err
	}

	return nil
}


