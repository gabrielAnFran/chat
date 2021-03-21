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

func AtualizarSenha(c *gin.Context) {
	if Token != "" {
		print("EndpointHIT")
		corpoRequisicao, erro := ioutil.ReadAll(c.Request.Body)
		if erro != nil {
			c.JSON(400, gin.H{
				"mensagem": "erro ao ler o corpo da requisição",
			})
			return
		}
		var msg models.AtualizarSenha
		if erro = json.Unmarshal(corpoRequisicao, &msg); erro != nil {
			c.JSON(400, gin.H{
				"mensagem": "erro ao ler o corpo da requisição",
			})
			return
		}

		fmt.Println(msg)
		
		var usuario models.Users
		if erro := banco.DBClient.Where("users.login = ? AND users.senha = ?", msg.Login, msg.Senha).Find(&usuario); erro.Error != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"mensagem": "Login e ou senha invalidos. Informe os dados corretamente",
			})
			return
		}
		fmt.Println("usuariologin >", usuario.Login)
		fmt.Println("usuariosessao >", UsuarioSessao)

		if usuario.Login == UsuarioSessao {
			if erro := banco.DBClient.Model(&usuario).Where("users.login = ? ", msg.Login).Update(&models.Users{Senha: msg.NovaSenha}); erro.Error != nil {
				c.JSON(http.StatusNonAuthoritativeInfo, gin.H{
					"message": "Erro ao atualizar usuário",
				})
				return
			}

			c.JSON(http.StatusOK, usuario)

		} else {
			c.JSON(400, gin.H{
				"mensagem": "Voce não tem autorização de alterar esse usuario e senha.",
			})

		}
	} else {
		c.JSON(400, gin.H{
			"mensagem": "Essa ação requere autenticação. Por favor, faça login",
		})

	}

}
