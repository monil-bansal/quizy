package model

// Create struct to hold info about new quiz
type Question struct {
	Question string   `json:"question"`
	Options  []string `json:"options"`
	Answer   int      `json:"answer"`
}

type Quiz struct {
	Title     string     `json:"title"`
	Questions []Question `json:"questions"`
}
