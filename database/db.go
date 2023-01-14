package database

import (
	"log"

	"github.com/ErickCelestino/Api_Rest_em_GO_Com_Gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	stringDeConexao := "host=localhost user=root password=root dbname=root port=5436 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Panic("Erro ao conectar com o banco de dados")
	}
	//Gerando os dados no banco de dados
	DB.AutoMigrate(&models.Aluno{})
}
