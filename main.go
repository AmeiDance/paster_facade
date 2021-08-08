package main

import (
	"fmt"
	"io/ioutil"

	"github.com/ameidance/paster_facade/client"
	"github.com/ameidance/paster_facade/client/consul"
	"github.com/ameidance/paster_facade/constant"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v3"
)

var router *gin.Engine

func main() {
	ginConf, err := getGinConfig()
	if err != nil {
		klog.Errorf("[main] get gin config failed. err:%v", err)
		panic(err)
	}

	client.InitRedis()
	consul.InitConsul(ginConf.Port)
	client.InitRpc()

	if err = router.Run(fmt.Sprintf("%s:%d", ginConf.Address, ginConf.Port)); err != nil {
		klog.Errorf("[main] server stopped with error. err:%v", err)
		panic(err)
	}
}

type _GinConf struct {
	Address string `yaml:"address"`
	Port    int    `yaml:"port"`
}

func getGinConfig() (*_GinConf, error) {
	conf := new(_GinConf)
	file, err := ioutil.ReadFile(constant.GIN_CONF_PATH)
	if err != nil {
		klog.Errorf("[getGinConfig] open file failed. err:%v", err)
		return nil, err
	}
	if err = yaml.Unmarshal(file, conf); err != nil {
		klog.Errorf("[getGinConfig] unmarshal file failed. err:%v", err)
		return nil, err
	}
	return conf, nil
}
