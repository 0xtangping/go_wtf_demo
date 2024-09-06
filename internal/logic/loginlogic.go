package logic

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"
	"time"

	"wtf_demo/internal/svc"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/go-github/v33/github"
	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/oauth2"
	oauth2git "golang.org/x/oauth2/github"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) GithubLogin(code string) (string, error) {
	config := &oauth2.Config{
		ClientID:     l.svcCtx.Config.GithubOAuth.ClientID,
		ClientSecret: l.svcCtx.Config.GithubOAuth.ClientSecret,
		RedirectURL:  l.svcCtx.Config.GithubOAuth.RedirectURL,
		Endpoint:     oauth2git.Endpoint,
	}

	token, err := config.Exchange(l.ctx, code)
	if err != nil {
		return "", fmt.Errorf("github OAuth exchange failed: %v", err)
	}

	client := github.NewClient(config.Client(l.ctx, token))
	user, _, err := client.Users.Get(l.ctx, "")
	if err != nil {
		return "", fmt.Errorf("failed to get github user: %v", err)
	}

	// Generate JWT token
	return l.generateToken(fmt.Sprintf("github_%d", user.GetID()))
}

func (l *LoginLogic) EthereumLogin(address, signature, message string) (string, error) {
	// Verify Ethereum signature
	if !verifyEthereumSignature(address, signature, message) {
		return "", errors.New("invalid Ethereum signature")
	}

	// Generate JWT token
	return l.generateToken(fmt.Sprintf("ethereum_%s", address))
}

func (l *LoginLogic) generateToken(userId string) (string, error) {
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Auth.AccessExpire

	claims := make(jwt.MapClaims)
	claims["exp"] = now + accessExpire
	claims["iat"] = now
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	// 签名并生成 token 字符串
    signedToken, err := token.SignedString([]byte(l.svcCtx.Config.Auth.AccessSecret))
    if err != nil {
        return "", err
    }

    return signedToken, nil
}


func verifyEthereumSignature(address, signature, message string) bool {
	// Remove '0x' prefix if present
	signature = strings.TrimPrefix(signature, "0x")

	// Decode the signature
	sig, err := hex.DecodeString(signature)
	if err != nil || len(sig) != 65 {
		return false
	}

	// Adjust the recovery id (v)
	sig[64] -= 27

	// Add Ethereum message prefix
	prefixedMessage := fmt.Sprintf("\x19Ethereum Signed Message:\n%d%s", len(message), message)

	// Hash the prefixed message
	hash := crypto.Keccak256Hash([]byte(prefixedMessage))

	// Recover the public key
	pubKey, err := crypto.SigToPub(hash.Bytes(), sig)
	if err != nil {
		return false
	}

	// Convert the public key to an address
	recoveredAddress := crypto.PubkeyToAddress(*pubKey)

	// Compare the recovered address with the provided address
	return strings.EqualFold(recoveredAddress.Hex(), address)
}