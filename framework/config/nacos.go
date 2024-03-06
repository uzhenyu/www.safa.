package config

import (
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/v2/clients"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/v2/clients/naming_client"
	"github.com/nacos-group/nacos-sdk-go/v2/common/constant"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
)

const (
	ip   = "127.0.0.1"
	port = 8848
)

var client config_client.IConfigClient
var clie naming_client.INamingClient

func GetClient() error {
	var err error
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(ip, port, constant.WithContextPath("/nacos")),
	}
	cc := *constant.NewClientConfig(
		constant.WithNamespaceId(""),
		constant.WithTimeoutMs(5000),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogDir("/tmp/nacos/log"),
		constant.WithCacheDir("/tmp/nacos/cache"),
		constant.WithLogLevel("debug"),
	)
	client, err = clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
	return err
}

func GetConfig(group, dataID string) (string, error) {
	if client == nil {
		// 初始化 client
		err := GetClient()
		if err != nil {
			return "", err
		}
	}
	content, err := client.GetConfig(vo.ConfigParam{
		DataId: dataID,
		Group:  group,
	})
	if err != nil {
		return "", err
	}
	return content, nil
}

func ListenConfig(group, dataID string) error {
	return client.ListenConfig(vo.ConfigParam{
		DataId: "dataId",
		Group:  "group",
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Println("group:" + group + ", dataId:" + dataId + ", data:" + data)
		},
	})
}

func Service(host string, ports int) error {
	// 创建clientConfig
	clientConfig := constant.ClientConfig{
		NamespaceId:         "", // 如果需要支持多namespace，我们可以场景多个client,它们有不同的NamespaceId。当namespace是public时，此处填空字符串。
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "/tmp/nacos/log",
		CacheDir:            "/tmp/nacos/cache",
		LogLevel:            "debug",
	}

	serverConfigs := []constant.ServerConfig{
		{
			IpAddr: ip,
			Port:   port,
		},
	}

	cli, _ := clients.CreateNamingClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})

	_, err := cli.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          host,
		Port:        uint64(ports),
		ServiceName: "he ke hue da sha bi",
		Weight:      10,
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
		Metadata:    map[string]string{"idc": "shanghai"},
		ClusterName: "DEFAULT",       // 默认值DEFAULT
		GroupName:   "DEFAULT_GROUP", // 默认值DEFAULT_GROUP
	})
	if err != nil {
		return err
	}
	return nil
}

func NacosService() (string, int, error) {
	data, err := clie.GetService(vo.GetServiceParam{
		ServiceName: "DEFAULT",
		GroupName:   "DEFAULT_GROUP",
		Clusters:    []string{"cluster-a"},
	})
	if err != nil {
		return "", 0, err
	}
	var host string
	var ports int

	for _, val := range data.Hosts {
		host = val.Ip
		ports = int(val.Port)
	}
	return host, ports, nil

}
