package crud

import (
	"chat/banco"
	"chat/models"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var Token string
var UsuarioSessao string
var IdSessao uint

func Logar(c *gin.Context) {
	var u models.Users
	//jogamos os dados recebidos em json para  a variavel u do tipo users
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusUnprocessableEntity, "Invalid json provided")
		return
	}
	//verificação de login
	if erro := banco.DBClient.Where(" users.login = ? AND users.senha = ?", u.Login, u.Senha).Find(&u).Error; erro != nil {
		c.JSON(http.StatusUnauthorized, "Por favor, forneça os dados de login e senha corretamente")
		return
	}
	//pegamos o id do usuario baseado no login, pois caso deixassemos o id na struct users, ele nao seria auto incrementado, o que acarretaria em problemas, ou
	//trabalhos a mais ao criar um usuario... teriamos que saber qual foi o ultimo id e enviar um id o queal seria id do novo usuario.....
	var u2 models.AtualizarSenha
	banco.DBClient.Raw("SELECT id_usuario FROM users WHERE login = ?", u.Login).Scan(&u2)
	id := u2.IdUsuario
	fmt.Println("id: ", id)

	//até aqui eu consigo criar um token, porém nao estou aplicando ele da maneira correta.
	token, err := CreateToken(uint64(id))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	c.JSON(http.StatusOK, token)
	//atribuo valores a essas variaveis para poder usá-las fora desse escopo, e assim fazer as restriçoes e autorizaçoes baseadas no login
	Token = token
	UsuarioSessao = u.Login
	IdSessao = id

}

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
