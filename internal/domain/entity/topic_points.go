package entity

type TopicPoints struct {
	Id         int32    `gorm:"primary_key;AUTO_INCREMENT;column:id"`
	UserId     int32    `gorm:"foreign_key;column:user_id"`
	QuestionId int32    `gorm:"foreign_key;column:question_id"`
	Points     int8     `gorm:"column:points"`
	Status     int8     `gorm:"column:status"`
	User       User     `gorm:"foreign_key:UserId"`
	Question   Question `gorm:"foreign_key:QuestionId"`
}
