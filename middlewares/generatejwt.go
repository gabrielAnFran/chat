package middlewares

import (
	"fmt"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func CreateToken(userid uint64) (string, error) {
	var err error
	//Criando o token de acesso
	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd") //o ideal seria estar nos arquivos env
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userid
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	fmt.Println(token)
	return token, nil
}
