package model

import (
    "fmt"
    "sync"
)

var (
    contacts = make(map[string]Contact)
    mu       sync.Mutex
)

type Contact struct {
    Name    string
    Address string
    Phone   string
}

func GetAllContacts() string {
    mu.Lock()
    defer mu.Unlock()

    var result string
    for _, contact := range contacts {
        result += fmt.Sprintf("姓名: %s, 住址: %s, 电话: %s\n", contact.Name, contact.Address, contact.Phone)
    }
    if result == "" {
        return "没有联系人信息。"
    }
    return result
}

func AddContact(name, address, phone string) {
    mu.Lock()
    defer mu.Unlock()
    contacts[name] = Contact{Name: name, Address: address, Phone: phone}
}

func ModifyContact(oldName, newName, address, phone string) {
    mu.Lock()
    defer mu.Unlock()
    contact := contacts[oldName]
    delete(contacts, oldName)
    contact.Name = newName
    contact.Address = address
    contact.Phone = phone
    contacts[newName] = contact
}

func DeleteContact(name string) {
    mu.Lock()
    defer mu.Unlock()
    delete(contacts, name)
}

func ContactExists(name string) bool {
    mu.Lock()
    defer mu.Unlock()
    _, exists := contacts[name]
    return exists
}
