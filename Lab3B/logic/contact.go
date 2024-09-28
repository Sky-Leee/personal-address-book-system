package logic

import "Lab3/Lab3B/model"

func ViewContacts() string {
	return model.GetAllContacts()
}

func AddContact(name, address, phone string) string {
	if model.ContactExists(name) {
		return "联系人已存在。"
	}
	model.AddContact(name, address, phone)
	return "联系人添加成功。"
}

func ModifyContact(oldName, newName, address, phone string) string {
	if !model.ContactExists(oldName) {
		return "联系人不存在。"
	}
	model.ModifyContact(oldName, newName, address, phone)
	return "联系人修改成功。"
}

func DeleteContact(name string) string {
	if !model.ContactExists(name) {
		return "联系人不存在。"
	}
	model.DeleteContact(name)
	return "联系人删除成功。"
}
