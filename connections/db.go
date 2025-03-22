package connection

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitConnection() {
	dsn := "user=Authora password=Authora dbName=Authora sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	sqlDb, err := db.DB()

	if err != nil {
		log.Fatal("Failed to get database instance:", err)
	}

	err := sqlDb.Ping()
	if err != nil {

		log.Fatal("Failed to ping database:", err)
	} else {
		fmt.Println("Connected to the database successfully!")
	}
}
