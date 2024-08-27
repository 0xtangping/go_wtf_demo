package types

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

type CreateArticleReq struct {
    Title   string `json:"title"`
    Content string `json:"content"`
    Author  string `json:"author"`
}

type CreateArticleResp struct {
    ID int64 `json:"id"`
}

type GetArticleReq struct {
    ID int64 `json:"id"`
}

type GetArticleResp struct {
    ID      int64  `json:"id"`
    Title   string `json:"title"`
    Content string `json:"content"`
    Author  string `json:"author"`
}

type UpdateArticleReq struct {
    ID      int64  `json:"id"`
    Title   string `json:"title"`
    Content string `json:"content"`
}

type DeleteArticleReq struct {
    ID int64 `json:"id"`
}
