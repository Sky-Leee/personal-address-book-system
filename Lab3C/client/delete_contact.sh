#!/bin/bash

read -p "请输入要删除的姓名: " name

curl -X DELETE "http://localhost:8080/delete?name=$name"
