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

///mongodb+srv://diattamohamet30_db_user:9ErMiqc0oBOjN1n5@cluster0.wdabvqz.mongodb.net/?appName=Cluster0
