package main

import (
	"net/http"		
	r "getswift/routing"
	"github.com/gorilla/handlers"
)

func main() {

	// r.Test2()
	
	router := r.NewRouter()
	headersOk := handlers.AllowedHeaders([]string{"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})

	http.ListenAndServe(":7070", handlers.CORS(originsOk, headersOk, methodsOk)(router))
	//start_server();


}

