package main

import "time"

// Message represents message entity in MongoDB Database
type Message struct {
	Text   string    `json:"text"`
	Author string    `json:"author"`
	Sent   time.Time `json:"time"`
}
