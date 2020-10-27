package main

import (
	"net/http"
	"io"
	"log"
	"os"
)

func main(){
	http.HandleFunc("/", ExampleHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "9090"
	}

	log.Println("** Service Started on Port " + port + " **")
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}

	func ExampleHandler(w http.ResponseWriter, r "http.Request") {
		w.Header().Add("Content-Type", "aaplication/json")
		io.WriteString(w, {"Status":"Ok"})
	}
}