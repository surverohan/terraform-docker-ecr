package main

import (
	"net/http"
)

/* Created by Rohan Surve on 02/08/2020. Appication Router routing configurations */

type Route struct {
	Name        string
	Method      string
	Pattern     string
	Description string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes Routes

func init() {
	routes = Routes{
		
		Route{
			"TodoShow",
			"GET",
			"/book/{bookId}",
			"Get specific book with ID",
			FetchBook,
		},
		Route{
			"TodoIndex",
			"GET",
			"/fetchallbooks",
			"Get all Books",
			FetchAllBooks,
		},

		Route{
			"info",
			"GET",
			"/info",
			"Get Paths info",
			AppInfo,
		},

		Route{
			"Index",
			"GET",
			"/",
			"Index",
			Index,
		},
	}
}
