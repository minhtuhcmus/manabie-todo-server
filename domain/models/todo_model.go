package models

import "time"

type Todo struct {
	ID        int       `json:id`
	Title     string    `json:"title"`
	Text      string    `json:"text"`
	Done      string    `json:"done"`
	FkUser    string    `json:"fkUser"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
