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
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s",
		cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBHost, cfg.DBPort, cfg.SslMode)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Panic("Failed to connect to the database:", err)
		return nil, err
	}
	migration()
	fmt.Println("Database connected and migrations applied successfully")
	migration()
	return DB, nil
}

func migration() {
	DB.AutoMigrate(&model.Business{})
	DB.AutoMigrate(&model.User{})
	fmt.Println("Database migration successfull.")

}
