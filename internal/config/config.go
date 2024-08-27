package config

import (
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
    rest.RestConf
    Database struct {
        Driver string
        Source string
    }
	GithubOAuth struct {
		ClientID     string
		ClientSecret string
		RedirectURL  string
	}
    Ethereum struct {
        ChainID int64
		EthereumRPC string
    }
    Auth struct {
        AccessSecret string
        AccessExpire int64
    }
}