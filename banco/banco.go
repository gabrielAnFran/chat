package banco

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // GORM
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "chatgo"
)

//DBClient é o export do bd
var DBClient *gorm.DB

//IniciarMigracaoBD Inicia conexao com bando
func ConectarDB() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, erro := gorm.Open("postgres", psqlInfo)
	if erro != nil {
		fmt.Println(erro.Error())
		panic("Falha ao conectar ao banco de dados")
	}

	DBClient = db

}
