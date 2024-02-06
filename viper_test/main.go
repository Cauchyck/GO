package main

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/fsnotify/fsnotify"
)

type ServerConfig struct{
	ServiceName string `mapstructure:"name"`
	Port int `mapstructure:"port"`
}

type MysqlConfig struct {
	Host string `mapstructure:"host"`
	Port int `mapstructure:"port"`
}
func main(){
	v := viper.New()
	v.SetConfigFile("viper_test/config.yaml")
	if err:=v.ReadInConfig(); err != nil{
		panic(err)
	}
	serverConfig := ServerConfig{}
	if err := v.Unmarshal(&serverConfig); err != nil{
		panic(err)
	}
	fmt.Println(serverConfig)

	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event){
		fmt.Println("config file changed:", e.Name)
		_ = v.ReadInConfig()
		_ = v.Unmarshal(&serverConfig)
		fmt.Println(serverConfig)
	})



}