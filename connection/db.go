package connection

import (
	"Authora/model"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDb() {
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

	err := sqlDb.AutoMigrate(&model.Business{})

	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	log.Println("Database connected and migrations applied successfully")
	return sqlDb
}
