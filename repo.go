// This is a BAD substitute for a Database
// It is just for quick tesing. 

package main 

import "fmt"

var currentId int

var todos Todos

// Initial Data
func init() {
	RepoCreateTodo( Todo{ Name: "Write presentation" } )
	RepoCreateTodo( Todo{ Name: "Host Meeting" } )
}


func RepoFindTodo( id int ) Todo {
	for _, t := range todos {
		if t.Id == id {
			return t
		}
	}
	// returns empty 'Todo' if none found
	return Todo{}
}


func RepoCreateTodo( t Todo ) Todo {
	currentId += 1
	t.Id = currentId
	todos = append( todos, t )
	return t
}


func RepoDestoryTodo( id int ) error {
	for i, t := range todos {
		if t.Id == id {
			todos = append( todos[:i], todos[i+1:]... )
			return nil
		}
	}
	return fmt.Errorf( "Could not find and delete the Todo of id: %d", id )
}