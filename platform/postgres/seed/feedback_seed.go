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
			ProductId:   3,
			UserId:      1,
			Status:      enum.TripActive,
			CreatedAt:   time.Now(),
			UpdatedAt:   nil,
			DeletedAt:   nil,
			Product:     entity.Product{},
			User:        entity.User{},
		},
		{
			Id:          0,
			Description: "mükemmel",
			Star:        5,
			ProductId:   2,
			UserId:      1,
			Status:      enum.TripActive,
			CreatedAt:   time.Now(),
			UpdatedAt:   nil,
			DeletedAt:   nil,
			Product:     entity.Product{},
			User:        entity.User{},
		},
		{
			Id:          0,
			Description: "idare eder",
			Star:        3,
			ProductId:   1,
			UserId:      1,
			Status:      enum.TripActive,
			CreatedAt:   time.Now(),
			UpdatedAt:   nil,
			DeletedAt:   nil,
			Product:     entity.Product{},
			User:        entity.User{},
		},
		{
			Id:          0,
			Description: "çok güzel",
			Star:        5,
			ProductId:   6,
			UserId:      1,
			Status:      enum.TripActive,
			CreatedAt:   time.Now(),
			UpdatedAt:   nil,
			DeletedAt:   nil,
			Product:     entity.Product{},
			User:        entity.User{},
		},
		{
			Id:          0,
			Description: "mükemmel",
			Star:        5,
			ProductId:   5,
			UserId:      1,
			Status:      enum.TripActive,
			CreatedAt:   time.Now(),
			UpdatedAt:   nil,
			DeletedAt:   nil,
			Product:     entity.Product{},
			User:        entity.User{},
		},
		{
			Id:          0,
			Description: "idare eder",
			Star:        3,
			ProductId:   4,
			UserId:      1,
			Status:      enum.TripActive,
			CreatedAt:   time.Now(),
			UpdatedAt:   nil,
			DeletedAt:   nil,
			Product:     entity.Product{},
			User:        entity.User{},
		},
		{
			Id:          0,
			Description: "çok güzel",
			Star:        5,
			ProductId:   9,
			UserId:      1,
			Status:      enum.TripActive,
			CreatedAt:   time.Now(),
			UpdatedAt:   nil,
			DeletedAt:   nil,
			Product:     entity.Product{},
			User:        entity.User{},
		},
		{
			Id:          0,
			Description: "mükemmel",
			Star:        5,
			ProductId:   8,
			UserId:      1,
			Status:      enum.TripActive,
			CreatedAt:   time.Now(),
			UpdatedAt:   nil,
			DeletedAt:   nil,
			Product:     entity.Product{},
			User:        entity.User{},
		},
		{
			Id:          0,
			Description: "idare eder",
			Star:        3,
			ProductId:   7,
			UserId:      1,
			Status:      enum.TripActive,
			CreatedAt:   time.Now(),
			UpdatedAt:   nil,
			DeletedAt:   nil,
			Product:     entity.Product{},
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
