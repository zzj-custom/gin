package nacos

import (
	"fmt"
	"github.com/gookit/goutil/dump"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"sync"
)

func (c *Config) clientConfig() *constant.ClientConfig {
	clientConfig := *constant.NewClientConfig(
		constant.WithNamespaceId(c.NamespaceId), //当namespace是public时，此处填空字符串。
		constant.WithTimeoutMs(c.TimeoutMS),
		constant.WithNotLoadCacheAtStart(true),
		constant.WithLogLevel("warn"),
		constant.WithUsername(c.UserName),
		constant.WithPassword(c.Password),
	)
	return &clientConfig
}

func (c *Config) serverConfig() []constant.ServerConfig {
	serverConfigs := []constant.ServerConfig{
		*constant.NewServerConfig(
			c.Host,
			c.Port,
			constant.WithScheme(c.Scheme),
			constant.WithContextPath(fmt.Sprintf("/%s", c.ContextPath)),
		),
	}
	return serverConfigs
}

var (
	configOnce sync.Once
	config     config_client.IConfigClient
	handler    Handler
)

func (c *Config) dClient() config_client.IConfigClient {
	configOnce.Do(func() {
		config = c.dynamicConfigurationClient()
	})
	return config
}

// DynamicConfigurationClient 动态配置客户端
func (c *Config) dynamicConfigurationClient() config_client.IConfigClient {
	configClient, err := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  c.clientConfig(),
			ServerConfigs: c.serverConfig(),
		},
	)
	if err != nil {
		fmt.Printf("连接nacos动态配置客户端失败，错误：%v", err)
		return nil
	}
	return configClient
}

// GetConfig 读取配置列表
func (c *Config) GetConfig() {
	// 读取配置
	configClient := c.dClient()
	config, err := configClient.GetConfig(
		vo.ConfigParam{
			DataId: "ch-gateway-router.json",
			Group:  "DEFAULT_GROUP",
		},
	)
	if err != nil {
		fmt.Printf("读取信息失败, err: %v", err)
	}
	dump.P(config)
}

// GetAllConfigFile 读取所有的配置文件
func (c *Config) GetAllConfigFile() {
	configClient := c.dClient()
	configPage, err := configClient.SearchConfig(vo.SearchConfigParam{
		Search:   "blur",
		DataId:   "",
		Group:    "",
		PageNo:   1,
		PageSize: int(c.PageSize),
	})
	if err != nil {
		fmt.Printf("读取配置文件失败，err: %v", err)
		return
	}
	for _, item := range configPage.PageItems {
		handler.Handle(item.DataId, item.Content)
		c.listenConfig(item.DataId, item.Group)
	}
}

func (c *Config) listenConfig(dataId string, group string) {
	configClient := c.dClient()
	err := configClient.ListenConfig(vo.ConfigParam{
		DataId: dataId,
		Group:  group,
		OnChange: func(namespace, group, dataId, data string) {
			fmt.Printf("data:%s\n", data)
			// 更新数据
			handler.Handle(dataId, data)
		},
	})
	if err != nil {
		fmt.Printf("监听出现问题，err:%v", err)
	}
	return
}
