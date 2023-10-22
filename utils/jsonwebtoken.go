package utils

import (
	"fmt"
	"perpustakaan/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/gommon/log"
)

type JSONWebToken struct {
	AccessToken  string `json:"access-token"`
	RefreshToken string `json:"refresh-token"`
}

func GenerateToken(id int, role string) *JSONWebToken {
	config := config.LoadServerConfig()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": id,
		"role": role,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	})

	access, err := token.SignedString([]byte(config.SIGN_KEY))
	refresh := getRefreshToken(id, role, config.REFRESH_KEY)
	
	if err != nil {
		log.Error(err.Error())
		access = ""
	}

	return &JSONWebToken{
		AccessToken: access,
		RefreshToken: refresh,
	}
}

func ExtractToken(accessToken string) map[string]any {
	cfg := config.LoadServerConfig()

	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(accessToken, &claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return []byte(cfg.SIGN_KEY), nil
	})

	if err != nil {
		log.Error("Error Parsing Token : ", err.Error())
		return nil
	}

	if token.Valid {
		return map[string]any {
			"role": claims["role"],
			"user-id": claims["id"],
		}
	}

	return nil
}

func ExtractRefreshToken(refreshToken string) jwt.MapClaims {
	cfg := config.LoadServerConfig()

	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(refreshToken, claims, func(t *jwt.Token) (interface{}, error) {
		return cfg.REFRESH_KEY, nil
	})

	if err != nil {
		log.Error("Error Parsing Refresh Token : ", err.Error())
		return nil
	}

	if token.Valid {
		return token.Claims.(jwt.MapClaims)
	}

	return nil
}

func getRefreshToken(id int, role string, refreshKey string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": id,
		"role": role,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	refresh, err := token.SignedString([]byte(refreshKey))

	if err != nil {
		log.Error(err.Error())
		refresh = ""
	}

	return refresh
}