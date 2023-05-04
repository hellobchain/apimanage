package global

import (
	"ApiManager/internal/app/libs"
	"os"
)

var (
	DbConfig                = make(map[string]string)
	SiteConfig              = make(map[string]string)
	SessionDrive            string
	SessionOption           = make(map[string]interface{})
	SessionDriveRedisConfig = make(map[string]interface{})
	GinRunMode              string
	GinWriteLog             bool
)

func ReadConfig() error {
	rootPath, err := os.Getwd()
	if err != nil {
		return err
	}
	config, err := libs.ReadIniFile(rootPath + "/deployments/config/config.ini")
	if err != nil {
		return err
	}

	// 站点配置
	SiteConfig["http_port"], err = config.GetConfig("site.http_port")
	if err != nil {
		return err
	}

	// 数据库配置
	DbConfig["host"], err = config.GetConfig("mysql.host")
	if err != nil {
		return err
	}
	DbConfig["username"], err = config.GetConfig("mysql.username")
	if err != nil {
		return err
	}
	DbConfig["password"], err = config.GetConfig("mysql.password")
	if err != nil {
		return err
	}
	DbConfig["database"], err = config.GetConfig("mysql.database")
	if err != nil {
		return err
	}
	DbConfig["port"], err = config.GetConfig("mysql.port")
	if err != nil {
		return err
	}

	// session驱动类型
	SessionDrive, err = config.GetConfig("session.driver")
	if err != nil {
		return err
	}

	// session 配置项
	SessionOption["max_age"], err = config.GetConfigToInt("session.max_age")
	if err != nil {
		return err
	}
	SessionOption["path"], err = config.GetConfig("session.path")
	if err != nil {
		return err
	}
	SessionOption["http_only"], err = config.GetConfigToBool("session.http_only")
	if err != nil {
		return err
	}
	SessionOption["domain"], err = config.GetConfig("session.domain")
	if err != nil {
		return err
	}
	SessionOption["secure"], err = config.GetConfigToBool("session.secure")
	if err != nil {
		return err
	}

	// 若session驱动为redis则读redis配置项
	if SessionDrive == "redis" {
		SessionDriveRedisConfig["size"], err = config.GetConfigToInt("session_driver_redis.max_idel_con")
		if err != nil {
			return err
		}
		SessionDriveRedisConfig["network"], err = config.GetConfig("session_driver_redis.network")
		if err != nil {
			return err
		}
		SessionDriveRedisConfig["address"], err = config.GetConfig("session_driver_redis.address")
		if err != nil {
			return err
		}
		SessionDriveRedisConfig["password"], err = config.GetConfig("session_driver_redis.password")
		if err != nil {
			return err
		}
	}

	// 运行模式
	GinRunMode, err = config.GetConfig("site.gin_run_mode")
	if err != nil {
		return err
	}

	// 是否记录运行日志
	GinWriteLog, err = config.GetConfigToBool("site.gin_write_log")
	if err != nil {
		return err
	}
	return nil
}
