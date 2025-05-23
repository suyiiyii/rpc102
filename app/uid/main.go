// Code generated by kitex v0.9.1, Customize by suyiiyii at https://github.com/suyiiyii/cwgo-template
package main

import (
	"net"
	"os"
	"rpc102/app/uid/conf"

	"rpc102/app/uid/biz/dal"
	"rpc102/app/uid/kitex_gen/uid/uidservice"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/doutokk/doutok/common/serversuite"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
)

func main() {
	// use `go run cmd/gorm/main.go` to migrate the database
	dal.Init()
	// use `go run cmd/gorm_gen/main.go` to generate the code
	//query.SetDefault(mysql.DB)
	opts := kitexInit()

	svr := uidservice.NewServer(new(UidServiceImpl), opts...)

	err := svr.Run()
	if err != nil {
		klog.Error(err.Error())
	}
}

func kitexInit() (opts []server.Option) {
	// address
	addr, err := net.ResolveTCPAddr("tcp", conf.GetConf().Kitex.Address)
	if err != nil {
		panic(err)
	}
	opts = append(opts, server.WithServiceAddr(addr))

	// service info
	opts = append(opts, server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{
		ServiceName: conf.GetConf().Kitex.Service,
	}))

	// registry
	opts = append(opts,
		server.WithSuite(serversuite.CommonServerSuite{
			CurrentServiceName: conf.GetConf().Kitex.Service,
			RegistryAddr:       conf.GetConf().Registry.RegistryAddress[0],
		}))

	// klog
	logger := kitexlogrus.NewLogger()
	klog.SetLogger(logger)
	klog.SetLevel(conf.LogLevel())
	klog.SetOutput(os.Stdout)
	return
}
