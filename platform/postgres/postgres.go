package postgres

import (
	"attempt4/internal/domain/entity"
	"attempt4/platform/app_log"
	"attempt4/platform/postgres/seed"
	"attempt4/platform/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializeDatabase(dsn string) (*gorm.DB, error) {
	db := ConnectToDb(dsn)
	err := db.AutoMigrate(
		&entity.Product{},
		&entity.User{},
		&entity.Order{},
		&entity.Role{},
		&entity.Wallet{},
		&entity.Submission{},
		&entity.WalletOperation{},
		&app_log.ApplicationLog{},
		&entity.Trip{},
		&entity.Question{},
		&entity.TopicPoints{},
	)
	if err != nil {
		return nil, err
	}

	seed.UserSeed(db)
	seed.ProductSeed(db)
	seed.RolSeed(db)
	seed.WalletSeed(db)
	seed.TripSeed(db)
	seed.QuestionSeed(db)

	return db, nil
}
func ConnectToDb(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		zap.Logger.Fatalf("Failed to connect the database %s", err)
	}
	return db
}
