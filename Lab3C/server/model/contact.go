package model

import (
	"sync"
)

var (
	contacts = make(map[string]Contact)
	mu       sync.Mutex
)

type Contact struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
}

func GetAllContacts() []Contact {
	mu.Lock()
	defer mu.Unlock()

	contactList := []Contact{}
	for _, contact := range contacts {
		contactList = append(contactList, contact)
	}
	return contactList
}

func AddContact(contact Contact) string {
	mu.Lock()
	defer mu.Unlock()
	if _, exists := contacts[contact.Name]; exists {
		return "联系人已存在。"
	}
	contacts[contact.Name] = contact
	return "联系人添加成功。"
}

func ModifyContact(contact Contact) string {
	mu.Lock()
	defer mu.Unlock()
	if _, exists := contacts[contact.Name]; !exists {
		return "联系人不存在。"
	}
	contacts[contact.Name] = contact
	return "联系人修改成功。"
}

func DeleteContact(name string) string {
	mu.Lock()
	defer mu.Unlock()
	if _, exists := contacts[name]; !exists {
		return "联系人不存在。"
	}
	delete(contacts, name)
	return "联系人删除成功。"
}
