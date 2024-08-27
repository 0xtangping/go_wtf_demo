package handler

import (
	"net/http"

	"wtf_demo/internal/logic"
	"wtf_demo/internal/svc"
	"wtf_demo/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func GithubLoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GithubLoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.GithubLogin(req.Code)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, &types.LoginResp{Token: resp})
		}
	}
}

func EthereumLoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.EthereumLoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.EthereumLogin(req.Address, req.Signature, req.Message)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, &types.LoginResp{Token: resp})
		}
	}
}