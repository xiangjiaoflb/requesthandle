package requesthandle

import "net/http"

//记录日志时的key
const (
	//请求的数据
	RequestData = "requestdata"

	//返回的数据
	ResponseData = "responsedata"
)

// GetRequest 请求的结构
type GetRequest func() (*http.Request, error)

// HandleResponse 处理返回的结构
type HandleResponse func(resp *http.Response) (bool, error)
