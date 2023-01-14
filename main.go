package main

import (
	"github.com/ErickCelestino/Api_Rest_em_GO_Com_Gin/models"
	"github.com/ErickCelestino/Api_Rest_em_GO_Com_Gin/routes"
)

func main() {
	//Mock
	models.Alunos = []models.Aluno{
		{Nome: "Erick", CPF: "000000000001", RG: "1000000000"},
		{Nome: "Ana", CPF: "000000000002", RG: "2000000000"},
	}
	routes.HandleResquests()
}
