package usecases

import "github.com/gin-gonic/gin"

type GoVersionInterface interface {
	//条件查找
	Fetch(*gin.Context) (int, *List)
	//更新表信息
	Update(*gin.Context) (int, *Response)
}
