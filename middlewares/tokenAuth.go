package middlewares

import (
	"fmt"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// JwtMicroservices é o Middleware responsável por validar a autenticidade de uma requisição no BackEnd de um microservice.
func JwtMicroservices() gin.HandlerFunc {
	fmt.Println("Entrada da funçao")
	return func(c *gin.Context) {

		fmt.Println("Entrada da funçao within the function")

		tokenString := c.GetHeader("Authorization")
		if tokenString != "" {

			bearerToken := strings.Split(tokenString, "Bearer ")
			if len(bearerToken) == 2 {

				_, err := jwt.Parse(bearerToken[1], func(token *jwt.Token) (interface{}, error) {
					if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("There was an error")
					}
					return []byte("jdnfksdmfksd"), nil
				})

				if err != nil {
					c.AbortWithStatusJSON(401, gin.H{"authErrors": []string{err.Error()}})
					return
				}
				return

			}

			c.AbortWithStatusJSON(401, gin.H{"authErrors": []string{"Invalid authorization bearer token"}})
			return
		}

		c.AbortWithStatusJSON(401, gin.H{"authErrors": []string{"An authorization header is required"}})
		return
	}

}
