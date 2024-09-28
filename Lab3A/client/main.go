package main

import (
	"Lab3/Lab3A/model"
	"encoding/json"
	"fmt"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		return
	}
	defer conn.Close()

	for {
		var choice int
		fmt.Println("选择操作: 1. 查看 2. 添加 3. 修改 4. 删除 5. 退出")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			viewContacts(conn)
		case 2:
			addContact(conn)
		case 3:
			updateContact(conn)
		case 4:
			deleteContact(conn)
		case 5:
			return
		default:
			fmt.Println("无效选择")
		}
	}
}

func viewContacts(conn net.Conn) {
	var action = "view"
	json.NewEncoder(conn).Encode(action)

	var contacts map[string]model.Contact
	json.NewDecoder(conn).Decode(&contacts)

	fmt.Println("联系人信息:")
	for _, contact := range contacts {
		fmt.Printf("姓名: %s, 地址: %s, 电话: %s\n", contact.Name, contact.Address, contact.Phone)
	}
}

func addContact(conn net.Conn) {
	var action = "add"
	json.NewEncoder(conn).Encode(action)

	var contact model.Contact
	fmt.Println("输入姓名:")
	fmt.Scan(&contact.Name)
	fmt.Println("输入地址:")
	fmt.Scan(&contact.Address)
	fmt.Println("输入电话:")
	fmt.Scan(&contact.Phone)

	json.NewEncoder(conn).Encode(contact)
}

func updateContact(conn net.Conn) {
	var action = "update"
	json.NewEncoder(conn).Encode(action)

	var contact model.Contact
	fmt.Println("输入姓名:")
	fmt.Scan(&contact.Name)
	fmt.Println("输入新的地址:")
	fmt.Scan(&contact.Address)
	fmt.Println("输入新的电话:")
	fmt.Scan(&contact.Phone)

	json.NewEncoder(conn).Encode(contact)
}

func deleteContact(conn net.Conn) {
	var action = "delete"
	json.NewEncoder(conn).Encode(action)

	var name string
	fmt.Println("输入要删除的联系人姓名:")
	fmt.Scan(&name)

	json.NewEncoder(conn).Encode(name)
}
