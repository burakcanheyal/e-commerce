package dto

type PanelDto struct {
	Id       int32 `json:"id"`
	UserId   int32 `json:"user_id"`
	Response int8  `json:"response"`
}
