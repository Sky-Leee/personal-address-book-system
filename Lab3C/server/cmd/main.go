package main

import (
	"Lab3/Lab3C/server/handler"
	"net/http"
)

func main() {
	http.HandleFunc("/contacts", handler.ContactsHandler)
	http.HandleFunc("/add", handler.AddContactHandler)
	http.HandleFunc("/modify", handler.ModifyContactHandler)
	http.HandleFunc("/delete", handler.DeleteContactHandler)

	http.ListenAndServe(":8080", nil)
}
