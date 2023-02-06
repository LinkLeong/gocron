/*
 * @Author: a624669980@163.com a624669980@163.com
 * @Date: 2022-10-21 14:56:53
 * @LastEditors: a624669980@163.com a624669980@163.com
 * @LastEditTime: 2022-12-03 22:01:38
 * @FilePath: /gocron/pkg/sqlite/db.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package sqlite

import (
	"gocron/model"
	"time"

	"github.com/IceWhaleTech/CasaOS-Common/utils/file"
	"github.com/IceWhaleTech/CasaOS-Common/utils/logger"
	"github.com/glebarez/sqlite"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var gdb *gorm.DB

func GetDb(dbPath string) *gorm.DB {
	if gdb != nil {
		return gdb
	}
	file.IsNotExistMkDir(dbPath)
	db, err := gorm.Open(sqlite.Open(dbPath+"/cron.db"), &gorm.Config{})
	c, _ := db.DB()
	c.SetMaxIdleConns(10)
	c.SetMaxOpenConns(100)
	c.SetConnMaxIdleTime(time.Second * 1000)
	if err != nil {
		logger.Error("sqlite connect error", zap.Any("db connect error", err))
		panic("sqlite connect error")
		return nil
	}
	gdb = db
	err = db.AutoMigrate(&model.Corns{}, &model.CornsLog{}, &model.Notify{})
	if err != nil {
		logger.Error("check or create db error", zap.Any("error", err))
	}
	return db
}
