package seed

import (
	"attempt4/core/internal/domain/entity"
	"attempt4/core/platform/hash"
	"gorm.io/gorm"
	"time"
)

func UserSeed(db *gorm.DB) {
	//Todo:Phone number ekle
	users := []entity.User{
		{0,
			"burak12570",
			"12345678",
			"burakcanheyal@gmail.com ",
			"Burak Can",
			"Heyal",
			1,
			"412563",
			time.Now(),
			time.Date(2000, time.Month(9), 18, 0, 0, 0, 0, time.UTC)},

		{0,
			"Fanahey",
			"1234578a",
			"fatihmeral@outlook.com",
			"Fatih",
			"Meral",
			1,
			"947628",
			time.Now(),
			time.Date(1999, time.Month(5), 24, 0, 0, 0, 0, time.UTC)},
	}

	var size int64
	db.Model(&users).Count(&size)
	if size == 0 {
		for _, u := range users {
			encryptPass, _ := hash.EncryptPassword(u.Password)
			u.Password = encryptPass
			db.Create(&u)
		}
	}
}
