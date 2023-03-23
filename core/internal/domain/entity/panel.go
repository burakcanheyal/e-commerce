package entity

type Panel struct {
	Id          int32 `json:"key_id" gorm:"primary_key;AUTO_INCREMENT"`
	OperationId int8
	Status      int8
}
