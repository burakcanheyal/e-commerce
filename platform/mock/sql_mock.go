package mock

import (
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func MockedDb() (*gorm.DB, sqlmock.Sqlmock) {
	testDB, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		panic(err)
	}
	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 testDB,
		PreferSimpleProtocol: true,
	})

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		panic("Cannot open stub database")
	}

	return db, mock
}

/*

testDB, _, err := sqlmock.New()
	if err != nil {
		panic("sqlmoc.New() occurs an error")
	}

	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 testDB,
		PreferSimpleProtocol: true,
	})

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		panic("Cannot open stub database")
	}

	err = db.AutoMigrate(
		&entity.Product{},
		&entity.User{},
		&entity.Order{},
		&entity.Role{},
		&entity.Wallet{},
		&entity.Submission{},
		&entity.WalletOperation{},
	)
	if err != nil {
		panic("mock does not work")
	}
	return db
*/
