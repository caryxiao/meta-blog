# Meta Blog - Go语言博客系统

一个基于Go语言和Gin框架开发的现代化博客系统，支持用户管理、文章发布、评论互动等功能，并集成了完整的操作日志记录系统。

## 项目特性

- 🚀 基于Gin框架的高性能Web服务
- 🔐 JWT身份认证和授权
- 📝 完整的博客CRUD操作
- 💬 评论系统
- 📊 操作日志记录
- 🗄️ MySQL数据库支持
- 🔧 数据库迁移管理
- 📦 依赖注入容器
- 🎯 RESTful API设计

## 项目结构

```
meta-blog/
├── cmd/                    # 应用程序入口
│   ├── blog/              # 博客服务主程序
│   │   └── main.go
│   └── migrate/           # 数据库迁移工具
│       └── main.go
├── configs/               # 配置文件
│   ├── config.dev.yaml   # 开发环境配置
│   └── config.yaml       # 默认配置
├── db/                    # 数据库相关
│   ├── db.go             # 数据库连接
│   └── migrations/       # 数据库迁移文件
│       ├── 000001_init_schema.up.sql
│       ├── 000001_init_schema.down.sql
│       ├── 000002_create_logs_table.up.sql
│       └── 000002_create_logs_table.down.sql
├── internal/              # 内部包
│   ├── config/           # 配置管理
│   ├── di/               # 依赖注入
│   ├── handler/          # HTTP处理器
│   ├── middleware/       # 中间件
│   ├── model/            # 数据模型
│   ├── repository/       # 数据访问层
│   ├── request/          # 请求结构体
│   ├── response/         # 响应结构体
│   ├── router/           # 路由配置
│   ├── service/          # 业务逻辑层
│   └── utils/            # 工具函数
├── go.mod                # Go模块文件
├── go.sum                # 依赖校验文件
└── README.md             # 项目说明文档
```

## API路由

### 用户相关
- `POST /api/register` - 用户注册
- `POST /api/login` - 用户登录

### 文章相关
- `GET /api/posts` - 获取文章列表
- `POST /api/posts` - 创建文章 (需要认证)
- `GET /api/posts/:id` - 获取文章详情
- `PUT /api/posts/:id` - 更新文章 (需要认证)
- `DELETE /api/posts/:id` - 删除文章 (需要认证)

### 评论相关
- `POST /api/posts/:id/comments` - 为指定文章创建评论 (需要认证)
- `GET /api/posts/:id/comments` - 获取指定文章的评论列表

## 核心功能

### 1. 用户管理
- 用户注册和登录
- 密码加密存储
- JWT令牌认证
- 用户信息管理

### 2. 文章管理
- 文章的增删改查
- 文章列表分页
- 用户权限控制（只能操作自己的文章）
- 文章状态管理

### 3. 评论系统
- 文章评论功能
- 评论列表分页
- 评论权限控制

### 4. 操作日志
- 用户注册日志
- 文章操作日志（创建、更新、删除）
- 评论操作日志
- 详细的操作信息记录（IP地址、用户代理等）

### 5. 安全特性
- JWT身份认证
- 密码哈希加密
- 请求参数验证
- 用户权限控制

## 技术栈

- **后端框架**: Gin (Go)
- **数据库**: MySQL
- **ORM**: GORM
- **认证**: JWT
- **配置管理**: Viper
- **数据库迁移**: golang-migrate
- **日志**: 自定义日志系统
- **依赖注入**: 自定义DI容器

## 环境要求

- Go 1.19+
- MySQL 8.0+
- golang-migrate工具

## 部署步骤

### 1. 克隆项目
```bash
git clone git@github.com:caryxiao/meta-blog.git
cd meta-blog
```

### 2. 安装依赖
```bash
go mod download
```

### 3. 配置数据库
创建MySQL数据库：
```sql
CREATE DATABASE meta_blog CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

### 4. 配置文件
复制并修改配置文件：
```bash
cp configs/config.yaml configs/config.dev.yaml
```

编辑 `configs/config.dev.yaml`，配置数据库连接信息：
```yaml
database:
  host: localhost
  port: 3306
  username: root
  password: your_password
  dbname: meta_blog
  charset: utf8mb4
  parseTime: true
  loc: Local

server:
  port: 8080
  mode: debug

jwt:
  secret: your_jwt_secret_key
  expire: 24h
```

### 5. 安装数据库迁移工具
```bash
# macOS
brew install golang-migrate

# 或者使用go install
go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```

### 6. 运行数据库迁移

#### 方法一：使用golang-migrate工具
```bash
migrate -path db/migrations -database "mysql://username:password@tcp(localhost:3306)/meta_blog?charset=utf8mb4&parseTime=True&loc=Local" up
```

#### 方法二：使用项目内置迁移工具（推荐）
```bash
# 执行所有迁移（默认action=up）
go run cmd/migrate/main.go

# 或者明确指定action
go run cmd/migrate/main.go -action=up

# 回滚一个迁移
go run cmd/migrate/main.go -action=down

# 删除所有表（慎用）
go run cmd/migrate/main.go -action=drop

# 强制设置迁移版本
go run cmd/migrate/main.go -action=force

# 指定环境（默认为dev）
APP_ENV=prod go run cmd/migrate/main.go -action=up
```

### 7. 启动服务
```bash
# 开发环境
go run cmd/blog/main.go

# 或者编译后运行
go build -o bin/blog cmd/blog/main.go
./bin/blog
```

### 8. 验证部署
访问 `http://localhost:8080` 确认服务正常运行。

## API 文档

### 在线 API 文档

项目提供了完整的 API 文档，您可以通过以下链接查看：

**Apifox API 文档**: [https://apifox.com/apidoc/shared-f305b3b6-0d63-45d2-bf69-ba6cd98865cb](https://apifox.com/apidoc/shared-f305b3b6-0d63-45d2-bf69-ba6cd98865cb)

该文档包含：
- 📋 完整的 API 接口列表
- 📝 详细的请求参数说明
- 💡 响应示例和数据结构
- 🧪 在线接口测试功能
- 📖 接口使用说明和示例代码

### API 使用示例

#### 用户注册
```bash
curl -X POST http://localhost:8080/api/register \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "email": "test@example.com",
    "password": "password123"
  }'
```

#### 用户登录
```bash
curl -X POST http://localhost:8080/api/login \
  -H "Content-Type: application/json" \
  -d '{
    "username": "testuser",
    "password": "password123"
  }'
```

#### 创建文章（需要认证）
```bash
curl -X POST http://localhost:8080/api/posts \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_JWT_TOKEN" \
  -d '{
    "title": "我的第一篇文章",
    "content": "这是文章内容...",
    "summary": "文章摘要"
  }'
```

### 获取文章列表
```bash
curl -X GET "http://localhost:8080/api/posts?page=1&page_size=10"
```

## 开发指南

### 添加新的API端点
1. 在 `internal/request/` 中定义请求结构体
2. 在 `internal/response/` 中定义响应结构体
3. 在 `internal/handler/` 中实现处理器方法
4. 在 `internal/router/router.go` 中注册路由
5. 如需数据库操作，在 `internal/repository/` 和 `internal/service/` 中实现相应逻辑

### 数据库迁移

#### 创建迁移文件
```bash
# 创建新的迁移文件
migrate create -ext sql -dir db/migrations -seq migration_name
```

#### 执行迁移
```bash
# 方法一：使用golang-migrate工具
migrate -path db/migrations -database "mysql://..." up
migrate -path db/migrations -database "mysql://..." down 1

# 方法二：使用项目内置工具（推荐）
go run cmd/migrate/main.go -action=up
go run cmd/migrate/main.go -action=down
```

## 许可证

本项目采用 MIT 许可证。详情请参阅 LICENSE 文件。

## 贡献

欢迎提交 Issue 和 Pull Request 来改进这个项目！

## 联系方式

如有问题或建议，请通过以下方式联系：
- 提交 GitHub Issue
- 发送邮件至项目维护者