#!/bin/bash

# 使用 JSON 格式传递联系人信息
read -p "请输入姓名: " name
read -p "请输入住址: " address
read -p "请输入电话: " phone

curl -X POST http://localhost:8080/add \
    -H "Content-Type: application/json" \
    -d "{\"name\":\"$name\", \"address\":\"$address\", \"phone\":\"$phone\"}"
