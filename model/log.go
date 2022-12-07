/*
 * @Author: a624669980@163.com a624669980@163.com
 * @Date: 2022-10-31 18:27:49
 * @LastEditors: a624669980@163.com a624669980@163.com
 * @LastEditTime: 2022-12-03 21:29:01
 * @FilePath: /gocron/model/log.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package model

type CornsLog struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`    //cron的名称
	CronId int    `json:"cron_id"` //cron的id
	Log    string `json:"log"`     //日志
	Time   int64  `json:"time"`    //执行时间
}

func (c *CornsLog) TableName() string {
	return "crons_log"
}
