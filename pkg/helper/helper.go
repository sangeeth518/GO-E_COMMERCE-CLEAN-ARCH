package helper

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/sangeeth518/go-Ecommerce/pkg/config"
	interfaces "github.com/sangeeth518/go-Ecommerce/pkg/helper/interface"
	"github.com/sangeeth518/go-Ecommerce/pkg/utils/models"
	"github.com/twilio/twilio-go"
	"golang.org/x/crypto/bcrypt"
)

type helper struct {
	config config.Config
}

func NewHelper(config config.Config) interfaces.Helper {
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

func (h *helper) GenerateTokenClient(user models.UserDetailsResponse) (string, error) {
	Claims := &AuthCustomClaims{
		ID:    user.Id,
		Email: user.Email,
		Role:  "user",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims)
	tokenstrig, err := token.SignedString([]byte("usersecret"))
	if err != nil {
		return "", err
	}
	return tokenstrig, nil
}

func (h *helper) PasswordHashing(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", errors.New("password hashing error")
	}
	hash := string(hashed)
	return hash, nil

}

func (h *helper) CompareHashPassword(password string, givenpass string) error {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(givenpass))
	if err != nil {
		return err
	}
	return nil
}
