package entity

import "time"

type Submission struct {
	Id                  int32      `gorm:"primary_key;AUTO_INCREMENT;column:id"`
	SubmissionNumber    string     `gorm:"type:varchar(16);column:submission_number"`
	SubmissionType      int8       `gorm:"type:smallint;column:submission_type"`
	Status              int8       `gorm:"type:smallint;column:status"`
	AppliedUserId       int32      `gorm:"column:applied_user_id"`
	ReceiverUserId      *int32     `gorm:"column:receiver_user_id"`
	OperationDate       time.Time  `gorm:"column:operation_date"`
	OperationResultDate *time.Time `gorm:"column:operation_result_date"`
}
