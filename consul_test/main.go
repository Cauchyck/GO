package main

import (
	"fmt"

	"github.com/hashicorp/consul/api"
)

func Register(address string, port int, name string, tags []string, id string) error {
	cfg := api.DefaultConfig()

	cfg.Address = "127.0.0.1:8500"

	client, err := api.NewClient(cfg)

	if err  != nil {
		panic(err)
	}

	check := &api.AgentServiceCheck{
		HTTP: "http://127.0.0.1:8889/health",
		Timeout: "5s",
		Interval: "5s",
		DeregisterCriticalServiceAfter: "10s",
	}

	registration := new(api.AgentServiceRegistration)
	registration.Name = name
	registration.ID = id
	registration.Port = port
	registration.Tags = tags
	registration.Address = address
	registration.Check = check

	err = client.Agent().ServiceRegister(registration)
	if err != nil{
		panic(err)
	}

	return nil

}

func AllServices(){
	cfg := api.DefaultConfig()
	cfg.Address = "127.0.0.1:8500"

	client, err := api.NewClient(cfg)
	if err != nil {
		panic(err)
	}

	data, err := client.Agent().Services()
	if err != nil {
		panic(err)
	}
	for key, _ := range data {
		fmt.Println(key)
	}
}

func FilterSerivice(){
	cfg := api.DefaultConfig()
	cfg.Address = "127.0.0.1:8500"

	client, err := api.NewClient(cfg)

	if err != nil {
		panic(err)
	}

	data, err := client.Agent().ServicesWithFilter(`Service == "user-web"`)
	if err != nil {
		panic(err)
	}
	for key, _ := range data {
		fmt.Println(key)
	}
}




func main(){
	// err := Register("127.0.0.1", 8889, "user-web", []string{"mxshop", "bobby"}, "user-web")
	// if err != nil {
	// 	panic(err)
	// }
	// AllServices()
	FilterSerivice()
}