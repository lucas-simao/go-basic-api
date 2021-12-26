package database

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/mitchellh/mapstructure"
)

var connectionDB *sqlx.DB

func Connect() *sqlx.DB {

	var (
		db  *sqlx.DB
		err error
	)

	if connectionDB != nil {
		db = connectionDB
		return db
	}

	databaseConfig := readEnvs()

	db, err = sqlx.Connect(databaseConfig.Type, fmt.Sprintf("user=%s port=%s password=%s host=%s dbname=%s sslmode=%s",
		databaseConfig.Username,
		databaseConfig.Port,
		databaseConfig.Password,
		databaseConfig.Host,
		databaseConfig.Database,
		databaseConfig.SSL,
	))

	if err != nil {
		panic(err)
	}

	connectionDB = db

	return db
}

func readEnvs() (databaseConfig DatabaseConfig) {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err.Error())
	}

	t := reflect.TypeOf(databaseConfig)

	mapDatabase := make(map[string]interface{})

	for i := 0; i < t.NumField(); i++ {
		mapDatabase[t.Field(i).Name] = os.Getenv(strings.ToUpper(t.Field(i).Name))
	}

	err = mapstructure.Decode(mapDatabase, &databaseConfig)

	if err != nil {
		log.Fatal(err.Error())
	}

	return
}
