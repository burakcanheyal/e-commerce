package dto

type FeedbackDto struct {
	Id          int32  `json:"id"`
	Description string `json:"description"`
	Star        int32  `json:"star"`
	TripId      int32  `json:"trip_id"`
	UserId      int32  `json:"user_id"`
	Status      int8   `json:"status"`
	UserName    string `json:"user_name"`
}
