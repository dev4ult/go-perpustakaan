package helpers

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

func (h *helper) GenerateToken(id int, role string) *JSONWebToken {
	config := config.LoadServerConfig()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": id,
		"role": role,
		"iat": time.Now().Unix(),
		"exp": time.Now().Add(time.Second * 15).Unix(),
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

func  ExtractToken(accessToken string, isRefreshToken bool) map[string]any {
	cfg := config.LoadServerConfig()
	key := []byte(cfg.SIGN_KEY)

	if isRefreshToken {
		key = []byte(cfg.REFRESH_KEY)
	}

	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(accessToken, &claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return key, nil
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

