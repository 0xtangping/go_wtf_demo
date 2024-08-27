package logic

import (
	"context"
	"wtf_demo/internal/model"
	"wtf_demo/internal/svc"
	"wtf_demo/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ArticleLogic struct {
    logx.Logger
    ctx    context.Context
    svcCtx *svc.ServiceContext
}

func NewArticleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ArticleLogic {
    return &ArticleLogic{
        Logger: logx.WithContext(ctx),
        ctx:    ctx,
        svcCtx: svcCtx,
    }
}

func (l *ArticleLogic) CreateArticle(req *types.CreateArticleReq) (*types.CreateArticleResp, error) {
    article := model.Article{
        Title:   req.Title,
        Content: req.Content,
        Author:  req.Author,
    }

    result, err := l.svcCtx.ArticleModel.Insert(l.ctx, &article)
    if err != nil {
        return nil, err
    }

    lastInsertId, err := result.LastInsertId() // 获取 LastInsertId 和 错误信息
    if err != nil {
        return nil, err
    }
    
    return &types.CreateArticleResp{
        ID: lastInsertId,
    }, nil
}

func (l *ArticleLogic) GetArticle(req *types.GetArticleReq) (*types.GetArticleResp, error) {
    article, err := l.svcCtx.ArticleModel.FindOne(l.ctx, req.ID)
    if err != nil {
        return nil, err
    }

    return &types.GetArticleResp{
        ID:      article.Id,
        Title:   article.Title,
        Content: article.Content,
        Author:  article.Author,
    }, nil
}

func (l *ArticleLogic) UpdateArticle(req *types.UpdateArticleReq) error {
    article := model.Article{
        Id:      req.ID,
        Title:   req.Title,
        Content: req.Content,
    }

    return l.svcCtx.ArticleModel.Update(l.ctx, &article)
}

func (l *ArticleLogic) DeleteArticle(req *types.DeleteArticleReq) error {
    return l.svcCtx.ArticleModel.Delete(l.ctx, req.ID)
}
