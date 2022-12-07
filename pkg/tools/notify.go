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
)

type FeishuModel struct {
	MsgType string       `json:"msg_type"`
	Content ContentModel `json:"content"`
}
type ContentModel struct {
	Text string `json:"text"`
}

func SendNotify(url string, t string, message string) ([]byte, error) {
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
	}
	return nil, nil
}
