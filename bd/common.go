package bd

import (
	"database/sql"
	"fmt"
	"os"

	".gambituser/models"
	".gambituser/secretm"
	_ "github.com/go-sql-driver/mysql"
)

var SecretModel models.SecretRDSJson
var err error
var Db *sql.DB

func ReadSecret() error {
	SecretModel, err = secretm.GetSecret(os.Getenv("SecretName"))

	return err
}

func DbConnect() error {
	ReadSecret()

	Db, err = sql.Open("mysql", ConnStr(SecretModel))

	if err != nil {
		fmt.Println("Error al conectar con la base de datos" + err.Error())
		return err
	}

	err = Db.Ping()
	if err != nil {
		fmt.Println("Error al conectar con la base de datos", err.Error())
		return err
	}

	fmt.Println("Conexion exitosa")

	return nil
}

func ConnStr(clave models.SecretRDSJson) string {
	var dbUser, authToken, dbEndpoint, DbName string
	dbUser = clave.Username
	authToken = clave.Password
	dbEndpoint = clave.Host
	DbName = "gambit"

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?allowCleartextPasswords=true", dbUser, authToken, dbEndpoint, DbName)
	fmt.Println("Dsn db" + dsn)
	return dsn
}