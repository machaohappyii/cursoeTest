# Gin User API

基于 Gin 框架的用户管理 API，支持 MySQL 数据库。

## 功能特性

- RESTful API 设计
- MySQL 数据库支持
- 用户 CRUD 操作
- 输入验证
- 统一错误处理
- 中间件支持 (CORS, 日志, 错误恢复)
- 环境变量配置

## 项目结构

```
gin-user-api/
├── main.go                    # 应用入口
├── config/
│   ├── config.go             # 配置管理
│   └── database.go           # 数据库配置
├── controllers/
│   └── user_controller.go    # 用户控制器
├── middleware/
│   └── common.go             # 通用中间件
├── models/
│   └── user.go               # 用户模型
├── routes/
│   └── router.go             # 路由配置
├── services/
│   └── user_service.go       # 用户业务逻辑
├── utils/
│   └── response.go           # 统一响应工具
└── .env.example              # 环境变量示例
```

## 安装和运行

### 1. 克隆项目
```bash
git clone <repository-url>
cd gin-user-api
```

### 2. 安装依赖
```bash
go mod tidy
```

### 3. 配置数据库
确保 MySQL 服务运行，并创建数据库：
```sql
CREATE DATABASE gin_user_api;
```

### 4. 配置环境变量
复制 `.env.example` 到 `.env` 并修改配置：
```bash
cp .env.example .env
```

编辑 `.env` 文件，设置数据库连接信息：
```
DB_HOST=localhost
DB_PORT=3306
DB_USERNAME=root
DB_PASSWORD=your_password
DB_NAME=gin_user_api
```

### 5. 运行应用
```bash
go run main.go
```

服务器将在 `http://localhost:8080` 启动。

## API 端点

### 健康检查
- `GET /health` - 服务状态检查

### 用户管理
- `POST /api/v1/users` - 创建用户
- `GET /api/v1/users` - 获取所有用户
- `GET /api/v1/users/:id` - 获取单个用户
- `PUT /api/v1/users/:id` - 更新用户
- `DELETE /api/v1/users/:id` - 删除用户

### 请求示例

#### 创建用户
```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "张三",
    "email": "zhangsan@example.com",
    "age": 25
  }'
```

#### 获取所有用户
```bash
curl http://localhost:8080/api/v1/users
```

#### 获取单个用户
```bash
curl http://localhost:8080/api/v1/users/1
```

#### 更新用户
```bash
curl -X PUT http://localhost:8080/api/v1/users/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "李四",
    "email": "lisi@example.com",
    "age": 30
  }'
```

#### 删除用户
```bash
curl -X DELETE http://localhost:8080/api/v1/users/1
```

## 环境变量

| 变量名 | 默认值 | 说明 |
|--------|--------|------|
| PORT | 8080 | 服务器端口 |
| GIN_MODE | debug | Gin 运行模式 |
| DB_DRIVER | mysql | 数据库驱动 |
| DB_HOST | localhost | 数据库主机 |
| DB_PORT | 3306 | 数据库端口 |
| DB_USERNAME | root | 数据库用户名 |
| DB_PASSWORD | | 数据库密码 |
| DB_NAME | gin_user_api | 数据库名称 |
| DB_CHARSET | utf8mb4 | 数据库字符集 |

## 数据库表结构

### users 表
| 字段 | 类型 | 说明 |
|------|------|------|
| id | uint | 主键，自增 |
| name | varchar | 用户姓名 |
| email | varchar | 用户邮箱（唯一） |
| age | int | 用户年龄 |
| created_at | datetime | 创建时间 |
| updated_at | datetime | 更新时间 |