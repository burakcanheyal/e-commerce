package entity

import "time"

type Submission struct {
	Id                  int32     `gorm:"primary_key;AUTO_INCREMENT;column:id"`
	SubmissionNumber    string    `gorm:"column:operation_number"`
	SubmissionType      int8      `gorm:"column:operation_type"`
	Status              int8      `gorm:"column:status"`
	AppliedUserId       int32     `gorm:"column:applied_user_id"`
	ReceiverUserId      int32     `gorm:"column:receiver_user_id"`
	OperationDate       time.Time `gorm:"column:operation_date"`
	OperationResultDate time.Time `gorm:"column:operation_result_date"`
}
