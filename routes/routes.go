package routes

import (
	"github.com/ErickCelestino/Api_Rest_em_GO_Com_Gin/controllers"
	"github.com/gin-gonic/gin"
)

func HandleResquests() {
	r := gin.Default()
	r.GET("/alunos", controllers.ExibeTodosAlunos)
	r.GET("/:nome", controllers.Saudacao)
	r.Run()
}
