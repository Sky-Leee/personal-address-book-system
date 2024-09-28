package main

import (
	"Lab3/Lab3A/model"
	"encoding/json"
	"fmt"
	"net"
	"sync"
)

var (
	contacts = make(map[string]model.Contact)
	mu       sync.Mutex
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	var action string

	// 接收客户端请求
	for {
		err := json.NewDecoder(conn).Decode(&action)
		if err != nil {
			fmt.Println("Error decoding:", err)
			return
		}

		mu.Lock()
		switch action {
		case "view":
			json.NewEncoder(conn).Encode(contacts)
		case "add":
			var contact model.Contact
			json.NewDecoder(conn).Decode(&contact)
			contacts[contact.Name] = contact
		case "update":
			var contact model.Contact
			json.NewDecoder(conn).Decode(&contact)
			contacts[contact.Name] = contact
		case "delete":
			var name string
			json.NewDecoder(conn).Decode(&name)
			delete(contacts, name)
		default:
			fmt.Println("Unknown action:", action)
		}
		mu.Unlock()
	}
}

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer ln.Close()

	fmt.Println("Server listening on port 8080")
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}
