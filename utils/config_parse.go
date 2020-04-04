package utils

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"sync"
)

//读取配置文件信息
type E struct {
	Environments `yaml:"environments"`
}

//请手动添加结构来实现yaml的解析
type Environments struct {
	Debug         bool   `yaml:"debug"`    //是否开启debug模式
	Server        string `yaml:"server"`   //服务运行的host:port
	DatabaseType  string `yaml:"db_type"`  //数据库类型
	DSN           string `yaml:"dsn"`      //数据库连接的源名称
	DatabaseDebug bool   `yaml:"db_debug"` //开启数据库debug模式
}

// conf 是一个全局的配置信息实例 项目运行只读取一次 是一个单例
var conf *E
var once sync.Once

// GetConfig 调用该方法会实例化conf 项目运行会读取一次配置文件 确保不会有多余的读取损耗
func GetConfig() *E {
	once.Do(func() {
		conf = new(E)
		yamlFile, err := ioutil.ReadFile("config.yaml")
		if err != nil {
			panic(err)
		}
		err = yaml.Unmarshal(yamlFile, conf)
		if err != nil {
			//读取配置文件失败,停止执行
			panic("read config file error:" + err.Error())
		}
	})
	return conf
}
