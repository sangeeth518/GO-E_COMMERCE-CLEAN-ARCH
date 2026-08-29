package helper

import (
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"path/filepath"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/sangeeth518/go-Ecommerce/pkg/config"
	interfaces "github.com/sangeeth518/go-Ecommerce/pkg/helper/interface"
	"github.com/sangeeth518/go-Ecommerce/pkg/utils/models"
	"golang.org/x/crypto/bcrypt"
)

type helper struct {
	config config.Config
	s3     *s3.Client
}

func NewHelper(config config.Config) (interfaces.Helper, error) {
	s3Client, err := InitS3Client(config)
	if err != nil {
		return nil, err
	}
	return &helper{
		config: config,
		s3:     s3Client,
	}, nil
}

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
	refreshTokenstrig, err := refreshToken.SignedString([]byte(h.config.RefreshToken))
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
	tokenstrig, err := token.SignedString([]byte(h.config.UserJWTToken))
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

// Initializing S3 client

func InitS3Client(config config.Config) (*s3.Client, error) {
	awscfg, err := awsConfig.LoadDefaultConfig(context.TODO(),
		awsConfig.WithRegion(config.AWSRegion),
		awsConfig.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(config.AWSKey, config.AWSSecret, "")),
	)
	if err != nil {
		return nil, err
	}
	client := s3.NewFromConfig(awscfg)
	return client, nil
}

//ADD PRODUCT IMAGE

func (h *helper) AddProductImage(ctx context.Context, file *multipart.FileHeader, productId int) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	ext := filepath.Ext(file.Filename)

	filename := fmt.Sprintf("products/%d%s%s", productId, uuid.New().String(), ext)

	_, err = h.s3.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(h.config.BucketName),
		Key:         aws.String(filename),
		Body:        src,
		ContentType: aws.String("image/" + strings.TrimPrefix(ext, ".")),
	})
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s",
		h.config.BucketName,
		h.config.AWSRegion,
		filename,
	)

	return url, nil

}
