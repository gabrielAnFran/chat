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

//Deletar a conversa baseada no numero do id do usuario que fez a requisição.
func DeletarConversasPorId(c *gin.Context) {

	if Token != "" {

		print("EndpointHIT")
		corpoRequisicao, erro := ioutil.ReadAll(c.Request.Body)
		if erro != nil {
			c.JSON(400, gin.H{
				"mensagem": "erro ao ler o corpo da requisição",
			})
			return
		}
		var msg models.LerMensagens
		if erro = json.Unmarshal(corpoRequisicao, &msg); erro != nil {
			c.JSON(400, gin.H{
				"mensagem": "erro ao ler o corpo da requisição",
			})
			return
		}
		fmt.Println(msg)
		if msg.IdRemetente == int(IdSessao) {
			var msgs []models.Mensagens
			banco.DBClient.Where("mensagens.id_remetente = ? ", msg.IdRemetente).Delete(&msgs)

			c.JSON(http.StatusOK, gin.H{
				"mensagem": "Mensagens deletadas",
			})
		}
	} else {
		c.JSON(400, gin.H{
			"mensagem": "Voce não tem autorização de alterar esse usuario e senha.",
		})
	}
	{
		c.JSON(400, gin.H{
			"mensagem": "Essa ação requere autenticação. Por favor, faça login",
		})

	}

}
