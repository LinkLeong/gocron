/*
 * @Author: a624669980@163.com a624669980@163.com
 * @Date: 2022-11-14 17:32:27
 * @LastEditors: a624669980@163.com a624669980@163.com
 * @LastEditTime: 2022-11-25 17:52:57
 * @FilePath: /gocron/pkg/tools/http_helper.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package tools

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/go-resty/resty/v2"
)

func HttpDo(method, path string, head map[string]string, data []byte) ([]byte, error) {
	client := &http.Client{}

	body1 := bytes.NewBuffer(data)
	//把form数据编下码
	req, err := http.NewRequest(strings.ToUpper(method), path, body1)
	if err != nil {
		return nil, err
	}
	for k, v := range head {
		req.Header.Add(k, v)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(body))
	return body, nil
}

func HttpNewRequest(method, url string, payload *strings.Reader) ([]byte, error) {
	// url := "https://open.feishu.cn/open-apis/bot/v2/hook/21980b82-3c60-4636-b6d2-04a7822a0185"
	// method := "POST"

	// payload := strings.NewReader(`{
	// 	"msg_type": "text",
	// 	"content": {
	// 		"text": "测试机器人"
	// 	}
	// }`)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	// req.Header.Add("User-Agent", "apifox/1.0.0 (https://www.apifox.cn)")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	fmt.Println(string(body))
	return body, nil
}

type RestyClient struct {
	client *resty.Client
}

func NewRestyClient() *RestyClient {
	return &RestyClient{
		client: resty.New(),
	}
}

func (rc *RestyClient) RestyGet(url string) (*resty.Response, error) {
	req := rc.client.R()
	return req.Get(url)
}

func (rc *RestyClient) RestyPost(url string) (*resty.Response, error) {
	req := rc.client.R()
	return req.Post(url)
}
