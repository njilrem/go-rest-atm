package services

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//jwt service
type JWTService interface {
	GenerateToken(cardNum string, expireDate string, cvv2 string, isValidAccount bool) string
	ValidateToken(token string) (*jwt.Token, error)
}
type authCustomClaims struct {
	CardNum        string `json:"cardNum"`
	ExpireDate     string `json:"expireDate"`
	Cvv2           string `json:"cvv2"`
	IsValidAccount bool   `json:"isValidAccount"`
	jwt.StandardClaims
}

type jwtServices struct {
	secretKey string
	issuer    string
}

func JWTAuthService() JWTService {
	return &jwtServices{
		secretKey: getSecretKey(),
		issuer:    "Cookie Bank",
	}
}

func getSecretKey() string {
	secret := os.Getenv("SECRET")
	if secret == "" {
		secret = "secret"
	}
	return secret
}

func (service *jwtServices) GenerateToken(cardNum string, expireDate string, cvv2 string, isValidAccount bool) string {
	claims := &authCustomClaims{
		cardNum,
		expireDate,
		cvv2,
		isValidAccount,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 15).Unix(),
			Issuer:    service.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//encoded string
	t, err := token.SignedString([]byte(service.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (service *jwtServices) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isValid := token.Method.(*jwt.SigningMethodHMAC); !isValid {
			return nil, fmt.Errorf("Invalid token", token.Header["alg"])

		}
		return []byte(service.secretKey), nil
	})

}