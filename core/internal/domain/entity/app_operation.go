package entity

import "time"

type AppOperation struct {
	Id                  int32 `json:"key_id" gorm:"primary_key;AUTO_INCREMENT"`
	OperationNumber     string
	OperationId         int8
	Status              int8
	AppliedUserId       int32
	ReceiverUserId      int32
	OperationDate       time.Time
	OperationResultDate time.Time
}
