/*
 * @Author: a624669980@163.com a624669980@163.com
 * @Date: 2022-11-28 14:55:35
 * @LastEditors: a624669980@163.com a624669980@163.com
 * @LastEditTime: 2022-12-03 22:29:56
 * @FilePath: /gocron/repository/db.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package repository

import (
	"fmt"
	"os"
	"time"

	"github.com/IceWhaleTech/CasaOS-Common/utils/logger"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var gdb *gorm.DB

func GetDb() *gorm.DB {

	url := os.Getenv("MYSQLURL")
	if url == "" {
		url = "192.168.2.3:45678"
		//url = "127.0.0.1:3306"
	}
	user := os.Getenv("MYSQLUSER")
	if user == "" {
		user = "root"
		//user = "casa"
	}
	password := os.Getenv("MYSQLPASSWORD")
	if password == "" {
		password = "123456"
	}
	// Refer https://github.com/go-sql-driver/mysql#dsn-data-source-name
	dsn := fmt.Sprintf("%v:%v@tcp(%v)/gocron?charset=utf8mb4&parseTime=True&loc=Local", user, password, url)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
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

	//	err = db.AutoMigrate(&model2.AppNotify{}, &model2.AppListDBModel{}, model2.SharesDBModel{}, model2.ConnectionsDBModel{})
	if err != nil {
		logger.Error("check or create db error", zap.Any("error", err))
	}
	return db
}
