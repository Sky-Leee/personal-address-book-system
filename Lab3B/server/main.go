package main

import (
	"Lab3/Lab3B/logic"
	"fmt"
	"net"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error starting server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("服务器启动，等待连接...")
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	for {
		command := make([]byte, 1024)
		n, err := conn.Read(command)
		if err != nil {
			return
		}
		input := strings.TrimSpace(string(command[:n]))
		parts := strings.Fields(input)

		if len(parts) == 0 {
			continue
		}

		var response string
		switch parts[0] {
		case "VIEW":
			response = logic.ViewContacts()
		case "ADD":
			if len(parts) != 4 {
				response = "参数不正确。"
				break
			}
			response = logic.AddContact(parts[1], parts[2], parts[3])
		case "MODIFY":
			if len(parts) != 5 {
				response = "参数不正确。"
				break
			}
			response = logic.ModifyContact(parts[1], parts[2], parts[3], parts[4])
		case "DELETE":
			if len(parts) != 2 {
				response = "参数不正确。"
				break
			}
			response = logic.DeleteContact(parts[1])
		default:
			response = "无效命令"
		}

		conn.Write([]byte(response + "\n"))
	}
}
