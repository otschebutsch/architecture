package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	type CurrentT struct {
		CurrentT string `json:"time"`
	}
	http.HandleFunc("/time", func(rw http.ResponseWriter, request *http.Request) {
		time := &CurrentT{CurrentT: time.Now().Format(time.RFC3339)}
		timej, _ := json.Marshal(time)
		rw.Header().Set("content-type", "aplication/json")
		if _, err := rw.Write([]byte(timej)); err != nil {
			fmt.Printf("Error writing response to the client: %s", err)
		}
	})
	log.Fatal(http.ListenAndServe(":8795", nil))
}
