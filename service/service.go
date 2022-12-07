/*
 * @Author: a624669980@163.com a624669980@163.com
 * @Date: 2022-10-21 16:05:12
 * @LastEditors: a624669980@163.com a624669980@163.com
 * @LastEditTime: 2022-11-28 15:15:45
 * @FilePath: /gocron/service/service.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package service

import (
	"gocron/repository"

	"github.com/robfig/cron/v3"
	"gorm.io/gorm"
)

var MyService Service

type Service interface {
	Cron() CronService
	Notify() NotifyService
}

func NewService(repository repository.Repository, cron *cron.Cron) Service {

	return &store{
		cron:   NewCronService(repository, cron),
		notify: NewNotifyService(repository),
	}
}

type store struct {
	db     *gorm.DB
	cron   CronService
	notify NotifyService
}

func (c *store) Cron() CronService {
	return c.cron
}

func (c *store) Notify() NotifyService {
	return c.notify
}
