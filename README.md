# 文章管理后端系统

为了学习go做的一个防wtf.academy的后端系统，网址：https://shopify.xyz/(仅作于学习交流，随时关闭)，使用 [Go-Zero](https://github.com/zeromicro/go-zero) 框架开发的文章管理后端系统，支持文章的增删改查操作以及基于登录权限的用户认证管理。


## 功能特色

- **用户认证**：基于 JWT 的安全登录认证机制。
- **文章管理**：支持创建、查看、更新和删除文章的完整操作。
- **基于角色的访问控制**：为管理员和普通用户提供不同的权限控制。
- **RESTful API**：符合 REST 风格的清晰接口。
- **错误处理**：对各种边界情况提供优雅的错误处理。

## 目录

- [安装](#安装)
- [配置](#配置)
- [API 接口](#api-接口)
- [运行项目](#运行项目)
- [贡献](#贡献)
- [许可证](#许可证)

## 安装

1. 克隆仓库：
    ```bash
    git clone https://github.com/0xtangping/go_wtf_demo.git
    ```
2. 进入项目目录：
    ```bash
    cd go_wtf_demo
    ```
3. 安装依赖：
    ```bash
    go mod tidy
    ```

4. 确保已经安装了 [Go-Zero](https://github.com/zeromicro/go-zero)。

5. 配置数据库（可以使用 MySQL、PostgreSQL 等）并更新配置文件。

## 配置

更新 `config.yaml` 文件中的数据库连接和 JWT 密钥配置。

```yaml
Name: wtf_demo-api
Host: 0.0.0.0
Port: 8080


Database:
  Driver: mysql
  #Source: root:QWEqwe123@tcp(127.0.0.1:3306)/wtf_demo?charset=utf8mb4&parseTime=true&loc=Local
Ethereum:
  ChainID: 1 # Mainnet
  EthereumRPC: https://cloudflare-eth.com

GithubOAuth:
  ClientID: ******
  ClientSecret: ******
  RedirectURL: https://shopify.xyz/login

Auth:
  AccessSecret: *****
  AccessExpire: 7200  # 2 hours

```

## API 接口

### 用户认证

| 方法 | 接口          | 描述          |
|------|---------------|---------------|
| POST | `/api/login/ethereum`  | 以太坊用户登录      |
| POST | `/api/login/github`  | 以太坊用户登录      |


### 文章管理 (CRUD)

| 方法   | 接口                      | 描述               |
|--------|---------------------------|--------------------|
| POST   | `/api/articles`            | 创建新文章         |
| GET    | `/api/articles`            | 获取所有文章       |
| GET    | `/api/articles/:id`        | 根据 ID 获取文章   |
| PUT    | `/api/articles/:id`        | 更新文章           |
| DELETE | `/api/articles/:id`        | 删除文章           |

### 请求示例

#### 登录

```bash
POST /api/login/github
Content-Type: application/json

{
  "code": "examplecode"
}
```

#### 创建文章

```bash
POST /api/articles
Authorization: Bearer <your-token>
Content-Type: application/json

{
  "title": "新文章",
  "content": "这是新文章的内容。"
}
```

#### 获取所有文章

```bash
GET /api/articles
Authorization: Bearer <your-token>
```

## 运行项目

1. 启动 Go-Zero 服务：
    ```bash
    go run main.go
    ```

2. 默认情况下，服务将运行在 `localhost:8080`。

3. 你可以使用 [Postman](https://www.postman.com/) 或 [curl](https://curl.se/) 测试 API。

## 未完善
1. 权限仅做登录判断，没做角色管理。
2. 文章没有category_id字段，仅做简单查询。

## 许可证

本项目使用 MIT 许可证。详情请参阅 [LICENSE](LICENSE) 文件。
