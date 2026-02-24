package api

type Question struct {
	QuestionText string `json:"questionText"`
	Answer       string `json:"answer"`
	Points       int    `json:"points"`
	Answered     bool   `json:"answered"`
}

type Category struct {
	ID        int64       `json:"id"`
	Name      string      `json:"name"`
	Questions []*Question `json:"questions"`
}

type Board struct {
	ID         string     `json:"id"`
	Categories []Category `json:"categories"`
}
