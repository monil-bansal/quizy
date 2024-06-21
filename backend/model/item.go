package model

// Create struct to hold info about new quiz
type Question struct {
	Question string   `json:"question"`
	Options  []string `json:"options"`
	Answer   int      `json:"answer"`
}

type Quiz struct {
	QuizId    string     `json:"quizId" dynamodbav:"quizId"`
	Title     string     `json:"title" dynamodbav:"title"`
	Questions []Question `json:"questions" dynamodbav:"questions"`
}

type UserSumission struct {
	UserId    string     `json:"userId" dynamodbav:"userId"`
	QuizId    string     `json:"quizId" dynamodbav:"quizId"`
	Score     int        `json:"score" dynamodbav:"score"`
	Title     string     `json:"title" dynamodbav:"title"` // have to store it twice or might have to read it from db : extra storeage is preferable over extra db read but then storage will be forever and db read only when submitting/accessing submission
	Questions []Question `json:"questions" dynamodbav:"questions"`
}
