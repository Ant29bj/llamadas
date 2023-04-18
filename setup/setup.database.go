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

	llenarLaamada := `
    INSERT INTO llamadas (fecha_llamada, duracion_llamada, remitente, destinatario) VALUES
    ('2023-04-16', '00:23:47', 'Juan', 'Pedro'),
    ('2023-04-16', '01:10:32', 'Ana', 'Luis'),
    ('2023-04-15', '00:17:54', 'Carlos', 'Sofía'),
    ('2023-04-15', '00:45:23', 'Marta', 'Javier'),
    ('2023-04-14', '00:10:12', 'Lucía', 'Manuel'),
    ('2023-04-14', '00:35:16', 'Sara', 'Pablo'),
    ('2023-04-13', '00:55:34', 'Juan', 'Luisa'),
    ('2023-04-13', '00:20:47', 'Pedro', 'Ana'),
    ('2023-04-12', '00:50:21', 'Sofía', 'Carlos'),
    ('2023-04-12', '00:15:39', 'Javier', 'Marta'),
    ('2023-04-11', '01:05:12', 'Manuel', 'Lucía'),
    ('2023-04-11', '00:30:08', 'Pablo', 'Sara'),
    ('2023-04-10', '00:42:56', 'Luisa', 'Juan'),
    ('2023-04-10', '00:12:34', 'Ana', 'Pedro'),
    ('2023-04-09', '00:23:47', 'Carlos', 'Sofía'),
    ('2023-04-09', '00:50:12', 'Marta', 'Javier'),
    ('2023-04-08', '00:09:38', 'Lucía', 'Manuel'),
    ('2023-04-08', '00:30:15', 'Sara', 'Pablo'),
    ('2023-04-07', '00:20:29', 'Juan', 'Luisa'),
    ('2023-04-07', '00:16:42', 'Pedro', 'Ana'),
    ('2023-04-06', '01:00:09', 'Sofía', 'Carlos'),
    ('2023-04-06', '00:14:37', 'Javier', 'Marta'),
    ('2023-04-05', '00:37:41', 'Manuel', 'Lucía'),
    ('2023-04-05', '00:25:18', 'Pablo', 'Sara'),
    ('2023-04-04', '00:55:27', 'Luisa', 'Juan'),
    ('2023-04-04', '00:09:56', 'Ana', 'Pedro'),
    ('2023-04-03', '00:18:23', 'Carlos', 'Sofía'),
    ('2023-04-03', '00:45:45', 'Marta', 'Javier'),
    ('2023-04-02', '00:11:32', 'Lucía', 'Manuel'),
    ('2023-04-02', '00:22:56', 'Sara', 'Pablo');`

	_, err := DBConn.Exec(tablaLlamadas)
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = DBConn.Exec(llenarLaamada)
	if err != nil {
		log.Fatal(err.Error())
	}
}
