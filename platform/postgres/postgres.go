package postgres

import (
	entity2 "attempt4/internal/domain/entity"
	seed2 "attempt4/platform/postgres/seed"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func InitializeDatabase(dsn string) *gorm.DB {
	db := ConnectToDb(dsn)
	err := db.AutoMigrate(
		&entity2.Product{},
		&entity2.User{},
		&entity2.Order{},
		&entity2.Role{},
		&entity2.Wallet{},
		&entity2.Submission{},
		&entity2.WalletOperation{},
	)
	if err != nil {
		return nil
	}

	seed2.UserSeed(db)
	seed2.ProductSeed(db)
	seed2.RolSeed(db)
	seed2.WalletSeed(db)

	return db
}
func ConnectToDb(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect the database")
	}
	return db
}
