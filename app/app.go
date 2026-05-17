package app

import (
	"gin-learning/database"

	"github.com/gin-gonic/gin"
)

type App struct {
	Router *gin.Engine
	Models database.ModelRegistery
}

func (a *App) Run() {

	a.Router.Run(":3000")
}
