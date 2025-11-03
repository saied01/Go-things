package cmd

import "time"

type Task struct {
	ID      int
	Content string
	Done    bool
	Created time.Time
}
