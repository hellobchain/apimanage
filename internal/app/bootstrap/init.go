package bootstrap

import (
	"ApiManager/internal/app/Validators"
	"ApiManager/internal/app/global"
	"database/sql"
	"log"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
)

var DbCon *sql.DB

func init() {
	initConfig()
	initValidation()
	initDatabase()
}

// 加载数据库连接
func initDatabase() {
	var err error
	host := global.DbConfig["host"]
	username := global.DbConfig["username"]
	password := global.DbConfig["password"]
	database := global.DbConfig["database"]
	port := global.DbConfig["port"]
	dns := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + database + "?parseTime=true"
	DbCon, err = sql.Open("mysql", dns)
	if err != nil {
		log.Fatal("[MYSQL ERROR] dns" + dns + "err:" + err.Error())
	}

	err = DbCon.Ping()
	if err != nil {
		log.Fatal("[MYSQL ERROR] ", err.Error())
	}
	DbCon.SetMaxIdleConns(10)
	DbCon.SetMaxOpenConns(50)
}

// 加载自定义表单验证器
func initValidation() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		_ = v.RegisterValidation("param_name", Validators.ParamName)
		_ = v.RegisterValidation("param_type", Validators.ParamType)
	}
}

// 读配置文件
func initConfig() error {
	return global.ReadConfig()
}
