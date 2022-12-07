/*
 * @Author: a624669980@163.com a624669980@163.com
 * @Date: 2022-10-25 15:29:24
 * @LastEditors: a624669980@163.com a624669980@163.com
 * @LastEditTime: 2022-10-25 15:29:39
 * @FilePath: /gocron/model/common.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package model

type Result struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}
