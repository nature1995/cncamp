package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"

	"hw3/internal/logger"
)

var C = new(Config)

func GetConfig(filePath string) error {
	vip := viper.New()
	if len(filePath) == 0 {
		// 获取项目的执行路径
		path, err := os.Getwd()
		if err != nil {
			return err
		}
		vip.AddConfigPath(path + "/config")
		// 设置文件的类型
	} else {
		vip.AddConfigPath(filePath)
	}
	vip.SetConfigName("config")
	vip.SetConfigType("yaml")

	// 进行配置读取
	if err := vip.ReadInConfig(); err != nil {
		return err
	}

	err := vip.Unmarshal(&C)
	if err != nil {
		return err
	}
	fmt.Printf("final config: %+v\n", *C)
	return nil
}

type Config struct {
	Env  string
	HTTP string
	Log  logger.Config
}
