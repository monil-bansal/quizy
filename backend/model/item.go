package model

// Create struct to hold info about new quiz
type Question struct {
	Question string   `json:"question"`
	Options  []string `json:"options"`
	Answer   int      `json:"answer"`
}

type Quiz struct {
	QuizId    string     `dynamodbav:"quizId"`
	Title     string     `json:"title" dynamodbav:"title"`
	Questions []Question `json:"questions" dynamodbav:"questions"`
}
