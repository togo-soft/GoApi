package models

// Go Version
type GoVersion struct {
	ID      int        `json:"-"`
	Version string     `gorm:"unique;not null" json:"version"` //版本
	Stable  bool       `json:"stable"`                         //是否为当前主线版本
	List    []GoBranch `json:"list"`
}

// Branch 版本分支信息
type GoBranch struct {
	ID          int    `json:"-"`                                 //ID
	GoVersionID int    `json:"-"`                                 //外键
	FileName    string `gorm:"size:32;not null" json:"file_name"` //名称
	Kind        string `gorm:"size:16;not null" json:"kind"`      //安装包类型
	Platform    string `gorm:"size:8;not null" json:"os"`         //平台
	Arch        string `gorm:"size:8;not null" json:"arch"`       //架构
	Size        string `gorm:"size:8;not null" json:"size"`       //大小
	CheckSum    string `gorm:"size:64;not null" json:"check_sum"` //校检
}

// GoParams 接收参数
type GoParams struct {
	Version  string `form:"version"` //版本号
	Platform string `form:"os"`      //平台
	Arch     string `form:"arch"`    //架构
	Kind     string `form:"kind"`    //包类型
	Stable   bool   `form:"stable"`  //是否稳定版
}
