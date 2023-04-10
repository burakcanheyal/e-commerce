package dto

type AppOperationDto struct {
	Id              int32  `json:"id"`
	UserId          int32  `json:"user_id"`
	OperationNumber string `json:"operation_number"`
	Response        int8   `json:"response"`
}
