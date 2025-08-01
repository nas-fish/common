package nacos

import (
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"github.com/nas-fish/common/choose"
	"github.com/nas-fish/common/env"
)

func NewNacosClient() (naming_client.INamingClient, error) {
	if env.IsProd() {
		return newProdNacosClient()
	} else {
		client, err := newProdNacosClient()
		if err != nil {
			return newBoeNacosClient()
		}
		return client, nil
	}
}

func newProdNacosClient() (naming_client.INamingClient, error) {
	return clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig: &constant.ClientConfig{
				NamespaceId: choose.If(env.IsProd(), "public", "boe"),
				Username:    "nacos",
				Password:    "wxl5211314",
			},
			ServerConfigs: []constant.ServerConfig{
				*constant.NewServerConfig("172.21.134.141", 8848),
			},
		},
	)
}

func newBoeNacosClient() (naming_client.INamingClient, error) {
	return clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig: &constant.ClientConfig{
				NamespaceId: "public",
				Username:    "nacos",
				Password:    "wxl5211314",
			},
			ServerConfigs: []constant.ServerConfig{
				*constant.NewServerConfig("172.21.134.141", 8848),
			},
		},
	)
}
