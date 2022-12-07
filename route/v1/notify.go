/*
 * @Author: a624669980@163.com a624669980@163.com
 * @Date: 2022-11-14 17:23:40
 * @LastEditors: a624669980@163.com a624669980@163.com
 * @LastEditTime: 2022-12-03 21:59:09
 * @FilePath: /gocron/route/v1/notify.go
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

	"github.com/labstack/echo/v4"
)

func PostNotifyAdd(w http.ResponseWriter, r *http.Request) {
	m := model.Notify{}
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
	// err = service.Service.Notify().AddNotify(m)
	// if err != nil {
	// 	result := model.Result{Code: 500, Msg: "error", Data: err}
	// 	json.NewEncoder(w).Encode(result)
	// 	return
	// }
	// w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(json.RawMessage(string(`{"code": 200, "msg": "添加成功"}`)))
}

func PutNotifyUpdate(w http.ResponseWriter, r *http.Request) {
	m := model.Notify{}
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
	// err = service.Service.Notify().UpdateNotify(m)
	// if err != nil {
	// 	result := model.Result{Code: 500, Msg: "error", Data: err}
	// 	json.NewEncoder(w).Encode(result)
	// 	return
	// }
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(json.RawMessage(string(`{"code": 200, "msg": "修改成功"}`)))
}
func GetNotifyList(w http.ResponseWriter, r *http.Request) {
	// list := service.Service.Notify().GetNotifyList()
	// result := model.Result{Code: 200, Msg: "success", Data: list}
	w.Header().Set("Content-Type", "application/json")
	//json.NewEncoder(w).Encode(result)
}
func DeleteNotifyDelete(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// id := vars["id"]
	// err := service.Service.Notify().DeleteNotify(id)
	// if err != nil {
	// 	result := model.Result{Code: 500, Msg: "error", Data: err}
	// 	json.NewEncoder(w).Encode(result)
	// 	return
	// }
	w.Header().Set("Content-Type", "application/json")
	//json.NewEncoder(w).Encode(json.RawMessage(string(`{"code": 200, "msg": "删除成功"}`)))
}

func (g *GoCron) GetNotifications(ctx echo.Context) error {
	list := service.MyService.Notify().GetNotifyList()
	notifies := []codegen.Notification{}
	for _, v := range list {
		t := codegen.NotificationType(v.Type)
		notifies = append(notifies, codegen.Notification{
			Title:   &v.Title,
			Content: &v.Content,
			Url:     &v.Url,
			Name:    &v.Name,
			Id:      &v.Id,
			Type:    &t,
		})
	}
	return ctx.JSON(http.StatusOK, codegen.GetNotificationsOK{Data: &notifies})
}

// Create a notification
// (POST /notify)
func (g *GoCron) CreateNotification(ctx echo.Context) error {
	m := codegen.Notification{}
	message := "OK"
	if err := ctx.Bind(&m); err != nil {
		message = err.Error()
		return ctx.JSON(http.StatusBadRequest, codegen.BadRequest{Message: message})
	}
	notify := model.Notify{
		Title:   *m.Title,
		Content: *m.Content,
		Type:    string(*m.Type),
		Url:     *m.Url,
		Name:    *m.Name,
	}
	err := service.MyService.Notify().AddNotify(notify)
	if err != nil {
		message = err.Error()
		return ctx.JSON(http.StatusBadRequest, codegen.BadRequest{Message: message})
	}
	return ctx.JSON(http.StatusOK, codegen.OK{Message: &message})
}

// Delete a notification
// (DELETE /notify/{id})
func (g *GoCron) DeleteNotification(ctx echo.Context, id int) error {
	message := "OK"
	err := service.MyService.Notify().DeleteNotify(strconv.Itoa(id))
	if err != nil {
		message = err.Error()
		return ctx.JSON(http.StatusBadRequest, codegen.BadRequest{Message: message})
	}
	return ctx.JSON(http.StatusOK, codegen.OK{Message: &message})
}
