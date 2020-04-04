package GoVersion

import (
	"GoApi/usecases"
	"github.com/gin-gonic/gin"
)

var version = usecases.NewGoVersionUC()

func Get(this *gin.Context) {
	this.JSON(version.Fetch(this))
}

func Update(this *gin.Context) {
	this.JSON(version.Update(this))
}
