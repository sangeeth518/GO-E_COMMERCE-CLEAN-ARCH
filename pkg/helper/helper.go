package helper

import (
	"bytes"
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
	"github.com/jung-kurt/gofpdf"
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

	// Return the S3 key
	return filename, nil
}

// GET PRESIGNED URL FOR IMAGE

func (h *helper) GetPresignedURL(ctx context.Context, key string) (string, error) {
	presignClient := s3.NewPresignClient(h.s3)
	presignedReq, err := presignClient.PresignGetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(h.config.BucketName),
		Key:    aws.String(key),
	}, s3.WithPresignExpires(time.Hour*1))
	if err != nil {
		return "", err
	}
	return presignedReq.URL, nil
}

// DELETE PRODUCT IMAGE FROM S3

func (h *helper) DeleteProductImageFromS3(ctx context.Context, key string) error {
	if key == "" {
		return errors.New("invalid s3 object key")
	}

	_, err := h.s3.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(h.config.BucketName),
		Key:    aws.String(key),
	})
	if err != nil {
		return err
	}

	return nil
}

// GenerateInvoicePDF creates a professional PDF invoice from order details
func (h *helper) GenerateInvoicePDF(details models.OrderDetailsResponse) ([]byte, error) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.SetMargins(15, 15, 15)
	pdf.AddPage()

	// Title
	pdf.SetFont("Arial", "B", 20)
	pdf.CellFormat(100, 10, "GO E-COMMERCE", "", 0, "L", false, 0, "")
	pdf.SetFont("Arial", "B", 14)
	pdf.CellFormat(80, 10, "TAX INVOICE", "", 1, "R", false, 0, "")

	pdf.Ln(4)
	pdf.SetDrawColor(200, 200, 200)
	pdf.Line(15, pdf.GetY(), 195, pdf.GetY())
	pdf.Ln(6)

	// Order Info & Shipping Address
	yBefore := pdf.GetY()

	// Left: Shipping Address
	pdf.SetFont("Arial", "B", 11)
	pdf.CellFormat(90, 6, "Billed To:", "", 1, "L", false, 0, "")
	pdf.SetFont("Arial", "", 10)
	pdf.CellFormat(90, 5, details.Address.Name, "", 1, "L", false, 0, "")
	pdf.CellFormat(90, 5, fmt.Sprintf("%s, %s", details.Address.HouseName, details.Address.Street), "", 1, "L", false, 0, "")
	pdf.CellFormat(90, 5, fmt.Sprintf("%s, %s - %s", details.Address.City, details.Address.State, details.Address.Pin), "", 1, "L", false, 0, "")
	pdf.CellFormat(90, 5, fmt.Sprintf("Phone: %s", details.Address.Phone), "", 1, "L", false, 0, "")

	yLeft := pdf.GetY()

	// Right: Order Metadata
	pdf.SetXY(110, yBefore)
	pdf.SetFont("Arial", "B", 11)
	pdf.CellFormat(85, 6, "Order Information:", "", 1, "L", false, 0, "")
	pdf.SetFont("Arial", "", 10)
	pdf.SetX(110)
	pdf.CellFormat(85, 5, fmt.Sprintf("Invoice / Order ID: #%d", details.Order.Id), "", 1, "L", false, 0, "")
	pdf.SetX(110)
	pdf.CellFormat(85, 5, fmt.Sprintf("Date: %s", details.Order.CreatedAt.Format("02 Jan 2006 15:04")), "", 1, "L", false, 0, "")
	pdf.SetX(110)
	pdf.CellFormat(85, 5, fmt.Sprintf("Payment Method: %s", details.Order.PaymentMethod), "", 1, "L", false, 0, "")
	pdf.SetX(110)
	pdf.CellFormat(85, 5, fmt.Sprintf("Payment Status: %s", details.Order.PaymentStatus), "", 1, "L", false, 0, "")
	pdf.SetX(110)
	pdf.CellFormat(85, 5, fmt.Sprintf("Order Status: %s", details.Order.OrderStatus), "", 1, "L", false, 0, "")

	yRight := pdf.GetY()

	// Move below the taller column
	if yLeft > yRight {
		pdf.SetY(yLeft + 8)
	} else {
		pdf.SetY(yRight + 8)
	}

	// Table Headers
	pdf.SetFillColor(240, 240, 240)
	pdf.SetFont("Arial", "B", 10)
	pdf.CellFormat(10, 8, "#", "1", 0, "C", true, 0, "")
	pdf.CellFormat(75, 8, "Product Name", "1", 0, "L", true, 0, "")
	pdf.CellFormat(20, 8, "Size", "1", 0, "C", true, 0, "")
	pdf.CellFormat(20, 8, "Qty", "1", 0, "C", true, 0, "")
	pdf.CellFormat(25, 8, "Unit Price", "1", 0, "R", true, 0, "")
	pdf.CellFormat(30, 8, "Total Price", "1", 1, "R", true, 0, "")

	// Table Rows
	pdf.SetFont("Arial", "", 9)
	for i, item := range details.OrderItems {
		pdf.CellFormat(10, 7, fmt.Sprintf("%d", i+1), "1", 0, "C", false, 0, "")
		pdf.CellFormat(75, 7, item.ProductName, "1", 0, "L", false, 0, "")
		pdf.CellFormat(20, 7, item.Size, "1", 0, "C", false, 0, "")
		pdf.CellFormat(20, 7, fmt.Sprintf("%d", item.Quantity), "1", 0, "C", false, 0, "")
		pdf.CellFormat(25, 7, fmt.Sprintf("%.2f", item.UnitPrice), "1", 0, "R", false, 0, "")
		pdf.CellFormat(30, 7, fmt.Sprintf("%.2f", item.TotalPrice), "1", 1, "R", false, 0, "")
	}

	// Total Price Summary
	pdf.Ln(2)
	pdf.SetFont("Arial", "B", 11)
	pdf.CellFormat(150, 8, "Final Amount:", "", 0, "R", false, 0, "")
	pdf.CellFormat(30, 8, fmt.Sprintf("Rs. %.2f", details.Order.FinalPrice), "1", 1, "R", true, 0, "")

	// Footer
	pdf.Ln(15)
	pdf.SetFont("Arial", "I", 9)
	pdf.SetTextColor(120, 120, 120)
	pdf.CellFormat(180, 5, "Thank you for shopping with Go-Ecommerce! This is a computer-generated invoice.", "", 1, "C", false, 0, "")

	var buf bytes.Buffer
	err := pdf.Output(&buf)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
