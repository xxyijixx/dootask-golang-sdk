# DooTask Go SDK

## 概述

DooTask Go SDK 是为 DooTask 团队协作平台开发的 Go 语言 SDK。该 SDK 提供了完整的 API 封装，支持用户管理、项目管理、系统配置、对话管理、文件管理、汇报管理、公共接口和审批工作流等核心功能模块。

## 特性

- 🚀 **模块化设计**: 按功能模块组织代码，便于维护和扩展
- 🔒 **类型安全**: 完整的类型定义和数据模型
- 🛠 **易于使用**: 简洁的 API 设计，支持链式调用
- 📊 **调试支持**: 内置请求/响应日志记录
- ⚙️ **可配置**: 支持自定义超时、认证、重试等配置
- 📝 **完整文档**: 详细的使用示例和 API 文档

## 目录结构

```
golang-sdk/
├── README.md              # 项目说明文档
├── go.mod                 # Go 模块定义
├── go.sum                 # Go 依赖校验
├── sdk/
│   ├── client.go          # 核心客户端
│   ├── config.go          # 配置管理
│   └── auth.go            # 认证相关
├── models/                # 数据模型
│   ├── common.go          # 通用模型
│   ├── users.go           # 用户相关模型
│   ├── project.go         # 项目相关模型
│   ├── system.go          # 系统相关模型
│   ├── dialog.go          # 对话相关模型
│   ├── file.go            # 文件相关模型
│   ├── report.go          # 汇报相关模型
│   ├── public.go          # 公共接口模型
│   └── approve.go         # 审批相关模型
├── api/                   # API 接口封装
│   ├── users/             # 用户模块 API
│   ├── project/           # 项目模块 API
│   ├── system/            # 系统模块 API
│   ├── dialog/            # 对话模块 API
│   ├── file/              # 文件模块 API
│   ├── report/            # 汇报模块 API
│   ├── public/            # 公共接口 API
│   └── approve/           # 审批模块 API
└── utils/                 # 工具函数
    ├── http.go            # HTTP 工具
    ├── response.go        # 响应处理
    └── error.go           # 错误处理
```
```
golang-sdk/
├── README.md              # 完整的项目文档和使用指南
├── go.mod/go.sum          # Go模块定义
├── sdk/
│   └── client.go          # 重构的核心客户端
├── models/                # 数据模型 (8个模块)
│   ├── common.go          # 通用模型
│   ├── users.go           # 用户相关模型
│   ├── project.go         # 项目相关模型
│   ├── system.go          # 系统相关模型
│   ├── dialog.go          # 对话相关模型
│   ├── file.go            # 文件相关模型
│   ├── report.go          # 汇报相关模型
│   ├── public.go          # 公共接口模型
│   └── approve.go         # 审批相关模型
├── api/                   # API接口封装 (待实现)
│   ├── users/             
│   ├── project/           
│   ├── system/            
│   ├── dialog/            
│   ├── file/              
│   ├── report/            
│   ├── public/            
│   └── approve/           
└── utils/                 # 工具函数
    └── utils.go           # 增强的工具函数
```

## 安装

```bash
go get github.com/xxyijixx/dootask-golang-sdk
```

## 快速开始

### 基本使用

```go
package main

import (
    "fmt"
    "log"
    "github.com/xxyijixx/dootask-golang-sdk/sdk"
    "github.com/xxyijixx/dootask-golang-sdk/api/users"
)

func main() {
    // 创建客户端
    client := sdk.NewClient("https://your-dootask-instance.com/api")

    // 设置认证令牌
    client.SetToken("your-auth-token")

    // 创建用户 API 客户端
    userClient := users.NewClient(client)

    // 获取用户信息
    user, err := userClient.Info()
    if err != nil {
        log.Fatal("获取用户信息失败:", err)
    }

    fmt.Printf("用户: %s (%s)\n", user.Name, user.Email)
}
```

### 高级配置

```go
package main

import (
    "time"
    "github.com/xxyijixx/dootask-golang-sdk/sdk"
)

func main() {
    // 自定义配置
    config := &sdk.Config{
        Timeout:   30 * time.Second,
        UserAgent: "MyApp/1.0",
        Insecure:  false, // 生产环境请设为 false
        Debug:     true,  // 开发环境开启调试
    }

    // 创建客户端
    client := sdk.NewClientWithConfig("https://your-dootask-instance.com/api", config)
    client.SetToken("your-auth-token")

    // 使用客户端...
}
```

## 模块说明

| 模块      | 说明           | 状态 | 优先级 |
|-----------|----------------|------|--------|
| users     | 用户管理       | 🔄 开发中 | ⭐⭐⭐ |
| project   | 项目管理       | ⏳ 待开发 | ⭐⭐⭐ |
| system    | 系统配置       | ⏳ 待开发 | ⭐⭐⭐ |
| dialog    | 对话管理       | ⏳ 待开发 | ⭐⭐  |
| file      | 文件管理       | ⏳ 待开发 | ⭐⭐  |
| report    | 汇报管理       | ⏳ 待开发 | ⭐⭐  |
| public    | 公共接口       | ⏳ 待开发 | ⭐    |
| approve   | 审批工作流     | ⏳ 待开发 | ⭐    |

## API 文档

### 客户端配置

#### NewClient(baseURL string) *Client
创建基础客户端实例

#### NewClientWithConfig(baseURL string, config *Config) *Client
创建带有自定义配置的客户端实例

#### SetToken(token string)
设置认证令牌

### 用户模块 (Users API)

#### 用户认证
- `Login(request *LoginRequest) (*LoginResponse, error)` - 用户登录
- `Logout() error` - 用户登出
- `RefreshToken() (*TokenResponse, error)` - 刷新令牌

#### 用户信息
- `Info() (*UserProfile, error)` - 获取当前用户信息
- `Update(request *UserUpdateRequest) (*User, error)` - 更新用户信息
- `ChangePassword(request *PasswordChangeRequest) error` - 修改密码

#### 用户管理 (管理员)
- `List(request *UserListRequest) (*UserListResponse, error)` - 获取用户列表
- `Create(request *UserCreateRequest) (*User, error)` - 创建用户
- `Delete(userID uint) error` - 删除用户

#### 部门管理
- `DepartmentList() ([]Department, error)` - 获取部门列表
- `DepartmentCreate(request *DepartmentCreateRequest) (*Department, error)` - 创建部门

### 项目模块 (Project API)

#### 项目管理
- `List(request *ProjectListRequest) (*ProjectListResponse, error)` - 获取项目列表
- `Create(request *ProjectCreateRequest) (*Project, error)` - 创建项目
- `Update(projectID uint, request *ProjectUpdateRequest) (*Project, error)` - 更新项目
- `Delete(projectID uint) error` - 删除项目

#### 项目任务
- `TaskList(projectID uint, request *TaskListRequest) (*TaskListResponse, error)` - 获取任务列表
- `TaskCreate(request *TaskCreateRequest) (*Task, error)` - 创建任务
- `TaskUpdate(taskID uint, request *TaskUpdateRequest) (*Task, error)` - 更新任务

### 系统模块 (System API)

#### 系统设置
- `Settings() (*SystemSettings, error)` - 获取系统设置
- `UpdateSettings(request *SystemSettingsUpdateRequest) error` - 更新系统设置

#### 系统信息
- `Info() (*SystemInfo, error)` - 获取系统信息
- `Version() (*VersionInfo, error)` - 获取版本信息

## 数据模型

### 通用模型
- `APIResponse` - 标准 API 响应
- `BaseModel` - 基础模型 (包含 ID、创建时间、更新时间)
- `PaginationParams` - 分页参数
- `PaginatedResponse` - 分页响应

### 用户模型
- `User` - 用户信息
- `UserProfile` - 用户详细信息
- `Department` - 部门信息
- `LoginRequest/Response` - 登录相关

### 项目模型
- `Project` - 项目信息
- `Task` - 任务信息
- `ProjectUser` - 项目成员关系

## 错误处理

SDK 提供了统一的错误处理机制：

```go
user, err := userClient.Info()
if err != nil {
    switch e := err.(type) {
    case *utils.APIError:
        fmt.Printf("API 错误 [%d]: %s\n", e.Ret, e.Msg)
    default:
        fmt.Printf("其他错误: %s\n", err.Error())
    }
    return
}
```

## 调试模式

开启调试模式可以查看请求和响应详情：

```go
config := &sdk.Config{
    Debug: true,
    // ... 其他配置
}
client := sdk.NewClientWithConfig(baseURL, config)
```

## 贡献指南

欢迎提交 Issue 和 Pull Request！

1. Fork 本项目
2. 创建特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 创建 Pull Request

## 许可证

本项目采用 MIT 许可证 - 查看 [LICENSE](../LICENSE) 文件了解详情