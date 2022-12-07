/*
 * @Author: a624669980@163.com a624669980@163.com
 * @Date: 2022-11-25 18:07:05
 * @LastEditors: a624669980@163.com a624669980@163.com
 * @LastEditTime: 2022-11-28 15:41:15
 * @FilePath: /gocron/repository/notify.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package repository

import (
	"gocron/model"

	"gorm.io/gorm"
)

type NotifyRepository interface {
	GetNotifyList() []model.Notify
	AddNotify(m model.Notify) error
	UpdateNotify(m model.Notify) error
	GetNOtifyById(id string) model.Notify
	DeleteNotify(id string) error
}

type notifyRepository struct {
	db *gorm.DB
}

func (n *notifyRepository) AddNotify(m model.Notify) error {
	return n.db.Create(&m).Error
}

func (n *notifyRepository) GetNotifyList() []model.Notify {
	m := []model.Notify{}
	n.db.Find(&m)
	return m
}
func (n *notifyRepository) UpdateNotify(m model.Notify) error {
	return n.db.Save(&m).Error
}
func (n *notifyRepository) GetNOtifyById(id string) model.Notify {
	m := model.Notify{}
	n.db.First(&m, id)
	return m
}
func (n *notifyRepository) DeleteNotify(id string) error {
	m := model.Notify{}
	return n.db.Delete(&m, id).Error
}

func NewNotifyRepository(db *gorm.DB) NotifyRepository {
	return &notifyRepository{db: db}
}
