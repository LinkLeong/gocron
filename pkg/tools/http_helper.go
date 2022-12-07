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
