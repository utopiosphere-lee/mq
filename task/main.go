package main

import (
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"task/conf"
	"task/core"
	"task/services"
)

func main()  {
	conf.Init()
	//	etcd注册件
	etcdReg := etcd.NewRegistry(
		registry.Addrs("127.0.0.1:2379"),
	)
	//	得到一个微服务实例
	microService := micro.NewService(
		micro.Name("rpcTaskService"),	//	微服务名称
		micro.Address("127.0.0.1:8083"),	//	当前微服务注册端口
		micro.Registry(etcdReg),
	)

	//	结构命令行参数 初始化
	microService.Init()
	//	服务注册
	_ = services.RegisterTaskServiceHandler(microService.Server(), new(core.TaskService))
	//	启动微服务
	_ = microService.Run()
}
