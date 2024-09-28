#!/bin/bash

# 使用 JSON 格式传递新的联系人信息
read -p "请输入要修改的姓名: " name
read -p "请输入新的住址: " address
read -p "请输入新的电话: " phone

curl -X POST http://localhost:8080/modify \
    -H "Content-Type: application/json" \
    -d "{\"name\":\"$name\", \"address\":\"$address\", \"phone\":\"$phone\"}"
