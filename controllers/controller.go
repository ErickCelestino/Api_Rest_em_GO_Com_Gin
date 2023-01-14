package controllers

import (
	"github.com/ErickCelestino/Api_Rest_em_GO_Com_Gin/models"
	"github.com/gin-gonic/gin"
)

func ExibeTodosAlunos(c *gin.Context) {
	c.JSON(200, models.Alunos)
}

func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API diz: ": "E ai " + nome + ", tudo beleza?",
	})

}
