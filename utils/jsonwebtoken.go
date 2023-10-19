package utils

import (
	"perpustakaan/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/gommon/log"
)

type JSONWebToken struct {
	AccessToken  string `json:"access-token"`
	RefreshToken string `json:"refresh-token"`
}

func GenerateToken(id int) *JSONWebToken {
	config := config.LoadServerConfig()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": id,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Hour * 1).Unix(),
	})

	access, err := token.SignedString([]byte(config.SIGN_KEY))
	refresh := getRefreshToken(config.REFRESH_KEY)
	
	if err != nil {
		log.Error(err.Error())
		access = ""
	}

	return &JSONWebToken{
		AccessToken: access,
		RefreshToken: refresh,
	}
}

func getRefreshToken(refreshKey string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	refresh, err := token.SignedString([]byte(refreshKey))

	if err != nil {
		log.Error(err.Error())
		refresh = ""
	}

	return refresh
}