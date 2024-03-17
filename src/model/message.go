package model

type Message struct {
	From    string `json:"from"`
	Content string `json:"content"`
	SentAt string `json:"sent_at,omitempty"`
}
