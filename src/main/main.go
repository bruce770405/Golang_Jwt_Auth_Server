package main

import (
	"data"
	"db"
	"github.com/codegangsta/negroni"
	"log"
	"login"
	"net/http"
)

type Predicate = func(int) bool

//type TestInterface interface {
//	save(student Student)
//}

func startServer() {
	// connection mongoDB
	db.Connection()

	// register
	http.HandleFunc("/login", login.Handler)

	http.Handle("/resource", negroni.New(
		negroni.HandlerFunc(login.ValidateTokenMiddleware),
		negroni.Wrap(http.HandlerFunc(data.ProtectedHandler)),
	))

	log.Println("Now listening...")
	http.ListenAndServe(":8080", nil)
}

func main() {
	startServer()
}
