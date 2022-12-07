/*
 * @Author: a624669980@163.com a624669980@163.com
 * @Date: 2022-10-31 18:20:10
 * @LastEditors: a624669980@163.com a624669980@163.com
 * @LastEditTime: 2022-10-31 18:20:27
 * @FilePath: /gocron/pkg/tools/shell.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package tools

import (
	"bytes"
	"os/exec"
	"strings"
)

func ExecShell(s string) (string, error) {
	cmd := exec.Command("/bin/bash", "-c", s)
	cmd.Stdin = strings.NewReader(s)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return out.String(), nil
}
