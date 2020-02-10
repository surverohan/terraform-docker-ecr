package main

/* Created by Rohan Surve on 02/08/2020. Application  book and info pojo */

type book struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Autor       string `json:"autor"`
}

type Books []book

type Info struct {
	Path        string `json:"path"`
	Description string `json:"description"`
}
type InfoAPI []Info
