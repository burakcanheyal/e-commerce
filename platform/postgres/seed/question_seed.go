package seed

import (
	"attempt4/internal/domain/entity"
	"attempt4/internal/domain/enum"
	"gorm.io/gorm"
)

func QuestionSeed(db *gorm.DB) {
	questions := []entity.Question{
		{
			Id:          0,
			Description: "Do visitors often engage in outdoor activities along rivers or in historic districts, such as picnics or leisurely walks?",
			Status:      enum.QuestionActive,
			Topic:       enum.QuestionNature,
		},
		{
			Id:          0,
			Description: "Is it possible to find secluded spots for relaxation amidst nature within historic areas?",
			Status:      enum.QuestionActive,
			Topic:       enum.QuestionNature,
		},
		{
			Id:          0,
			Description: "Are there guided tours available for visitors interested in learning more about the historical significance of local landmarks?",
			Status:      enum.QuestionActive,
			Topic:       enum.QuestionHistorical,
		},
		{
			Id:          0,
			Description: "Can visitors expect to see interactive exhibits or multimedia installations at museums that bring the area's history to life?",
			Status:      enum.QuestionActive,
			Topic:       enum.QuestionHistorical,
		},
		{
			Id:          0,
			Description: "Are there any organized adventure tours offered for exploring natural attractions, such as rivers or mountains?",
			Status:      enum.QuestionActive,
			Topic:       enum.QuestionAdventure,
		},
		{
			Id:          0,
			Description: "Is it common for travelers to discover hidden alleys or secret pathways while exploring historic districts?",
			Status:      enum.QuestionActive,
			Topic:       enum.QuestionAdventure,
		},
		{
			Id:          0,
			Description: "Are there frequent cultural performances or live music events held within historic areas?",
			Status:      enum.QuestionActive,
			Topic:       enum.QuestionCultural,
		},
		{
			Id:          0,
			Description: "Can visitors easily find local artisans selling traditional crafts or handmade souvenirs in cultural hubs?",
			Status:      enum.QuestionActive,
			Topic:       enum.QuestionCultural,
		},
		{
			Id:          0,
			Description: "Do many visitors choose to unwind by natural features like rivers or in historic districts, enjoying the tranquility of their surroundings?",
			Status:      enum.QuestionActive,
			Topic:       enum.QuestionRelaxation,
		},
		{
			Id:          0,
			Description: "Are there any designated relaxation spots within historic areas, such as parks or quiet squares, where tourists can take a break from sightseeing?",
			Status:      enum.QuestionActive,
			Topic:       enum.QuestionRelaxation,
		},
	}

	var size int64
	db.Model(&questions).Count(&size)
	if size == 0 {
		for _, p := range questions {
			db.Create(&p)
		}
	}
}
