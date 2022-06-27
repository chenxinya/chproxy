package internal

import (
	"fmt"
	"github.com/contentsquare/chproxy/config"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"gopkg.in/yaml.v3"
	"testing"
	"time"
)

func TestNacosGetConfig(t *testing.T) {
	var nacosConfig = NacosConfig{
		ServerPort:        8848,
		ServerIP:          "192.168.167.88",
		ClientNamespaceId: "bd1424e0-5b99-4170-85fb-24db3a7dc8f5",
	}
	cc := nacosConfig.InitNacosConf()
	//get config
	content, err := cc.GetConfig(vo.ConfigParam{
		DataId: "dev_chproxy",
		Group:  "chproxy_GROUP",
	})
	if err != nil {
		fmt.Printf("或者值出现问题%s\n", err)
	}
	fmt.Println("GetConfig,config :" + content)
	cfg := &config.Config{}
	fmt.Println("start yaml convert ----------")
	if err := yaml.Unmarshal([]byte(content), cfg); err != nil {
		fmt.Printf("ymal convert is error :%s \n", cfg)
	}
	fmt.Printf("config value is :%s\n", cfg)
}

func TestNacosListenConfig(t *testing.T) {
	var nacosConfig = NacosConfig{
		ServerPort:        8848,
		ServerIP:          "192.168.167.88",
		ClientNamespaceId: "bd1424e0-5b99-4170-85fb-24db3a7dc8f5",
	}
	cc := nacosConfig.InitNacosConf()

	//Listen config change,key=dataId+group+namespaceId.
	err := cc.ListenConfig(vo.ConfigParam{
		DataId: "dev_chproxy",
		Group:  "chproxy_GROUP",
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println("config changed group:" + group + ", dataId:" + dataId + ", content:" + data)
		},
	})
	if err != nil {
		fmt.Printf("listen config have error：%s\n", err)
	}

	time.Sleep(2 * time.Hour)
}
