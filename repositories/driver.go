package repositories

import (
	"GoApi/models"
	"GoApi/utils"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	// sqlite数据库-连接引擎
	engine *gorm.DB
	//配置文件句柄
	conf = utils.GetConfig()
)

func init() {
	var err error
	if engine, err = gorm.Open(conf.DatabaseType, conf.DSN); err != nil {
		panic("connect " + conf.DatabaseType + " error:" + err.Error())
	}
	//数据库连通性检测
	if err = engine.DB().Ping(); err != nil {
		panic("ping " + conf.DatabaseType + " error:" + err.Error())
	}
	// 全局禁用表名复数
	engine.SingularTable(true)
	//是否开启调试
	if !conf.DatabaseDebug {
		engine.LogMode(false)
	}
	//同步数据库结构
	engine.AutoMigrate(new(models.GoVersion), new(models.GoBranch))
}
