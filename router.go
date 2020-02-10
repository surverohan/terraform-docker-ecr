package main

import (
	"github.com/gorilla/mux"
)

/* Created by Rohan Surve on 02/08/2020. Application Router to route api's call to corrrespoding controller endpoints */

func AppRouter() *mux.Router {

	router := mux.AppRouter().StrictSlash(true)
	for _, route := range routes {

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandlerFunc)

	}
	return router

}
