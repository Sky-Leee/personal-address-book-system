package handler

import (
	"Lab3/Lab3C/server/model"
	"encoding/json"
	"net/http"
)

func ContactsHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodGet {
        contacts := model.GetAllContacts()
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(contacts)
    } else {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

func AddContactHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        var contact model.Contact
        if err := json.NewDecoder(r.Body).Decode(&contact); err != nil {
            http.Error(w, "Invalid request payload", http.StatusBadRequest)
            return
        }
        msg := model.AddContact(contact)
        w.Write([]byte(msg))
    } else {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

func ModifyContactHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        var contact model.Contact
        if err := json.NewDecoder(r.Body).Decode(&contact); err != nil {
            http.Error(w, "Invalid request payload", http.StatusBadRequest)
            return
        }
        msg := model.ModifyContact(contact)
        w.Write([]byte(msg))
    } else {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}

func DeleteContactHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodDelete {
        name := r.URL.Query().Get("name")
        msg := model.DeleteContact(name)
        w.Write([]byte(msg))
    } else {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
    }
}
