package seed

import (
	"attempt4/internal/domain/entity"
	"attempt4/internal/domain/enum"
	"gorm.io/gorm"
	"time"
)

func FeedbackSeed(db *gorm.DB) {
	feedbacks := []entity.Feedback{
		{
			Id:          0,
			Description: "çok güzel",
			Star:        5,
			TripId:      3,
			UserId:      1,
			Status:      enum.TripActive,
			CreatedAt:   time.Now(),
			UpdatedAt:   nil,
			DeletedAt:   nil,
			Trip:        entity.Trip{},
			User:        entity.User{},
		},
		{
			Id:          0,
			Description: "mükemmel",
			Star:        5,
			TripId:      3,
			UserId:      1,
			Status:      enum.TripActive,
			CreatedAt:   time.Now(),
			UpdatedAt:   nil,
			DeletedAt:   nil,
			Trip:        entity.Trip{},
			User:        entity.User{},
		},
	}
	var size int64
	db.Model(&feedbacks).Count(&size)
	if size == 0 {
		for _, p := range feedbacks {
			db.Create(&p)
		}
	}
}
