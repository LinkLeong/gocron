/*
 * @Author: a624669980@163.com a624669980@163.com
 * @Date: 2022-10-24 16:38:07
 * @LastEditors: a624669980@163.com a624669980@163.com
 * @LastEditTime: 2022-12-04 22:41:40
 * @FilePath: /gocron/route/v1/cron.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package v1

import (
	"encoding/json"
	"gocron/codegen"
	"gocron/model"
	"gocron/service"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/labstack/echo/v4"
)

// 获取定时任务列表
func GetCronList(w http.ResponseWriter, r *http.Request) {
	// list := service.Service.Cron().GetJobs()
	// result := model.Result{Code: 200, Msg: "success", Data: list}
	w.Header().Set("Content-Type", "application/json")
	//json.NewEncoder(w).Encode(result)
}

// 添加定时任务
func PostCronAdd(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	m := model.Corns{}
	dataByte, err := ioutil.ReadAll(r.Body)
	if err != nil {
		result := model.Result{Code: 500, Msg: "error", Data: err}
		json.NewEncoder(w).Encode(result)
		return
	}
	err = json.Unmarshal(dataByte, &m)
	if err != nil {
		result := model.Result{Code: 500, Msg: "error", Data: err}
		json.NewEncoder(w).Encode(result)
		return
	}
	err = service.MyService.Cron().AddJob(m)
	if err != nil {
		result := model.Result{Code: 500, Msg: "error", Data: err}
		json.NewEncoder(w).Encode(result)
		return
	}

	json.NewEncoder(w).Encode(json.RawMessage(string(`{"code": 200, "msg": " 任务添加成功"}`)))
}

// 删除定时任务
func DeleteCronDelete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	err := service.MyService.Cron().DeleteJob(id)
	if err != nil {
		result := model.Result{Code: 500, Msg: "error", Data: err}
		json.NewEncoder(w).Encode(result)
		return
	}

	json.NewEncoder(w).Encode(json.RawMessage(string(`{"code": 200, "msg": "任务删除成功"}`)))
}

func (g *GoCron) GetCronJobs(ctx echo.Context) error {
	list, err := service.MyService.Cron().GetJobs()
	message := "OK"
	if err != nil {
		message = err.Error()
		return ctx.JSON(http.StatusInternalServerError, codegen.BadRequest{Message: message})
	}
	crons := []codegen.Cron{}
	for _, v := range list {
		id := v.Id
		crons = append(crons, codegen.Cron{
			Id:      &id,
			Name:    v.Name,
			Command: v.Command,
			Time:    v.Time,
		})
	}
	return ctx.JSON(http.StatusOK, codegen.GetCronJobsOK{Message: &message, Data: &crons})
}
func (g *GoCron) CreateCronJob(ctx echo.Context) error {
	m := codegen.Cron{}
	message := "OK"
	if err := ctx.Bind(&m); err != nil {
		message = err.Error()
		return ctx.JSON(http.StatusBadRequest, codegen.BadRequest{Message: message})
	}

	corn := model.Corns{}
	corn.Command = m.Command
	corn.Name = m.Name
	corn.Time = m.Time
	corn.Notify = m.Notify
	err := service.MyService.Cron().AddJob(corn)
	if err != nil {
		message = err.Error()
		return ctx.JSON(http.StatusInternalServerError, codegen.BadRequest{Message: message})
	}
	return ctx.JSON(http.StatusOK, codegen.OKRequest{Message: &message})
}
func (g *GoCron) DeleteCronJob(ctx echo.Context, id int) error {
	message := "OK"
	err := service.MyService.Cron().DeleteJob(strconv.Itoa(id))
	if err != nil {
		message = err.Error()
		ctx.JSON(http.StatusInternalServerError, codegen.BadRequest{Message: message})
	}
	return nil
}
func (g *GoCron) UpdateCronJobStatus(ctx echo.Context, id int) error {
	m := codegen.CronStatus{}
	message := "OK"
	if err := ctx.Bind(&m); err != nil {
		message = err.Error()
		return ctx.JSON(http.StatusBadRequest, codegen.BadRequest{Message: message})
	}

	corn := model.Corns{}
	corn.Enable = string(m.Enable)
	corn.Id = id
	err := service.MyService.Cron().UpdateStatus(strconv.Itoa(id), corn.Enable)
	if err != nil {
		message = err.Error()
		ctx.JSON(http.StatusInternalServerError, codegen.BadRequest{Message: message})
	}
	return ctx.JSON(http.StatusOK, codegen.OKRequest{Message: &message})
}
