package config

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Load config
func Load(filepath string, val any, hook func(v *viper.Viper)) {
	v, err := loadConfig(filepath, hook)
	if err != nil {
		log.Fatalf("load config err:%v", err)
	}

	if err = v.Unmarshal(&val); err != nil {
		log.Fatalf("unmarshal config err:%v", err)
	}

	v.WatchConfig()
	// 注册每次配置文件发生变更后都会调用的回调函数
	v.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("Config file changed: %s", e.Name)
		// 每次配置文件发生变化，需要重新将其反序列化到结构体中
		if err := v.Unmarshal(&val); err != nil {
			panic(fmt.Errorf("unmarshal config error: %s", err.Error()))
		}
	})
}

// loadConfig load config file from given path
func loadConfig(filepath string, hook func(v *viper.Viper)) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigFile(filepath) // 如果指定了配置文件，则解析指定的配置文件
	v.SetConfigType("yaml")   // 设置配置文件格式为YAML
	v.AutomaticEnv()          // 读取匹配的环境变量
	v.SetEnvPrefix("mall")    // 读取环境变量的前缀为 mall
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	if hook != nil {
		hook(v)
	}
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			//配置文件未找到错误；如果需要可以忽略
			return nil, errors.New("config file not found")
		}
		// 配置文件被找到，但产生了另外的错误
		return nil, err
	}
	log.Println("Using config file: ", v.ConfigFileUsed(), " settings: ", v.AllSettings())

	return v, nil
}
