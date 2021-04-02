package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type interview struct {
	Message   string
	Timestamp int
}

func main() {
	request()
}

func AvitaInterview(response http.ResponseWriter, r *http.Request) {
	interviewtype := []interview{
		interview{
			Message:   "Avita Pharmacy Interview",
			Timestamp: 04022021,
		},
	}

	json.NewEncoder(response).Encode(interviewtype)

	fmt.Println("Endpoint", interviewtype)
}

func request() {
	http.HandleFunc("/", AvitaInterview)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
