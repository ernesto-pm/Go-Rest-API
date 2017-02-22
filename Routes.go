package main

import "net/http"

// Route defines struct to use for the routes
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes contains all the routes configured for the application
type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"MessageIndex",
		"GET",
		"/messages",
		MessageIndex,
	},
	Route{
		"MessageShow",
		"GET",
		"/messages/{messageID}",
		MessageShow,
	},
	Route{
		"PostMessage",
		"POST",
		"/messages",
		PostMessage,
	},
}
