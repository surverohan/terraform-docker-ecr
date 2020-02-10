package main

import (
	"log"
	"net/http"
)

/* Created by Rohan Surve on 02/08/2020. Application  bootstraping with listener configuration */

func main() {
	
	router := AppRouter()
	log.Fatal(http.ListenAndServe(":8080", router))

}
