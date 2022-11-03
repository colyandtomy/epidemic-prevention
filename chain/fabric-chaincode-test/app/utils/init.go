package utils

import (
	"fmt"
	"log"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/errors/retry"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/msp"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"

	mspclient "github.com/hyperledger/fabric-sdk-go/pkg/client/msp"
)

func SetupSDK() (*fabsdk.FabricSDK, error) {

	sdk, err := fabsdk.New(config.FromFile(ConfigFile))
	if err != nil {
		return nil, fmt.Errorf("实例化Fabric SDK失败: %v", err)
	}

	fmt.Println("Fabric SDK初始化成功")
	return sdk, nil
}

func CreatChannelClient(sdk *fabsdk.FabricSDK) (*channel.Client, error) {

	clientChannelContext := sdk.ChannelContext(
		ChannelID,
		fabsdk.WithUser("User1"))
	channelClient, err := channel.New(clientChannelContext)
	if err != nil {
		return nil, fmt.Errorf("创建应用通道客户端失败: %v", err)
	}
	fmt.Println("通道客户端创建成功，可以利用此客户端调用链码进行查询或执行事务.")

	return channelClient, nil
}

func CreateResMgmtClient(sdk *fabsdk.FabricSDK, orgAdmin, orgName string) (*resmgmt.Client, error) {

	clientContext := sdk.Context(
		fabsdk.WithUser(orgAdmin),
		fabsdk.WithOrg(orgName))
	if clientContext == nil {
		return nil, fmt.Errorf(orgName + "创建资源管理客户端Context失败")
	}
	fmt.Println(orgName + "创建资源管理客户端成功...")

	// 返回资源管理客户端实例
	resMgmtClient, err := resmgmt.New(clientContext)
	if err != nil {
		return nil, fmt.Errorf(orgName+"创建通道管理客户端失败: %v", err)
	}
	fmt.Println(orgName + "创建通道管理客户端成功...")
	return resMgmtClient, nil
}

func CreateChannel(sdk *fabsdk.FabricSDK) error {

	resMgmtClient, err := CreateResMgmtClient(sdk, OrgAdmin, OrdererOrgName)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	if !Initialized {
		fmt.Println("保存带有交易ID的通道")
		resMgmtClient, err = saveChannel(sdk, resMgmtClient)
		if err != nil {
			return fmt.Errorf("!保存带有交易ID的通道: %v", err)
		}
	} else {
		fmt.Println("加入通道已完成...")
	}
	ResMgmtClient = resMgmtClient
	return nil
}

func saveChannel(sdk *fabsdk.FabricSDK, resMgmtClient *resmgmt.Client) (*resmgmt.Client, error) {

	mspClient, err := mspclient.New(
		sdk.Context(),
		mspclient.WithOrg(Org1Name))
	if err != nil {
		log.Panicf("创建 msp 客户端失败: %s", err)
	}

	adminIdentity, err := mspClient.GetSigningIdentity(OrgAdmin)
	if err != nil {
		log.Panicf("获取标识失败: %s", err)
	}

	channelReq := resmgmt.SaveChannelRequest{
		ChannelID:         ChannelID,
		ChannelConfigPath: ChannelConfig,
		SigningIdentities: []msp.SigningIdentity{adminIdentity},
	}
	// 保存带有交易ID的渠道响应
	_, err = resMgmtClient.SaveChannel(channelReq,
		resmgmt.WithRetry(retry.DefaultResMgmtOpts),
		resmgmt.WithOrdererEndpoint(OrdererID))
	if err != nil {
		return nil, fmt.Errorf("创建应用通道失败: %v", err)
	}
	fmt.Println("通道创建成功...")
	return resMgmtClient, nil

}

func PeerJoinChannel(sdk *fabsdk.FabricSDK, orgName string) error {

	orgResMgmt, err := CreateResMgmtClient(sdk, OrgAdmin, orgName)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	err = orgResMgmt.JoinChannel(ChannelID,
		resmgmt.WithRetry(retry.DefaultResMgmtOpts),
		resmgmt.WithOrdererEndpoint(OrdererID))
	if err != nil {
		return fmt.Errorf("节点加入通道失败: %v", err)
	}
	fmt.Println(orgName + " 加入通道成功...")
	OrgResMgmt = orgResMgmt
	return nil
}
