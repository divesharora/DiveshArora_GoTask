package database

import (
	"log"
	"DiveshArora_GoTask/entity"
	"github.com/jinzhu/gorm"
)

//Connector variable used for CRUD operation's
var Connector *gorm.DB

//Connect creates MySQL connection
func Connect(connectionString string) error {
	var err error
	Connector, err = gorm.Open("mysql", connectionString)
	if err != nil {
		return err
	}
	log.Println("Connection was successful!!")
	return nil
}

func Migrate(table *entity.User) {
	Connector.AutoMigrate(&table)
	log.Println("Table migrated")
}
func MigrateLikes(table *entity.Likes) {
	Connector.AutoMigrate(&table)
	log.Println("Table migrated")
}