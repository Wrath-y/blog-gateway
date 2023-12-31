package main

import (
	"flag"
	"gateway/infrastructure/config"
	"gateway/infrastructure/util/consul"
	"gateway/infrastructure/util/db"
	"gateway/infrastructure/util/goredis"
	"gateway/infrastructure/util/logging"
	"gateway/launch/gin"
	"github.com/spf13/viper"
	"log"
	"syscall"
)

// setupConfigYaml 就绪配置文件
// 环境变量配置 NACOS_SKIP="Y", 可跳过下载配置
// 环境变量:
// NACOS_USE=false
// NACOS_NAMESPACE=""
// NACOS_SERVER=""
// NACOS_USERNAME=""
// NACOS_PASSWORD=""
func setupConfigYaml() {
	viper.AutomaticEnv()
	if envUse := viper.GetBool("NACOS_USE"); !envUse {
		log.Println("跳过从nacos下载配置文件")
		return
	}

	config.SetupNacosClient()
	config.DownloadNacosConfig()

	// 未使用k8s部署时监听nacos（已经被使用的变量不会体现出变化）
	//config.ListenNacos()

	// 使用k8s部署时监听nacos（已经被使用的变量不会体现出变化）
	config.ListenNacos(func(cnf string) {
		log.Println("当前进程将被停止")
		syscall.Kill(syscall.Getpid(), syscall.SIGTERM)
	})
}

func setup() {
	config.Setup()
	logging.Setup()
	goredis.Setup()
	db.Setup()
	consul.Setup()
}

func main() {
	flag.Parse()
	setupConfigYaml()

	setup()

	gin.RunGin()

	logging.New().Info("Server exited", "", nil)
}
