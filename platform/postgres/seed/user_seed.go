package seed

import (
	"attempt4/internal/domain/entity"
	"attempt4/internal/domain/enum"
	"attempt4/platform/hash"
	"gorm.io/gorm"
	"time"
)

func UserSeed(db *gorm.DB) {
	userUsername := [2]string{"burak12570", "Fanahey"}
	userPassword := [2]string{"12345678", "1234578a"}
	userMail := [2]string{"burakcanheyal@gmail.com", "fatihmeral@outlook.com"}
	userName := [2]string{"Burak can", "Fatih"}
	userSurname := [2]string{"Heyal", "Meral"}
	userPhone := [2]string{"+905316519484", "+905316519424"}
	firstUserBirthDate := time.Date(2000, time.Month(9), 18, 0, 0, 0, 0, time.UTC)
	secondUserBirthDate := time.Date(1999, time.Month(5), 24, 0, 0, 0, 0, time.UTC)
	users := []entity.User{
		{0,
			userUsername[0],
			userPassword[0],
			userMail[0],
			userName[0],
			userSurname[0],
			userPhone[0],
			enum.UserActiveStatus,
			nil,
			nil,
			&firstUserBirthDate,
			time.Now(), nil, nil},

		{0,
			userUsername[1],
			userPassword[1],
			userMail[1],
			userName[1],
			userSurname[1],
			userPhone[0],
			enum.UserActiveStatus,
			nil,
			nil,
			&secondUserBirthDate,
			time.Now(), nil, nil},
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
