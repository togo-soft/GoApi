package usecases

const (
	StatusOK    = 200
	ClientError = 400
	ServerError = 500

	//参数缺失
	ParamsDefect = 4001
	//参数错误
	ErrorParamsParse = 4002

	//数据库出错
	ErrorDBOperation = 5001

	//远程数据出错
	ErrorRemoteRead  = 5002
	ErrorRemoteParse = 5003
)

// Response 是交付层的基本回应
type Response struct {
	Code    int         `json:"code"`    //请求状态代码
	Message interface{} `json:"message"` //请求结果提示
	Data    interface{} `json:"data"`    //请求结果与错误原因
}

// List 会返回给交付层一个列表回应
type List struct {
	Code    int         `json:"code"`    //请求状态代码
	Count   int         `json:"count"`   //数据量
	Message interface{} `json:"message"` //请求结果提示
	Data    interface{} `json:"data"`    //请求结果
}
