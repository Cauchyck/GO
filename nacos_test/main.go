package main

import (
	"encoding/json"
	"fmt"
	"hello_go/nacos_test/config"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

func main() {
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr: "127.0.0.1",
			Port:   8848,
		},
	}

	clientConfig := constant.ClientConfig{
		NamespaceId:         "b6140e28-dea8-47a8-8bd2-e1f4a59173b5", //we can create multiple clients with different namespaceId to support multiple namespace.When namespace is public, fill in the blank string here.
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "tmp/nacos/log",
		CacheDir:            "tmp/nacos/cache",
		LogLevel:            "debug",
	}

	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"clientConfig":  clientConfig,
		"serverConfigs": serverConfigs,
	})
	if err != nil {
		panic(err)
	}

	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: "user-web.json",
		Group:  "dev",
	})

	if err != nil {
		panic(err)
	}
	// fmt.Println(content)
	serverConfig := config.ServerConfig{}
	json.Unmarshal([]byte(content), &serverConfig)
	fmt.Println(serverConfig)
	// err = configClient.ListenConfig(vo.ConfigParam{
	// 	DataId: "user-web.yaml",
	// 	Group: "dev",
	// 	OnChange: func(namespace, group, dataId, data string) {
	// 		fmt.Println("config file is changed")
	// 		fmt.Println("group:" + group + ", dataId:" + dataId + ", data:" + data)
	// 	},
	// })
	// time.Sleep(3000*time.Second)
}
