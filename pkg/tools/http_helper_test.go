/*
 * @Author: a624669980@163.com a624669980@163.com
 * @Date: 2022-11-14 19:01:33
 * @LastEditors: a624669980@163.com a624669980@163.com
 * @LastEditTime: 2022-11-25 18:03:57
 * @FilePath: /gocron/pkg/tools/http_helper_test.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package tools

import (
	"encoding/json"
	"fmt"
	"gocron/model"
	"testing"
)

func TestHttpDo(t *testing.T) {

	notify := model.Notify{}
	err := json.Unmarshal([]byte(`{
		"name": "飞书",
		"type": "post",
		"title": "标题",
		"body": [
		{
			"msg_type":"text",
			"content":{"text":"请求了baidu"}
		}
	],
		"url": "https://open.feishu.cn/open-apis/bot/v2/hook/f537994e-c5ca-4f13-b338-3e748a1a9d73",
		"header": [
		{
			"Content-Type": "application/json"
		}
	]
	}`), &notify)
	if err != nil {
		fmt.Println("1")
		fmt.Println(err)
	}

	//result, err := HttpDo(notify.Type, notify.Url, notify.Header, notify.Body)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//fmt.Println(result)
}
