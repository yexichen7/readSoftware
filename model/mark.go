package model

type Mark struct {
	ID      int    `json:"userID"`
	BookID  int    `json:"book_id"`
	Name    string `json:"name"`
	Page    int    `json:"page"`
	Content string `json:"content"`
}
