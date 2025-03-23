package connection

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/erfuuan/Authora/conf"
	"github.com/erfuuan/Authora/model"
)

var DB *gorm.DB

func InitDb(cfg *conf.Config) (*gorm.DB, error) {
	var err error
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBHost, cfg.DBPort)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
		return nil, err
	}
	migration()
	log.Println("Database connected and migrations applied successfully")
	migration()
	return DB, nil
}

func migration() {
	log.Println("Database migration started.")

	DB.AutoMigrate(&model.Business{})
	DB.AutoMigrate(&model.User{})

}
