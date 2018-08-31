package main

import (
	"fmt"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"

	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

const (
	channelID      = "mychannel"
	orgName        = "chenman" //"Org1"
	orgAdmin       = "Admin"
	ordererOrgName = "OrdererOrg"
)
const (
	sdkConfigFile      = "fabric_kafka/project/configfile/configfile.yaml"
	sdkValidClientUser = "Admin"
	sdkValidClientOrg1 = "chenman"
)

var (
	ccID = "hyctest"
)

// 首先 根据配置文件 获取sdk
// 要做什么操作，先实例化一个 对应的 客户端 来进行操作
// 实例化客户端 需要相应的 上下文
// 所以第二步骤就是 实例化 你需要操作 的客户端 所需要的上下文喽
// 用客户端 进行 相关操作就可以了

// 下面就按照我 搭建 环境测试的例子步骤 来 使用sdk
// 创建chanell channel_sdk_example
// peer 加入此 频道
// 安装 chaincode
// 实例化 chaincode
// 查询 chaincode

func createChannel(channelName string) bool {

	// 第一步 根据配置文件 获取sdk，1.获取配置，2.根据配置生成sdk
	sdkconfig := config.FromFile(sdkConfigFile)
	sdk, err := fabsdk.New(sdkconfig)

	// 检查 sdk 初始化成功没有
	if err != nil {
		fmt.Printf("get sdk have error =" + err.Error())
	}

	// 构建 创建频道需要的上下文,fabsdk 就是用来创建上下文的
	// 回一下 创建频道 必要条件
	// 1. 根据peer.yaml启动peer 容器，peer 跟orderer 连接起来
	// 2. 进入容器，根据channel.tx
	// 3. 创建频道
	// 看fabric 的上下文api有几个分别是，withOrg withUser WithIdentfiler
	// 相关的组织 ，相关的用户 相关的ident
	// 相关的组织，我这个节点先暂时定义 chenman ,
	createChannelContext := sdk.Context(fabsdk.WithOrg("chenman"), fabsdk.WithOrg("hyc0.example.com:7050"))

	// 创建channel 需要resm 客户端
	resmgmtClient, err := resmgmt.New(createChannelContext)
	if err != nil {
		fmt.Printf("resmgmtClient get error =" + err.Error())
	}

	// 有了客户端就创建channel
	_, err = resmgmtClient.SaveChannel(resmgmt.SaveChannelRequest{ChannelID: "mychannel", ChannelConfigPath: "fabric_kafka/channel-artifacts/mychannel.tx"})
	is := true
	if err != nil {
		is = false
	}

	return is
}

// func callChaincode(configOpt core.ConfigProvider, sdkOpts ...fabsdk.Option) {
// 	// 读取sdk 配置文件
// 	sdkConfig := config.FromFile(sdkConfigFile)

// 	// 根据配置文件生成相应的sdk
// 	sdk, err := fabsdk.New(sdkConfig)
// 	if err != nil {
// 		fmt.Printf("初始化sdk出错了" + err.Error())
// 	}
// 	defer sdk.Close()

// 	// 获取channel client
// 	client, err := mspclient.New(sdk.Context(), mspclient.WithOrg(orgName))
// 	if err != nil {
// 		fmt.Printf("client失败" + err.Error())
// 	}

// }

func main() {
	createChannel("mychannel")
}
