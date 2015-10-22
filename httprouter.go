package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Username struct {
	Name string `json:"name"`
}

type Message struct {
	Greet string `json:"greeting"`
}

func hello(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
	fmt.Fprintf(rw, "Hello, %s!\n", p.ByName("name"))
}

func fooHandler(rw http.ResponseWriter, req *http.Request, p httprouter.Params) {
	var usr Username
	json.NewDecoder(req.Body).Decode(&usr)
	var mess Message
	mess.Greet = "Hello, " + usr.Name + "!"
	j, _ := json.Marshal(mess)

	rw.Header().Set("Content-Type", "application/json")

	fmt.Fprintf(rw, "%s", j)
}

func main() {
	mux := httprouter.New()
	mux.GET("/hello/:name", hello)
	mux.POST("/hello", fooHandler)
	server := http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
