package service

import (
	"errtrack/internal/config"
	"fmt"
	"github.com/golang-jwt/jwt"
	"time"
)

type JWTService struct {
	secretKey string
}

func NewJWTService(securityConfig *config.SecurityConfig) *JWTService {
	return &JWTService{
		secretKey: securityConfig.JwtSecurityKey,
	}
}

// GenerateToken 生成 JWT Token，包含 claims 和过期时间
func (s *JWTService) GenerateToken(username string, expiration time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(expiration).Unix(),
	}

	// 创建 JWT Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// 使用密钥签名并生成 token 字符串
	tokenString, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %v", err)
	}

	return tokenString, nil
}

// ValidateToken 验证 JWT Token 是否有效
func (s *JWTService) ValidateToken(tokenString string) (*jwt.Token, error) {
	// 解析并验证 token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// 检查 token 的签名方法是否正确
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// 返回密钥用于验证签名
		return []byte(s.secretKey), nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to validate token: %v", err)
	}

	return token, nil
}

func (s *JWTService) ParseToken(tokenString string) (jwt.MapClaims, error) {
	token, err := s.ValidateToken(tokenString)
	if err != nil {
		return nil, err
	}

	// 确保 token 是有效的并且没有过期
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}
