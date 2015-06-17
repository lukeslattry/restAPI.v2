package main 

import (
	"log"
	"net/http"
	"time"
)

func Logger( inner http.Handler, name string ) http.Handler {
	return http.HandlerFunc( func( w http.ResponseWriter, r *http.Request ) {	
		
		start := time.Now()
		inner.ServeHTTP( w, r )
		
		log.PrintF(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURL,
			name,
			time.Since( start ),
		)
	})
}