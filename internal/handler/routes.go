package handler

import (
	"net/http"

	"wtf_demo/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/login/github",
				Handler: GithubLoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/login/ethereum",
				Handler: EthereumLoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/article/create",
				Handler: CreateArticleHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/article/get",
				Handler: GetArticleHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/article/update",
				Handler: UpdateArticleHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/article/delete",
				Handler: DeleteArticleHandler(serverCtx),
			},
		},
	)
}
