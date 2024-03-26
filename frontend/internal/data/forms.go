package data

type CreatePollForm struct {
	Title     string   `form:"title"`
	Questions []string `form:"inputAnswer"`
}

type AnswerPollForm struct {
	Answer int `form:"radio"`
}
