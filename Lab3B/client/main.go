package main

import (
    "fmt"
    "net"
    "os"
)

func main() {
    conn, err := net.Dial("tcp", "localhost:8080")
    if err != nil {
        fmt.Println("Error connecting:", err)
        os.Exit(1)
    }
    defer conn.Close()

    for {
        fmt.Println("1. 查看联系人")
        fmt.Println("2. 添加联系人")
        fmt.Println("3. 修改联系人")
        fmt.Println("4. 删除联系人")
        fmt.Println("5. 退出")

        var choice int
        fmt.Scan(&choice)

        switch choice {
        case 1:
            viewContacts(conn)
        case 2:
            addContact(conn)
        case 3:
            modifyContact(conn)
        case 4:
            deleteContact(conn)
        case 5:
            return
        default:
            fmt.Println("无效选择，请重新输入。")
        }
    }
}

func viewContacts(conn net.Conn) {
    conn.Write([]byte("VIEW\n"))
    response := make([]byte, 1024)
    n, _ := conn.Read(response)
    fmt.Println("联系人信息:", string(response[:n]))
}

func addContact(conn net.Conn) {
    var name, address, phone string
    fmt.Println("输入姓名:")
    fmt.Scan(&name)
    fmt.Println("输入住址:")
    fmt.Scan(&address)
    fmt.Println("输入电话:")
    fmt.Scan(&phone)

    command := fmt.Sprintf("ADD %s %s %s\n", name, address, phone)
    conn.Write([]byte(command))
    response := make([]byte, 1024)
    n, _ := conn.Read(response)
    fmt.Println(string(response[:n]))
}

func modifyContact(conn net.Conn) {
    var oldName, newName, address, phone string
    fmt.Println("输入要修改的姓名:")
    fmt.Scan(&oldName)
    fmt.Println("输入新姓名:")
    fmt.Scan(&newName)
    fmt.Println("输入新住址:")
    fmt.Scan(&address)
    fmt.Println("输入新电话:")
    fmt.Scan(&phone)

    command := fmt.Sprintf("MODIFY %s %s %s %s\n", oldName, newName, address, phone)
    conn.Write([]byte(command))
    response := make([]byte, 1024)
    n, _ := conn.Read(response)
    fmt.Println(string(response[:n]))
}

func deleteContact(conn net.Conn) {
    var name string
    fmt.Println("输入要删除的姓名:")
    fmt.Scan(&name)

    command := fmt.Sprintf("DELETE %s\n", name)
    conn.Write([]byte(command))
    response := make([]byte, 1024)
    n, _ := conn.Read(response)
    fmt.Println(string(response[:n]))
}
