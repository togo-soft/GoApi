package repositories

import "GoApi/models"

type GoVersionInterface interface {
	//添加数据
	Insert(*models.GoVersion) error
	//删除数据
	Reset() error
	//条件查询数据
	Fetch(*models.GoParams) ([]models.GoVersion, error)
}
