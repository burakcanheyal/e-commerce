package dto

type TopicPointsDto struct {
	Id         int32 `json:"id"`
	UserId     int32 `json:"user_id"`
	QuestionId int32 `json:"question_id"`
	Points     int8  `json:"points"`
	Status     int8  `json:"status"`
}
