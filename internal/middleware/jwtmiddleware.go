package middleware

import (
	"context"
	"errors"
	"net/http"
	"strings"

	"wtf_demo/internal/svc"

	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/rest/httpx"
)

type JWTMiddleware struct {
    svcCtx *svc.ServiceContext
}

// NewJWTMiddleware 创建一个新的 JWT 中间件实例
func NewJWTMiddleware(svcCtx *svc.ServiceContext) *JWTMiddleware {
    return &JWTMiddleware{
        svcCtx: svcCtx,
    }
}

// Handle 是实际执行中间件逻辑的函数
func (m *JWTMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // 从 Authorization 头中提取 token
        authHeader := r.Header.Get("Authorization")
        if authHeader == "" {
            httpx.Error(w, errors.New("Authorization header missing"))
            return
        }

        // 按照 "Bearer {token}" 的格式提取 token
        parts := strings.SplitN(authHeader, " ", 2)
        if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
            httpx.Error(w, errors.New("Invalid Authorization header format"))
            return
        }

        tokenStr := parts[1]

        // 解析并验证 JWT
        token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
            // 验证算法
            if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                return nil, errors.New("Unexpected signing method")
            }
            // 从 ServiceContext 中获取 JWT 密钥
            return []byte(m.svcCtx.Config.Auth.AccessSecret), nil
        })

        if err != nil || !token.Valid {
            httpx.Error(w, errors.New("Invalid or expired token:${}"))
            return
        }

        // 提取 JWT 中的 claims
        claims, ok := token.Claims.(jwt.MapClaims)
        if !ok || !token.Valid {
            httpx.Error(w, errors.New("Invalid token claims"))
            return
        }

        // 将 claims 保存到请求的上下文中
        ctx := context.WithValue(r.Context(), "user", claims)
        r = r.WithContext(ctx)

        // 调用下一个处理函数
        next(w, r)
    }
}
