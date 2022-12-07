/*
 * @Author: a624669980@163.com a624669980@163.com
 * @Date: 2022-11-14 17:18:18
 * @LastEditors: a624669980@163.com a624669980@163.com
 * @LastEditTime: 2022-12-03 21:33:01
 * @FilePath: /gocron/route/v1/log.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package v1

import (
	"gocron/codegen"
	"gocron/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetCronLog(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	//id := vars["id"]
	// logs, err := service.Service.Cron().GetJobLog(id)
	// if err != nil {
	// 	result := model.Result{Code: 500, Msg: "error", Data: err}
	// 	json.NewEncoder(w).Encode(result)
	// 	return
	// }
	// result := model.Result{Code: 200, Msg: "success", Data: logs}
	w.Header().Set("Content-Type", "application/json")
	//json.NewEncoder(w).Encode(result)
}
func (g *GoCron) GetLogById(ctx echo.Context, id int64) error {
	list, err := service.MyService.Cron().GetJobLog(strconv.FormatInt(id, 10))
	if err != nil {
		return ctx.JSON(http.StatusOK, codegen.BadRequest{Message: err.Error()})
	}

	logs := []codegen.Log{}
	for _, v := range list {
		time := strconv.FormatInt(v.Time, 10)
		id := v.Id
		logs = append(logs, codegen.Log{
			Id:     &id,
			Name:   v.Name,
			CronId: strconv.Itoa(v.CronId),
			Log:    v.Log,
			Time:   &time,
		})
	}
	return ctx.JSON(http.StatusOK, codegen.GetLogByIdOK{Data: &logs})
}
