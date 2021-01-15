package util

import (
	"bytes"
	contexts "context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gotcb/tencent/context"
)

const (
	//触发云函数
	invokeCloudFunctionURL = "https://tcb-api.tencentcloudapi.com/api/v2/envs/"
)

//Util 结构体
type Util struct {
	Context *context.Context
}

//NewUtil 创建tcb
func NewUtil(ct *context.Context) *Util {
	return &Util{
		Context: ct,
	}
}

//SendPostURL 发送post请求*******************************************************************公共函数
func (tcb *Util) SendPostURL(method string, url string, data []byte) (string, error) {

	buffer := bytes.NewBuffer(data)

	request, err := http.NewRequest(method, url, buffer)
	if err != nil {
		fmt.Printf("http.NewRequest%v", err)
		return "", err
	}

	request.Header.Add("Vary", "Origin")
	request.Header.Add("Access-Control-Allow-Origin", "*")                            //允许访问所有域
	request.Header.Add("Access-Control-Allow-Headers", "*")                           //header的类型
	request.Header.Add("content-type", "application/json;charset=utf-8")              //返回数据格式是json
	request.Header.Add("Access-Control-Allow-Methods", "PUT,POST,GET,DELETE,OPTIONS") //返回数据格式是json

	request.Header.Add("X-CloudBase-Authorization", tcb.Context.Authorization)
	request.Header.Add("X-CloudBase-TimeStamp", strconv.FormatInt(tcb.Context.Timestamp, 10))
	request.Header.Add("content-type", "application/json;charset=utf-8") //返回数据格式是json

	if err != nil {
		fmt.Printf("http.NewRequest%v", err)
		return "", err
	}
	client := http.Client{}                                      //创建客户端
	resp, err := client.Do(request.WithContext(contexts.TODO())) //发送请求
	if err != nil {
		fmt.Println(err)
		//添加一条失败日志
		return "", err
	}

	defer resp.Body.Close()
	respBytes, err := ioutil.ReadAll(resp.Body)

	return string(respBytes), err
}
