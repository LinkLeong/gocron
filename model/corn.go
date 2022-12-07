/*
 * @Author: a624669980@163.com a624669980@163.com
 * @Date: 2022-10-21 14:57:52
 * @LastEditors: a624669980@163.com a624669980@163.com
 * @LastEditTime: 2022-12-04 22:42:24
 * @FilePath: /gocron/model/corn.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package model

type Corns struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Type    int    `json:"type"`
	Command string `json:"command"` //可以是url,也可以是shell脚本的地址
	CronID  int    `json:"cron_id"` //cron的id
	Status  int    `json:"status"`  //0:未执行
	Notify  int    `json:"notify"`  //0:不通知,(通知表的id)
	Tags    string `json:"tags"`    //标签
	Time    string `json:"time"`    //执行时间
}

func (c *Corns) TableName() string {
	return "crons"
}
