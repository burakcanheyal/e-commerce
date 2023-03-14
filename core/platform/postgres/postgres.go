package postgres

import (
	"attempt4/core/internal/domain/entity"
	"attempt4/core/platform/postgres/seed"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func InitializeDatabase(dsn string) *gorm.DB {
	db := ConnectToDb(dsn)
	err := db.AutoMigrate(&entity.Product{}, &entity.User{}, &entity.Order{}, &entity.Key{})
	if err != nil {
		return nil
	}
	
	seed.ProductSeed(db)
	seed.UserSeed(db)
	seed.RolSeed(db)

	return db
}
func ConnectToDb(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database")
	}
	return db
}
