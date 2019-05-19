package main

import (
	"github.com/hxlb/pkg/config"
	"github.com/hxlb/pkg/log"
	"github.com/hxlb/user/handler"
	"github.com/hxlb/user/proto/user"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"time"
)

func main() {
	//加载配置项
	conf := config.New()
	err := conf.Load("./config/config.json")
	if err != nil {
		log.Logger("user").Fatal(err.Error())
		return
	}
	//
	reg := consul.NewRegistry(
		registry.Addrs(conf.GetString("consul")),
	)
	// New Service
	service := micro.NewService(
		micro.Name(conf.GetString("server_name")),
		micro.Version(conf.GetString("server_version")),
		micro.Registry(reg),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*15),
	)
	// 定义Service动作操作
	service.Init(
		micro.Action(func(c *cli.Context) {
			log.Logger("user").Debug("user-srv is start ...")
			// Register Handler
			com_hxlb_srv_user.RegisterUserServiceHandler(service.Server(), new(handler.User))
		}),
		micro.AfterStop(func() error {
			log.Logger("user").Debug("user-srv is stop ...")
			return nil
		}),
		micro.AfterStart(func() error {
			return nil
		}),
	)

	// Run service
	if err := service.Run(); err != nil {
		log.Logger("user").Fatal(err.Error())

	}
}
