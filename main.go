package main

import (
	"dns/config"
	"dns/controller"
	"dns/httpGin"
	"dns/model"
	"flag"
	"fmt"
	"os"
	"strings"
)

var etcdUrl = flag.String("etcd.url", "http://127.0.0.1:2379", "ETCD数据库的连接地址")
var etcdPath = flag.String("etcd.path", "/coredns", "ETCD数据库存储路径的前缀")
var port = flag.Int("port", 9101, "Web服务监听的端口")

func main() {
	flag.Parse()

	// 解析链接
	config.Etcd_url = append(config.Etcd_url, strings.Split(*etcdUrl, ",")...)

	if len(config.Etcd_url) == 0 {
		fmt.Println("地址不存在")
		os.Exit(1)
	}
	config.DBKeyPath = *etcdPath
	if config.DBKeyPath == "" {
		config.DBKeyPath = "/coredns"
	} else {
		// 判断根是否存在
		if !strings.HasPrefix(config.DBKeyPath, "/") {
			config.DBKeyPath = "/" + config.DBKeyPath
		}
	}
	//初始化检测etcd链接情况
	model.OninitCheck()

	// 链接后监听数据库变化
	controller.Oninit()

	// 启动Http服务
	httpGin.StartHttp(fmt.Sprint(*port))
}
