package handler

import (
	"net/http"
	"wtf_demo/internal/logic"
	"wtf_demo/internal/svc"
	"wtf_demo/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func CreateArticleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var req types.CreateArticleReq
        if err := httpx.Parse(r, &req); err != nil {
            httpx.Error(w, err)
            return
        }

        l := logic.NewArticleLogic(r.Context(), svcCtx)
        if resp, err := l.CreateArticle(&req);err != nil {
            httpx.Error(w, err)
        } else {
            httpx.OkJson(w, resp)
        }
    }
}

func GetArticleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var req types.GetArticleReq
        if err := httpx.Parse(r, &req); err != nil {
            httpx.Error(w, err)
            return
        }

        l := logic.NewArticleLogic(r.Context(), svcCtx)
        if resp, err := l.GetArticle(&req);err != nil {
            httpx.Error(w, err)
        } else {
            httpx.OkJson(w, resp)
        }
    }
}

func UpdateArticleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var req types.UpdateArticleReq
        if err := httpx.Parse(r, &req); err != nil {
            httpx.Error(w, err)
            return
        }
                
        l := logic.NewArticleLogic(r.Context(), svcCtx)
        if err := l.UpdateArticle(&req);err != nil {
            httpx.Error(w, err)
        } else {
            httpx.OkJson(w, map[string]string{"msg": "update successful"})
        }
    }
}

func DeleteArticleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var req types.DeleteArticleReq
        if err := httpx.Parse(r, &req); err != nil {
            httpx.Error(w, err)
            return
        }


        l := logic.NewArticleLogic(r.Context(), svcCtx)
        if err := l.DeleteArticle(&req);err != nil {
            httpx.Error(w, err)
        } else {
            httpx.OkJson(w, map[string]string{"msg": "delete successful"})
        }
    }
}


func GetArticlesHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        var req types.GetArticlesReq
        if err := httpx.Parse(r, &req); err != nil {
            httpx.Error(w, err)
            return
        }

        l := logic.NewArticleLogic(r.Context(), svcCtx)
        resp, err := l.GetArticles(&req)
        if err != nil {
            httpx.Error(w, err)
        } else {
            httpx.OkJson(w, resp)
        }
    }
}
