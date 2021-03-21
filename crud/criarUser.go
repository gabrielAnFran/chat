package crud

import (
	"chat/banco"
	"chat/models"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

func CriarUsuario(c *gin.Context) {
	//recebe o corpo da requisição via cliente
	corpoRequisicao, erro := ioutil.ReadAll(c.Request.Body)
	if erro != nil {
		c.JSON(400, gin.H{
			"mensagem": "erro ao ler o corpo da requisição",
		})
		return
	}
	//cria a variavel do tipo users onde serao guardadas as infromaçoes desjasonizadas
	var usuario models.Users
	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		c.JSON(400, gin.H{
			"mensagem": "erro ao ler o corpo da requisição",
		})
		return
	}

	fmt.Println(usuario)
	//criamos esse usuario no banco de dados
	banco.DBClient.Create(&usuario)

}
