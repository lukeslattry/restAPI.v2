package main 

import "time"

type Todo strcut {
	Name        string      `json: "name"`
	Completed   bool        `json: "completed"`
	Due         time.Time   `json: "due"`
}

tyoe Todo []Todo