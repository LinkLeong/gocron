/*
 * @Author: a624669980@163.com a624669980@163.com
 * @Date: 2022-11-15 17:48:57
 * @LastEditors: a624669980@163.com a624669980@163.com
 * @LastEditTime: 2022-12-06 22:22:48
 * @FilePath: /gocron/pkg/tools/notify_method/feishu.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package tools

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
)

type FeishuModel struct {
	MsgType string       `json:"msg_type"`
	Content ContentModel `json:"content"`
}
type ContentModel struct {
	Text string `json:"text"`
}

func SendNotify(url string, t string, message string) ([]byte, error) {
	client := resty.New()
	if t == "feishu" {
		data := FeishuModel{
			MsgType: "text",
			Content: ContentModel{
				Text: message,
			},
		}
		payloadBytes, err := json.Marshal(data)
		if err != nil {
			return nil, err
		}
		header := make(map[string]string)
		header["Content-Type"] = "application/json"
		return HttpDo("POST", url, header, payloadBytes)
	} else if t == "serverchan" {
		resp, err := client.R().
			SetHeader("Content-Type", "application/json").
			SetBody(map[string]string{
				"title": "testSeverChan",
				"desp":  message,
			}).
			Post(url)
		fmt.Println(resp.StatusCode())
		fmt.Println(resp.String())
		fmt.Println(err)
		return resp.Body(), err
	} else if t == "pushplus" {
		resp, err := client.R().
			SetHeader("Content-Type", "application/json").
			SetBody(map[string]string{
				"topic":   "123",
				"token":   "ae4f0531a0c2491b9be7c38a0f354150",
				"title":   "12",
				"content": message,
				// "template": "html",
				"channel": "wechat",
			}).
			Post(url)
		fmt.Println(resp.StatusCode())
		fmt.Println(resp.String())
		fmt.Println(err)
		return resp.Body(), err
	}
	return nil, nil
}
