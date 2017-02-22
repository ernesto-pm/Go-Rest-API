package main

import (
	"log"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func getMessages() []Message {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	c := session.DB("realtimeChat").C("messages")
	var result []Message
	err = c.Find(bson.M{}).All(&result)
	if err != nil {
		log.Fatal("error", err)
	}
	return result
}

func createMessage(message Message) {
	session, err := mgo.Dial("localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	c := session.DB("realtimeChat").C("messages")
	err = c.Insert(&Message{message.Text, message.Author, time.Now()})
	if err != nil {
		log.Fatal(err)
	}
}

/*
	err = c.Insert(&Message{"Hello?", "KetziaDB", time.Now()})
	if err != nil {
		log.Fatal(err)
	}
*/
