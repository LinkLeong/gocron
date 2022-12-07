/*
 * @Author: a624669980@163.com a624669980@163.com
 * @Date: 2022-11-28 14:50:22
 * @LastEditors: a624669980@163.com a624669980@163.com
 * @LastEditTime: 2022-11-28 15:10:35
 * @FilePath: /gocron/repository/repository.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package repository

var MyRepository Repository

type Repository interface {
	Notify() NotifyRepository

	Corn() CronRepository
}

func NewRepository() Repository {

	db := GetDb()

	return &store{
		notify: NewNotifyRepository(db),
		cron:   NewCronRepository(db),
	}
}

type store struct {
	notify NotifyRepository
	cron   CronRepository
}

func (c *store) Notify() NotifyRepository {
	return c.notify
}

func (c *store) Corn() CronRepository {
	return c.cron
}
