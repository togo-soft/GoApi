package usecases

import (
	"GoApi/models"
	"GoApi/repositories"
	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

var versionRepo = repositories.NewGoVersionRepo()

type GoVersionUC struct {
}

func NewGoVersionUC() GoVersionInterface {
	return &GoVersionUC{}
}

// Fetch 查询数据
func (g *GoVersionUC) Fetch(ctx *gin.Context) (int, *List) {
	var params = &models.GoParams{}
	if err := ctx.Bind(params); err != nil {
		return ClientError, &List{
			Code:    ErrorParamsParse,
			Message: "error:" + err.Error(),
		}
	}
	if data, err := versionRepo.Fetch(params); err != nil {
		return ServerError, &List{
			Code:    ErrorDBOperation,
			Message: "error:" + err.Error(),
		}
	} else {
		return StatusOK, &List{
			Code:    StatusOK,
			Count:   len(data),
			Message: "ok",
			Data:    data,
		}
	}
}

// Update 更新数据库中的版本数据
func (g *GoVersionUC) Update(ctx *gin.Context) (int, *Response) {
	//更新数据前 先刷新数据表
	_ = versionRepo.Reset()
	//抓取数据
	if err := grab(); err != nil {
		return ServerError, &Response{
			Code:    ErrorDBOperation,
			Message: err,
			Data:    nil,
		}
	}
	return StatusOK, &Response{
		Code:    StatusOK,
		Message: "ok",
		Data:    nil,
	}
}

// grab 抓取页面
func grab() error {
	// Request the HTML page.
	res, err := http.Get("https://golang.google.cn/dl")
	if err != nil || res.StatusCode != 200 {
		log.Fatal("请求出错:", err, "状态码:", res.StatusCode)
		return err
	}
	defer res.Body.Close()
	// 载入页面 分析
	dom, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return err
	}
	//获取的stable版本
	dom.Find(".toggleVisible").Each(func(i int, stable *goquery.Selection) {
		vers, _ := stable.Attr("id")
		//截取版本号
		version := strings.ReplaceAll(vers, "go", "")
		//版本获取完毕 开始获取各分支信息
		var v = &models.GoVersion{Version: version, Stable: true}
		stable.Find("tbody").Each(func(i int, tbody *goquery.Selection) {
			//进入tbody中
			tbody.Find("tr").Each(func(i int, tr *goquery.Selection) {
				//从进入tr 检测 Other Ports 列
				if first, exist := tr.Attr("class"); first == "first" && exist {
					return
				}
				var b = models.GoBranch{}
				tr.Find("td").Each(func(i int, td *goquery.Selection) {
					if i == 0 {
						//获取文件名
						b.FileName = td.Find("a").Text()
					} else if i == 1 {
						//获取包类型
						b.Kind = td.Text()
					} else if i == 2 {
						//获取平台
						b.Platform = td.Text()
					} else if i == 3 {
						//获取架构
						b.Arch = td.Text()
					} else if i == 4 {
						//获取文件大小
						b.Size = td.Text()
					} else if i == 5 {
						//获取检验和
						b.CheckSum = td.Find("tt").Text()
					}
				})
				//数据装载
				v.List = append(v.List, b)
			})
			//数据入数据库
			_ = versionRepo.Insert(v)
		})
	})
	//获取归档版本
	dom.Find("#archive").Find(".toggle").Each(func(i int, archive *goquery.Selection) {
		vers, _ := archive.Attr("id")
		//截取版本号
		version := strings.ReplaceAll(vers, "go", "")
		//版本获取完毕 开始获取各分支信息
		var v = &models.GoVersion{Version: version}
		//var b = []models.Branch{}
		archive.Find("tbody").Each(func(i int, tbody *goquery.Selection) {
			//进入tbody中
			tbody.Find("tr").Each(func(i int, tr *goquery.Selection) {
				//从进入tr
				var b = models.GoBranch{}
				tr.Find("td").Each(func(i int, td *goquery.Selection) {
					if i == 0 {
						//获取文件名
						b.FileName = td.Find("a").Text()
					} else if i == 1 {
						//获取包类型
						b.Kind = td.Text()
					} else if i == 2 {
						//获取平台
						b.Platform = td.Text()
					} else if i == 3 {
						//获取架构
						b.Arch = td.Text()
					} else if i == 4 {
						//获取文件大小
						b.Size = td.Text()
					} else if i == 5 {
						//获取检验和
						b.CheckSum = td.Find("tt").Text()
					}
				})
				//数据装载
				v.List = append(v.List, b)
			})
			//数据入数据库
			_ = versionRepo.Insert(v)
		})
	})
	return nil
}
