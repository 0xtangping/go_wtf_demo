package svc

import (
	"wtf_demo/internal/config"
	"wtf_demo/internal/model"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
    Config config.Config
    EthClient   *ethclient.Client
    ArticleModel model.ArticleModel

}

func NewServiceContext(c config.Config) *ServiceContext {
    conn := sqlx.NewMysql(c.Database.Source)
	ethClient, err := ethclient.Dial(c.Ethereum.EthereumRPC)
    articleModel := model.NewArticleModel(conn)
    if err != nil {
        panic(err)
    }
    return &ServiceContext{
        Config:    c,
        EthClient: ethClient,
        ArticleModel: articleModel,
    }
}
