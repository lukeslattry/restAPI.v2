package main 

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io"
	"io/ioutil"
	"strconv"
	"github.com/gorilla/mux"
)


func Index( w, http.ResponseWriter, r *http.Request ) {
	fmt.Fprintln( w, "Welcome!" )
}


func TodoIndex( w http.ResponseWriter, r *http.Request ) {
//	todo := Todos {
//		Todo{ Name: "Write Presentation" },
//		Todo{ Name: "Host Meeting" },
//	}

	w.Header().Set( "Content-Type", "application/json; charset=UTF-8" )
	w.WriteHeader( http.StatusOK )

	if err := json.NewEncoder( w ).Encode( todos ); err != nil {
		panic( err )
	}
}


func TodoShow( w http.ResponseWriter, r *http.Request ) {
	vars := mux.Vars( r )
	var todoId int
	var err error

	if todoId, err = strconv.Atoi( vars["todoId"] ); err != nil {
		panic( err )
	}

	todo := RepoFindTodo( todoId )
	
	if todo.Id > 0 {
		w.Header().Set( "Content-Type", "application/json; charset=UTF-8" )
		w.WriteHeader( http.StatusOK )

		if err := json.NewEncoder( w ).Encode( todo ); err != nil {
			panic( err )
		}
		return
	}

	// 404 that shit
	w.Header().Set( "Content-Type", "application/json; charset=UTF-8" )
	w.WriteHeader( http.StatusNotFound )

	if err := json.NewEncoder( w ).Encode( jsonErr{ Code: http.StatusNotFound, Text: "Not Found" } ); err != nil {
		panic( err )
	}
}


func TodoCreate( w http.ResponseWriter, r *http.Request ) {
	var todo Todo
	body, err := ioutil.ReadAll( io.LimitReader( r.Body, 1048576 ) )

	if err := r.Body.Close(); err != nil {
		panic( err )
	}

	if err := json.Unmarshal( body, &todo ); err != nil {
		w.Header().Set( "Content-Type", "application/json; charset=UTF-8" )
		w.WriteHeader( 442 ) //Unprocessable entity

		if er := json.NewEncoder( w ).Encode( err ); err != nil {
			panic( err )
		} 
	}

	t := RepoCreateTodo( todo )
	w.Header().Set( "Content-Type", "application/json; charset=UTF-8" )
	w.WriteHeader( http.StatusCreated )

	if err := json.NewEncoder( w ).Encode( t ); err != nil {
		panic( err )
	}
}