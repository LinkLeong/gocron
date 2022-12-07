/*
 * @Author: a624669980@163.com a624669980@163.com
 * @Date: 2022-10-31 15:20:06
 * @LastEditors: a624669980@163.com a624669980@163.com
 * @LastEditTime: 2022-10-31 15:23:00
 * @FilePath: /gocron/pkg/tools/command.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package tools

import (
	"bufio"
	"fmt"
	"os/exec"
)

func Command(cmdStr string) ([]string, error) {
	cmd := exec.Command("/bin/bash", "-c", cmdStr)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
		return []string{}, err
	}
	defer stdout.Close()
	if err = cmd.Start(); err != nil {
		fmt.Println(err)
		return []string{}, err
	}
	// str, err := ioutil.ReadAll(stdout)
	list := []string{}
	outputBuf := bufio.NewReader(stdout)
	for {
		output, _, err := outputBuf.ReadLine()
		if err != nil {
			if err.Error() != "EOF" {
				fmt.Printf("Error :%s\n", err)
			}
			break
		}
		list = append(list, string(output))
	}
	cmd.Wait()
	return list, nil
}
