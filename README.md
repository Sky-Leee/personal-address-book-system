# 通讯录管理系统

本项目是一个简单的通讯录管理系统，分为多个实验（Lab3A、Lab3B、Lab3C），每个实验展示了不同的架构和实现方式。

## 目录结构

```zsh
.
├── Lab3A
│   ├── client
│   │   └── main.go          # 客户端主程序
│   ├── model
│   │   └── contact.go       # 联系人模型
│   └── server
│       └── main.go          # 服务器主程序
├── Lab3B
│   ├── client
│   │   └── main.go          # 客户端主程序
│   ├── logic
│   │   └── contact.go       # 联系人逻辑处理
│   ├── model
│   │   └── db.go            # 数据库操作
│   └── server
│       └── main.go          # 服务器主程序
├── Lab3C
│   ├── client
│   │   ├── add_contact.sh    # 添加联系人脚本
│   │   ├── delete_contact.sh  # 删除联系人脚本
│   │   ├── get_contact.sh     # 查看联系人脚本
│   │   └── modify_contact.sh   # 修改联系人脚本
│   └── server
│       ├── cmd
│       │   └── main.go        # 服务器主程序
│       ├── handler
│       │   └── handler.go      # 请求处理逻辑
│       └── model
│           └── contact.go      # 联系人模型
├── README.md                   # 项目说明文件
└── go.mod                      # Go 模块文件
```

## 实验内容

### Lab3A
- 实现了一个基本的客户端-服务器架构的通讯录管理系统。
- 客户端通过 TCP 协议与服务器进行通信，支持基本的联系人管理功能。

### Lab3B
- 实现了一个三层架构的通讯录管理系统。
- 通过分层处理，增强了系统的可维护性和可扩展性。

### Lab3C
- 使用 B/S 架构实现通讯录管理系统。
- 提供了 RESTful API 接口，使用 `curl` 脚本进行客户端操作。

## 使用说明

1. **环境要求**: 确保已安装 Go 语言环境，并配置好工作区。

2. **运行项目**:
   - 进入对应的实验目录，例如 `Lab3A`：
     ```bash
     cd Lab3A
     ```
   - 启动服务器：
     ```bash
     go run server/main.go
     ```
   - 启动客户端：
     ```bash
     go run client/main.go
     ```

3. **使用脚本 (Lab3C)**:
   - 进入 `Lab3C/client` 目录，使用脚本管理联系人：
     ```bash
     chmod +x add_contact.sh delete_contact.sh get_contact.sh modify_contact.sh
     ./get_contact.sh        # 查看联系人
     ./add_contact.sh        # 添加联系人
     ./modify_contact.sh     # 修改联系人
     ./delete_contact.sh     # 删除联系人
     ```
