package requesthandle

import (
	"net"
	"net/http"
	"time"
)

func newHHTTPClient() *http.Client {
	client := &http.Client{Transport: &http.Transport{
		Dial: func(netw, addr string) (net.Conn, error) {
			c, err := net.DialTimeout(netw, addr, time.Second*5) //设置建立连接超时
			if err != nil {
				return nil, err
			}
			c.SetDeadline(time.Now().Add(6 * time.Second)) //设置发送接收数据超时
			return c, nil
		},
	}}
	return client
}

// AllRequest ...
func AllRequest(getReq GetRequest, callback HandleResponse) error {
	for {
		//创建http请求
		req, err := getReq()
		if err != nil {
			return err
		}

		client := newHHTTPClient()
		resp, err := client.Do(req)
		if err != nil {
			return err
		}

		defer resp.Body.Close()
		//处理返回结果
		ok, err := callback(resp)
		if ok {
			time.Sleep(time.Second)
			continue
		}
		if err != nil {
			return err
		}

		return nil
	}
}
