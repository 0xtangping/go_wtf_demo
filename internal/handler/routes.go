package handler

import (
	"net/http"

	"wtf_demo/internal/middleware"
	"wtf_demo/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
    jwtMiddleware := middleware.NewJWTMiddleware(serverCtx)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/api/login/github",
				Handler: GithubLoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/login/ethereum",
				Handler: EthereumLoginHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/api/articles",
				Handler: jwtMiddleware.Handle(CreateArticleHandler(serverCtx)),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/articles/:id",
				Handler: GetArticleHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/api/articles/:id",
				Handler: jwtMiddleware.Handle(UpdateArticleHandler(serverCtx)),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/api/articles/:id",
				Handler: jwtMiddleware.Handle(DeleteArticleHandler(serverCtx)),
			},
			{
				Method:  http.MethodGet,
				Path:    "/api/articles",
				Handler: GetArticlesHandler(serverCtx),
			},
		},
	)

}
