package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

// Index function returns the query to the index URL
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the messages REST API")
}

// MessageIndex function to render the messages
func MessageIndex(w http.ResponseWriter, r *http.Request) {
	messages := getMessages()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(messages); err != nil {
		panic(err)
	}
}

func PostMessage(w http.ResponseWriter, r *http.Request) {
	var message Message
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &message); err != nil {
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(422)
		if err2 := json.NewEncoder(w).Encode(err); err2 != nil {
			panic(err2)
		}
	}

	if message.Author != "" && message.Text != "" {
		createMessage(message)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode("Message created"); err != nil {
			panic(err)
		}
	} else {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusConflict)
		if err := json.NewEncoder(w).Encode("Please input message body"); err != nil {
			panic(err)
		}
	}

}

// MessageShow function to show specific functions
func MessageShow(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	//messageID := vars["messageId"]
	//fmt.Fprintln(w, "Message shown:", messageId)
}
