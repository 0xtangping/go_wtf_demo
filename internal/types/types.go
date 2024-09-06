package types

import (
	"time"
)

type GithubLoginReq struct {
	Code string `json:"code"`
}

type EthereumLoginReq struct {
	Address   string `json:"address"`
	Signature string `json:"signature"`
	Message   string `json:"message"`
}

type LoginResp struct {
	Token string `json:"token"`
}

type Article struct {
    ID        int64     `json:"id"`        // 文章ID
    Title     string    `json:"title"`     // 文章标题
    Content   string    `json:"content"`   // 文章内容
    Author    string    `json:"author"`    // 作者
    CreatedAt time.Time `json:"created_at"` // 创建时间
    UpdatedAt time.Time `json:"updated_at"` // 更新时间
}

type CreateArticleReq struct {
    Title   string `json:"title"`
    Content string `json:"content"`
    Author  string `json:"author"`
}

type CreateArticleResp struct {
    ID int64 `json:"id"`
}

type GetArticleReq struct {
    ID int64 `path:"id"`
}

type GetArticleResp struct {
    ID        int64     `json:"id"`        // 文章ID
    Title     string    `json:"title"`     // 文章标题
    Content   string    `json:"content"`   // 文章内容
    Author    string    `json:"author"`    // 作者
    CreatedAt time.Time `json:"created_at"` // 创建时间
    UpdatedAt time.Time `json:"updated_at"` // 更新时间
}

type GetArticlesReq struct {
    Limit   int64 `form:"limit"`
    Page    int64 `form:"page"`
}

type GetArticlesResp struct {
    Articles  []Article `json:"articles"`    // 文章列表
}

type UpdateArticleReq struct {
    ID      int64  `path:"id"`
    Title   string `json:"title"`
    Content string `json:"content"`
}

type DeleteArticleReq struct {
    ID int64 `path:"id"`
}
