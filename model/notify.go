/*
 * @Author: a624669980@163.com a624669980@163.com
 * @Date: 2022-10-21 17:52:37
 * @LastEditors: a624669980@163.com a624669980@163.com
 * @LastEditTime: 2022-11-25 18:03:17
 * @FilePath: /gocron/model/notify.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package model

type Notify struct {
	Id      int
	Name    string
	Type    string //feishu,
	Title   string //通知标题
	Content string
	Url     string //通知地址
}

func (Notify) TableName() string {
	return "notify"
}
