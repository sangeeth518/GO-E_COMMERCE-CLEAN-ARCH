package helper

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/sangeeth518/go-Ecommerce/pkg/config"
	"github.com/sangeeth518/go-Ecommerce/pkg/utils/models"
	"github.com/twilio/twilio-go"
)

type helper struct {
	config config.Config
}

func NewHelper(config config.Config) *helper {
	return &helper{
		config: config,
	}
}

var client *twilio.RestClient

type AuthCustomClaims struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

func (h *helper) GenerateTokenAdmin(admin models.AdminDetailResponse) (string, string, error) {
	accessTokenClaims := &AuthCustomClaims{
		ID:    admin.ID,
		Email: admin.Email,
		Role:  "admin",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 20).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	refreshTokenClaims := &AuthCustomClaims{
		ID:    admin.ID,
		Email: admin.Email,
		Role:  "admin",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 30).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	accessTokenstring, err := accessToken.SignedString([]byte(h.config.JWTToken))
	if err != nil {
		return "", "", err
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	refreshTokenstrig, err := refreshToken.SignedString([]byte("refreshsecret"))
	if err != nil {
		return "", "", err
	}
	return accessTokenstring, refreshTokenstrig, nil

}
