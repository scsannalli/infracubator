package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func Greet(name string) string {
	return "Hello " + name
}

func main() {
	fmt.Printf("Serving Application !!!!!!! \n")
	http.HandleFunc("/", getRoot)
	http.HandleFunc("/game", getGame)
	err := http.ListenAndServe(":3333", nil)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	if _, err := io.WriteString(w, "Hello World!\n"); err != nil {
		log.Fatal(err)
	}
}

func getGame(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got /game request\n")
	if _, err := io.WriteString(w, "Starting Game\n"); err != nil {
		log.Fatal(err)
	}
}
