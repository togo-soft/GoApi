package repositories

import (
	"GoApi/models"
	"errors"
	"reflect"
)

type GoVersionRepo struct {
}

func NewGoVersionRepo() GoVersionInterface {
	return &GoVersionRepo{}
}

// 添加数据
func (g *GoVersionRepo) Insert(record *models.GoVersion) error {
	return engine.Create(record).Error
}

// Reset 删除数据 删表重建
func (g *GoVersionRepo) Reset() error {
	if err := engine.DropTableIfExists(&models.GoVersion{}, &models.GoBranch{}).Error; err != nil {
		return err
	}
	if err := engine.CreateTable(&models.GoVersion{}, &models.GoBranch{}).Error; err != nil {
		return err
	}
	return nil
}

// Fetch 查询数据
func (g *GoVersionRepo) Fetch(params *models.GoParams) ([]models.GoVersion, error) {
	var list []models.GoVersion
	if reflect.DeepEqual(params, &models.GoParams{}) {
		err := engine.Preload("List").Find(&list).Error
		return list, err
	}
	if params.Version != "" && params.Stable {
		//筛选版本
		engine.Where("version = ? AND stable = ?", params.Version, params.Stable).Find(&list)
	} else if params.Version != "" {
		engine.Where("version = ? ", params.Version).Find(&list)
	} else if params.Stable {
		engine.Where("stable = ?", params.Stable).Find(&list)
	} else {
		//不筛选版本
		engine.Find(&list)
	}
	//数据项为空
	if len(list) == 0 {
		return nil, errors.New("找不到该版本号:" + params.Version)
	}
	//从版本中选择下属分支
	var filter = make(map[string]interface{}, 4)
	if params.Platform != "" {
		filter["platform"] = params.Platform
	}
	if params.Kind != "" {
		filter["kind"] = params.Kind
	}
	if params.Arch != "" {
		filter["arch"] = params.Arch
	}
	for i := 0; i < len(list); i++ {
		engine.Where(filter).Model(&list[i]).Related(&list[i].List, "go_branch")
	}
	return list, nil
}
