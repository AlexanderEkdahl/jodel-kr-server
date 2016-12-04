package main

import "net/http"

// Route ...
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes ...
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
	Route{
		"HandlerPostComment",
		"POST",
		"/post_comment",
		HandlerPostComment,
	},
	Route{
		"HandlerGetMessagesWithUser",
		"GET",
		"/get_user",
		HandlerGetMessagesWithUser,
	},
}
