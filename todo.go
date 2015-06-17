package main 

import "time"

type Todo struct {
	Id          int         `json: "id"`
	Name        string      `json: "name"`
	Completed   bool        `json: "completed"`
	Due         time.Time   `json: "due"`
}

tyoe Todo []Todo