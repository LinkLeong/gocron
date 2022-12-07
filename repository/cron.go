/*
 * @Author: a624669980@163.com a624669980@163.com
 * @Date: 2022-11-28 15:08:27
 * @LastEditors: a624669980@163.com a624669980@163.com
 * @LastEditTime: 2022-12-04 23:08:59
 * @FilePath: /gocron/repository/cron.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package repository

import (
	"gocron/model"

	"gorm.io/gorm"
)

type CronRepository interface {
	CreateCronLog(m model.CornsLog) error
	AddJob(m *model.Corns) error
	DeleteJob(id string) error
	GetJobById(id string) (model.Corns, error)
	UpdateJob() error
	GetJobs() ([]model.Corns, error)
	GetCronJobsByCronId(cron_id int) ([]model.CornsLog, error)
	GetJobLog(id string) ([]model.CornsLog, error)
}

type cronStruct struct {
	db *gorm.DB
}

func (c *cronStruct) CreateCronLog(m model.CornsLog) error {
	return c.db.Create(&m).Error
}
func (c *cronStruct) GetJobLog(name string) ([]model.CornsLog, error) {
	var list []model.CornsLog
	tx := c.db.Where("name= ?", name).Find(&list)
	return list, tx.Error
}
func (c *cronStruct) AddJob(m *model.Corns) error {
	return c.db.Create(m).Error
}

func (c *cronStruct) DeleteJob(id string) error {
	return c.db.Model(&model.Corns{}).Delete("id = ?", id).Error
}
func (c *cronStruct) GetJobById(id string) (model.Corns, error) {
	m := model.Corns{}
	tx := c.db.Where("id= ?", id).First(&m)
	return m, tx.Error
}
func (c *cronStruct) GetCronJobsByCronId(cron_id int) ([]model.CornsLog, error) {
	list := []model.CornsLog{}
	tx := c.db.Where("cron_id= ?", cron_id).Find(&list)
	return list, tx.Error
}

//
func (c *cronStruct) GetJobs() ([]model.Corns, error) {
	var list []model.Corns
	tx := c.db.Find(&list)
	return list, tx.Error
}

func (c *cronStruct) UpdateJob() error {
	return nil
}

func NewCronRepository(db *gorm.DB) CronRepository {
	return &cronStruct{db: db}
}
