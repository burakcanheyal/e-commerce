package dto

type QuestionDto struct {
	Id              int32  `json:"id"`
	Description     string `json:"description"`
	Status          int8   `json:"status"`
	Topic           int8   `json:"topic"`
	AnswerQuestions int8   `json:"answer_questions"`
}
