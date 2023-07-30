package it_test

import (
	"fmt"
	"io"
	"net/http"
	"log"
)

func GameServer(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://localhost:3333/game")
	if err != nil {
		log.Fatal("Error getting response from localhost. Full error: " + err.Error())
		fmt.Fprintf(w, "err")
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(w, "err")
	}
	fmt.Fprint(w, string(body))
}
