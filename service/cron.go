/*
 * @Author: a624669980@163.com a624669980@163.com
 * @Date: 2022-10-21 16:41:38
 * @LastEditors: a624669980@163.com a624669980@163.com
 * @LastEditTime: 2022-12-06 22:20:33
 * @FilePath: /gocron/service/cron.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package service

import (
	"errors"
	"fmt"
	"gocron/model"
	"gocron/pkg/tools"
	"gocron/repository"
	"strconv"
	"time"

	"github.com/robfig/cron/v3"
)

type CronService interface {
	AddJob(m model.Corns) error
	DeleteJob(id string) error
	GetJobById(id string)
	GetJobs() ([]model.Corns, error)
	GetJobLog(id string) ([]model.CornsLog, error)
	UpdateStatus(id, enable string) error
	GetCronByDB()
}

type cronStruct struct {
	rp   repository.Repository
	cron *cron.Cron
}

func (c *cronStruct) GetCronByDB() {
	list, _ := MyService.Cron().GetJobs()
	// fmt.Println(list, err)
	for i := 0; i < len(list); i++ {
		aaa := list[i]
		//c.cron.AddJob("", cron.Job.Job)
		if aaa.Enable == "enable" {
			cronId, _ := c.cron.AddFunc(aaa.Time, func() {
				tools.Command(aaa.Command)
				if aaa.Notify > 0 {
					notify := c.rp.Notify().GetNOtifyById(strconv.Itoa(aaa.Notify))
					if notify.Id > 0 {
						result, err := tools.SendNotify(notify.Url, notify.Type, notify.Content)
						if err != nil {
							fmt.Println(err)
						}
						fmt.Println(result)
					}
				}
			})
			c.rp.Corn().UpdateCronIdById(strconv.Itoa(list[i].Id), strconv.FormatInt(int64(cronId), 10))
		} else if aaa.Enable == "unable" {
			fmt.Println("cron id:", aaa.Id, "\n this cron enable status is unable")
		} else {
			fmt.Println("enable parameter error")
		}
		c.cron.Start()
	}
}

func (c *cronStruct) GetJobLog(id string) ([]model.CornsLog, error) {
	m, err := c.rp.Corn().GetJobById(id)
	if err != nil {
		return nil, err
	}
	list, err := c.rp.Corn().GetJobLog(m.Name)
	if err != nil {
		return nil, err
	}
	return list, nil
}
func (c *cronStruct) GetJobLogByCronId(cron_id int) ([]model.CornsLog, error) {
	list, err := c.rp.Corn().GetCronJobsByCronId(cron_id)
	if err != nil {
		return nil, err
	}
	return list, nil
}
func (c *cronStruct) AddJob(m model.Corns) error {

	cronId, err := c.cron.AddFunc(m.Time, func() {
		l, err := tools.ExecShell(m.Command)
		if err != nil {
			fmt.Println(err)
		}
		logM := model.CornsLog{
			Name:   m.Name,
			Log:    l,
			CronId: m.Id,
			Time:   time.Now().Unix(),
		}
		c.rp.Corn().CreateCronLog(logM)
		if m.Notify > 0 {
			notify := c.rp.Notify().GetNOtifyById(strconv.Itoa(m.Notify))
			if notify.Id > 0 {
				result, err := tools.SendNotify(notify.Url, notify.Type, notify.Content)
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println(result)
			}
		}
	})
	if err != nil {
		return err
	}
	c.cron.Start()
	m.CronID = int(cronId)
	m.Enable = "enable"
	return c.rp.Corn().AddJob(&m)
}

func (c *cronStruct) DeleteJob(id string) error {
	m, err := c.rp.Corn().GetJobById(id)
	if err != nil || m.CronID == 0 {
		return errors.New("cronId is null")
	}
	c.cron.Remove(cron.EntryID(m.CronID))
	c.cron.Start()
	c.rp.Corn().DeleteJob(id)
	return nil
}
func (c *cronStruct) GetJobById(id string) {
	c.rp.Corn().GetJobById(id)
}
func (c *cronStruct) GetJobs() ([]model.Corns, error) {
	return c.rp.Corn().GetJobs()
}

func (c *cronStruct) UpdateStatus(id string, enable string) error {
	m, err := c.rp.Corn().GetJobById(id)
	if err != nil || m.CronID == 0 {
		return errors.New("cronId is null")
	}
	if enable == "unable" {
		c.cron.Remove(cron.EntryID(m.CronID))
		c.rp.Corn().UpdateStatus(id, enable)
	} else if enable == "enable" {
		c.cron.AddFunc(m.Time, func() {
			l, err := tools.ExecShell(m.Command)
			if err != nil {
				fmt.Println(err)
			}
			logM := model.CornsLog{
				Name:   m.Name,
				Log:    l,
				CronId: m.Id,
				Time:   time.Now().Unix(),
			}
			c.rp.Corn().CreateCronLog(logM)
			if m.Notify > 0 {
				notify := c.rp.Notify().GetNOtifyById(strconv.Itoa(m.Notify))
				if notify.Id > 0 {
					result, err := tools.SendNotify(notify.Url, notify.Type, notify.Content)
					if err != nil {
						fmt.Println(err)
					}
					fmt.Println(result)
				}
			}
		})
		c.cron.Start()
		c.rp.Corn().UpdateStatus(id, enable)

	} else {
		return errors.New("enable should be enable or unable")
	}
	return nil
}

func NewCronService(rp repository.Repository, cron *cron.Cron) CronService {
	return &cronStruct{rp: rp, cron: cron}
}
