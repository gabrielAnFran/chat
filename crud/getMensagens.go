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

func GetMensagensDuasPessoas(c *gin.Context) {

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
		//id sessao é a variavel que nos diz quem está logado, assim evitando que usuario x faça alteraçoes que cabem a y ....
		if msg.IdRemetente == int(IdSessao) {
			fmt.Println(msg)
			var msgs []models.Mensagens
			banco.DBClient.Where("mensagens.id_remetente = ? AND mensagens.id_destinatario = ? OR  mensagens.id_remetente = ? AND mensagens.id_destinatario = ?", msg.IdRemetente, msg.IdDestinatario, msg.IdDestinatario, msg.IdRemetente).Find(&msgs)

			c.JSON(http.StatusOK, msgs)
		} else {
			c.JSON(400, gin.H{
				"mensagem": "Voce nao está autorizado a operar nem visualizar transaçoes do usuario que solicitou.",
			})

		}
	} else {
		c.JSON(400, gin.H{
			"mensagem": "Essa ação requere autenticação. Por favor, faça login",
		})

	}

}
