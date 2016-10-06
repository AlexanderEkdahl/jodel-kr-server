package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"HandlerGetMessages",
		"GET",
		"/get",
		HandlerGetMessages,
	},
	Route{
		"HandlerPostMessage",
		"POST",
		"/post",
		HandlerPostMessage,
	},
}
