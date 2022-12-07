//go:generate bash -c "mkdir -p codegen && go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.12.2 -generate types,server,spec -package codegen api/gocron/openapi.yaml > codegen/gocron_api.go"
/*
 * @Author: a624669980@163.com a624669980@163.com
 * @Date: 2022-10-21 11:47:53
 * @LastEditors: a624669980@163.com a624669980@163.com
 * @LastEditTime: 2022-11-29 15:22:15
 * @FilePath: /gocron/main.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package main

import (
	"gocron/repository"
	"gocron/route"
	"gocron/service"
	"log"
	"net/http"
	"time"

	"github.com/IceWhaleTech/CasaOS/pkg/utils/loger"
	"github.com/robfig/cron/v3"
)

func init() {
	loger.LogInit()
	repository.MyRepository = repository.NewRepository()
	service.MyService = service.NewService(repository.MyRepository, cron.New(cron.WithSeconds()))
}
func main() {
	r := route.InitV1Router()
	srv := &http.Server{
		Handler: r,
		Addr:    "0.0.0.0:8008",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
	// c := cron.New(cron.WithSeconds())
	// c.AddJob()

}
