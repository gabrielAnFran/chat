package crud

import (
	"chat/banco"
	"chat/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func EnviarMensagem(c *gin.Context) {
	fmt.Println("token", Token)
	//caso o token seja vazio, quer dizer q o user nao está logado, logado nao tem autorização para fazer nada por aqui...
	if Token != "" {

		//lemos o corpo da requisição e o guardamos
		corpoRequisicao, erro := ioutil.ReadAll(c.Request.Body)
		if erro != nil {
			c.JSON(400, gin.H{
				"mensagem": "erro ao ler o corpo da requisição",
			})
			return
		}

		//fazemos unmarshall em msg
		var msg models.Mensagens
		if erro = json.Unmarshal(corpoRequisicao, &msg); erro != nil {
			c.JSON(400, gin.H{
				"mensagem": "erro ao ler o corpo da requisição",
			})
			return
		}

		fmt.Println(msg)
		//Validação de acesso a transação
		//Caso todas validaçoes passem, criamos a mensagem e guardamos no banco de dados
		if msg.IdRemetente == int(IdSessao) {
			if erro := banco.DBClient.Create(&msg).Error; erro != nil {
				c.JSON(400, gin.H{
					"mensagem": "ERRO AO CRIAR USUARIO",
				})
			}

			c.JSON(http.StatusAccepted, msg)
		} else {
			c.JSON(400, gin.H{
				"mensagem": "Voce não tem autorização de operar esse usuario.",
			})

		}
	} else {
		c.JSON(400, gin.H{
			"mensagem": "Essa ação requere autenticação. Por favor, faça login",
		})

	}
}
