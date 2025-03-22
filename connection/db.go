package connection

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/erfuuan/Authora/model"
)

var DB *gorm.DB

func InitDb() (*gorm.DB, error) {
	var err error
	dsn := "user=Authora password=Authora dbname=Authora host=localhost port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
		return nil, err
	}

	DB.AutoMigrate(&model.Business{})

	log.Println("Database connected and migrations applied successfully")
	migration()
	return DB, nil
}

func migration() {
	log.Println("Database migration started.")

	DB.AutoMigrate(&model.Business{})

}
