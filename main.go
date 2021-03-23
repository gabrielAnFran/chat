package main

import (
	"chat/banco"
	"chat/crud"
	"chat/middlewares"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	banco.ConectarDB()
	defer banco.DBClient.Close()

Auth := r.Group("") //não precisa de autenticação
Api := r.Group("") //precisam de autenticação
Api.Use(middlewares.JwtMicroservices())

	Auth.POST("/criarUsuario", crud.CriarUsuario)//cria um usuario no bd
	Api.POST("/enviarMensagem", crud.EnviarMensagem)// envia mensagens do usuasrio id tal para usuario id tal, porem é dependente do login.
	Auth.GET("/login", crud.Logar)// confere se login e senha batem, chama jwt cria um token e passa as autorizaçoes... Estou passando como string !!!! MELHORAR !!!!
	Api.GET("/getMensagens2pessoas", crud.GetMensagensDuasPessoas)//Busca conversas entre 2 pessoas
	Api.GET("/getTodasMensagens", crud.GetTodasMensagens)//busca todas as conversas onde tal pessoa estava envolvida
	Api.PUT("/atualizarSenha", crud.AtualizarSenha)// atualiza senha
	Api.DELETE("/deletarConversas", crud.DeletarConversasPorId)// deleta todas mensagens enviadas por tal pessoa(a que está logada)

	if erro := r.Run(":5000"); erro != nil {
		log.Fatal(erro.Error())
	}

}

