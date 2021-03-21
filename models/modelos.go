package models

type Users struct {
	Nome      string `json:"nome"`
	Login     string `json:"login"`
	Senha     string `json:"senha"`
}


type Mensagens struct {
	IdRemetente    int    `json:"id_remetente"`
	IdDestinatario int    `json:"id_destinatario"`
	Mensagem       string `json:"mensagem"`
}

type LerMensagens struct {
	IdRemetente    int `json:"id_1"`
	IdDestinatario int `json:"id_2"`
}

type AtualizarSenha struct {
	IdUsuario uint   `json:"id_usuario"`
	Login     string `json:"login"`
	Senha     string `json:"senha"`
	NovaSenha string `json:"novasenha"`
}
