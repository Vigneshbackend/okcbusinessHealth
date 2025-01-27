package driver

import (
	"errors"
	"fmt"
	"os"

	"DeviceConnect/model"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"

	//it is the db connector for postgres
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var err error

func initialMigration() {

	db.AutoMigrate(&model.Transactiondetails{})
	db.AutoMigrate(&model.Accounts{})
	db.AutoMigrate(&model.Merchantdetails{})

}

//Connect function returns database pointer
func Connect() (*gorm.DB, error) {

	err := godotenv.Load(".env")

	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")

	fmt.Println(username, password, port, dbName, host)
	if username == "" || password == "" || port == "" || dbName == "" {
		err := errors.New("Environmental variables not found to connect database")
		fmt.Println(err)
		return nil, err
	}

	dbConfig := "port=" + port + " user=" + username + " dbname=" + dbName + " password=" + password + " host=" + host + " sslmode=disable"

	db, err = gorm.Open("postgres", dbConfig)
	if err != nil {
		fmt.Println("Failed to connect database", err)
	} else {
		fmt.Println("Database connection Successfull")
		initialMigration()
	}

	return db, err

}
