package cmd

import "time"

type Task struct {
	ID      int       `json:"id"`
	Content string    `json:"content"`
	Done    bool      `json:"done"`
	Created time.Time `json:"created"`
}
