package internal

import (
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

type NacosConfig struct {
	ServerIP          string `yaml:"server_ip"`
	ServerPort        uint64 `yaml:"server_port"`
	ClientNamespaceId string `yaml:"client_namespace_id"`
	DataId            string `yaml:"data_id"`
	Group             string `yaml:"group"`
	ConfigClient      config_client.IConfigClient
}

func (n *NacosConfig) InitNacosConf() config_client.IConfigClient {
	sc := []constant.ServerConfig{
		{
			IpAddr: n.ServerIP,
			Port:   n.ServerPort,
		},
	}
	cc := constant.ClientConfig{
		NamespaceId:         n.ClientNamespaceId, //namespace id
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "nacos/log",
		CacheDir:            "nacos/cache",
		LogLevel:            "debug",
	}

	// a more graceful way to create config client
	client, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	n.ConfigClient = client
	if err != nil {
		panic(err)
	}
	return client
}

func (n *NacosConfig) GetNacosConfig() config_client.IConfigClient {
	sc := []constant.ServerConfig{
		{
			IpAddr: n.ServerIP,
			Port:   n.ServerPort,
		},
	}
	cc := constant.ClientConfig{
		NamespaceId:         n.ClientNamespaceId, //namespace id
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "nacos/log",
		CacheDir:            "nacos/cache",
		LogLevel:            "debug",
	}

	// a more graceful way to create config client
	client, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)

	if err != nil {
		panic(err)
	}
	return client
}
