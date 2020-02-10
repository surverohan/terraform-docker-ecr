package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

/* Created by Rohan Surve on 02/09/2020. Application  controller implementation */

func AppInfo(w http.ResponseWriter, r *http.Request) {
	var infoAPI InfoAPI
	for _, values := range routes {

		infoAPI = append(infoAPI,
			Info{Path: values.Pattern, Description: values.Description})

	}

	json.NewEncoder(w).Encode(infoAPI)
}

func FetchAllBooks(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(books)
}

func FetchBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId, err := strconv.Atoi(vars["bookId"])
	if err == nil {

		for _, values := range books {
			if bookId == values.ID {

				json.NewEncoder(w).Encode(values)
			}

		}
	}
}


func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}


