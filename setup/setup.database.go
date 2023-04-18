package setup

import (
	"database/sql"
	"io/ioutil"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v3"
)

var GlobalConfig Config = SetConfigInfo()
var DBConn sql.DB = SetDatabaseConnection()

func SetConfigInfo() Config {
	configFile, err := ioutil.ReadFile("config.yml")
	if err != nil {
		log.Fatal(err)
	}

	var config Config

	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		panic(err)
	}

	return config
}

func SetDatabaseConnection() sql.DB {
	dbUrlInfo :=
		GlobalConfig.Database.Username +
			":" + GlobalConfig.Database.Password +
			"@tcp(" + GlobalConfig.Database.Host +
			":" + GlobalConfig.Database.Port + ")/" +
			GlobalConfig.Database.Name

	//"root:root@tcp(127.0.0.1:3306)/llamadas"

	db, err := sql.Open(GlobalConfig.Database.Type, dbUrlInfo)
	if err != nil {
		log.Fatal(err.Error())
	}
	return *db

}

func SetDatabase() {

	tablaLlamadas := `
    CREATE TABLE IF NOT EXISTS llamadas (
      id INT AUTO_INCREMENT PRIMARY KEY,
      fecha_llamada DATE,
      duracion_llamada TIME,
      remitente VARCHAR(255),
      destinatario VARCHAR(255)
    );`
	_, err := DBConn.Exec(tablaLlamadas)
	if err != nil {
		log.Fatal(err.Error())

	}
}
