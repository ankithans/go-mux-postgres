package main

import (
	routes "./routes"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	const port string = ":8000"

	router.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Server is Up and Running...")
	})
	router.HandleFunc("/posts", routes.GetPosts).Methods("GET")
	router.HandleFunc("/posts", routes.AddPost).Methods("POST")

	log.Println("Server Listening on port", port)
	log.Fatal(http.ListenAndServe(port, router))
}
