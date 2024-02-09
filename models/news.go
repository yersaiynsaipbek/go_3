package models

type News struct {
	ID         int     `json:"id"`
	Title      string  `json:"title"`
	Content    string  `json:"content"`
	Estimation float32 `json:"estimation"`
	Category   string  `json:"category"`
}
