/*
 * @Author: a624669980@163.com a624669980@163.com
 * @Date: 2022-11-09 17:47:01
 * @LastEditors: a624669980@163.com a624669980@163.com
 * @LastEditTime: 2022-11-28 15:49:52
 * @FilePath: /gocron/service/notify.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package service

import (
	"gocron/model"
	"gocron/repository"
)

type NotifyService interface {
	GetNotifyList() []model.Notify
	AddNotify(m model.Notify) error
	UpdateNotify(m model.Notify) error
	GetNOtifyById(id string) model.Notify
	DeleteNotify(id string) error
}

type notifyService struct {
	rp repository.Repository
}

func (n *notifyService) AddNotify(m model.Notify) error {
	return n.rp.Notify().AddNotify(m)
}

func (n *notifyService) GetNotifyList() []model.Notify {
	return n.rp.Notify().GetNotifyList()
}
func (n *notifyService) UpdateNotify(m model.Notify) error {
	return n.rp.Notify().UpdateNotify(m)
}
func (n *notifyService) GetNOtifyById(id string) model.Notify {
	return n.rp.Notify().GetNOtifyById(id)
}
func (n *notifyService) DeleteNotify(id string) error {
	return n.rp.Notify().DeleteNotify(id)
}

func NewNotifyService(rp repository.Repository) NotifyService {
	return &notifyService{rp: rp}
}
