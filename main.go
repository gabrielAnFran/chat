package main

import (
	"chat/banco"
	"chat/crud"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	banco.ConectarDB()
	defer banco.DBClient.Close()

	r.POST("/criarUsuario", crud.CriarUsuario)//cria um usuario no bd
	r.POST("/enviarMensagem", crud.EnviarMensagem)// envia mensagens do usuasrio id tal para usuario id tal, porem é dependente do login.
	r.GET("/login", crud.Logar)// confere se login e senha batem, chama jwt cria um token e passa as autorizaçoes... Estou passando como string !!!! MELHORAR !!!!
	r.GET("/getMensagens2pessoas", crud.GetMensagensDuasPessoas)//Busca conversas entre 2 pessoas
	r.GET("/getTodasMensagens", crud.GetTodasMensagens)//busca todas as conversas onde tal pessoa estava envolvida
	r.PUT("/atualizarSenha", crud.AtualizarSenha)// atualiza senha
	r.DELETE("/deletarConversas", crud.DeletarConversasPorId)// deleta todas mensagens enviadas por tal pessoa(a que está logada)

	if erro := r.Run(":5000"); erro != nil {
		log.Fatal(erro.Error())
	}

}

