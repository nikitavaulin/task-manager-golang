package tools_jwt

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
	tools_envparser "github.com/nikitavaulin/task-manager-golang/internal/core/tools/env_parser"
)

func getSecretKey() []byte {
	return []byte(tools_envparser.GetEnvVarOrDefault("JWT_SECRET", "secret"))
}

func GenerateToken(claims jwt.MapClaims) (string, error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secretKey := getSecretKey()

	signedToken, err := jwtToken.SignedString(secretKey)
	if err != nil {
		return "", fmt.Errorf("failed to sign jwt token: %w", err)
	}

	return signedToken, nil
}

func DecodeClaims(token string) (map[string]any, error) {
	jwtToken, err := jwt.Parse(token, func(t *jwt.Token) (any, error) {
		return getSecretKey(), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil {
		return nil, fmt.Errorf("failed to parse jwt token: %w", err)
	}
	if !jwtToken.Valid {
		return nil, fmt.Errorf("invalid jwt token")
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("failed to convert claims to MapClaims")
	}

	return claims, nil
}
