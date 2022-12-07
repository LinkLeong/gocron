/*
 * @Author: a624669980@163.com a624669980@163.com
 * @Date: 2022-11-29 15:15:34
 * @LastEditors: a624669980@163.com a624669980@163.com
 * @LastEditTime: 2022-11-29 15:19:36
 * @FilePath: /gocron/route/v1/route.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package v1

import (
	"gocron/codegen"
)

type GoCron struct{}

func NewGoCron() codegen.ServerInterface {
	return &GoCron{}
}
