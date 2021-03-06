package middleware

import (
	"fmt"
	"github.com/njilrem/go-rest-atm/services"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthorizeJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BearerSchema = "Bearer "
		authHeader := c.GetHeader("Authorization")
		tokenString := authHeader[len(BearerSchema):]
		token, err := services.JWTAuthService().ValidateToken(tokenString)
		if err != nil {
			fmt.Println("Validation")
			fmt.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
			fmt.Println("Claims Error")
			fmt.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if token.Valid {
		} else {
			fmt.Println("Error")
			fmt.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

	}
}

func AuthorizeAdminJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BearerSchema = "Bearer "
		authHeader := c.GetHeader("Authorization")
		tokenString := authHeader[len(BearerSchema):]
		token, err := services.JWTAuthService().ValidateToken(tokenString)
		if err != nil {
			fmt.Println("Validation")
			fmt.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
			fmt.Println("Claims Error")
			fmt.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			isAdmin := claims["isValidAccount"]
			if isAdmin == true {
				fmt.Println("ADMIN TOKEN USED")
			} else {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
		} else {
			fmt.Println("Error")
			fmt.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

	}
}

func AuthorizeTransactionJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		const BearerSchema = "Bearer "
		authHeader := c.GetHeader("Authorization")
		tokenString := authHeader[len(BearerSchema):]
		token, err := services.JWTAuthService().ValidateToken(tokenString)
		if err != nil {
			fmt.Println("Validation")
			fmt.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
			fmt.Println("Claims Error")
			fmt.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if token.Valid {
			//claims := token.Claims.(jwt.MapClaims)
			//var transactionInfo dto.AuthTransaction
			//err := c.ShouldBind(transactionInfo)
			return
		} else {
			fmt.Println(err)
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

	}
}