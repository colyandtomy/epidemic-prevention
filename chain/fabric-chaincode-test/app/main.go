package main

import (
	"fmt"

	"github.com/hyperledger/fabric-chaincode-test/app/service"
	"github.com/hyperledger/fabric-chaincode-test/app/utils"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

func org1Start(sdk *fabsdk.FabricSDK) {

	fmt.Println("开始Org1.......")
	//创建Org1通道
	err := utils.CreateChannel(sdk)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if !utils.Initialized {
		//Org1加入通道
		err = utils.PeerJoinChannel(sdk, utils.Org1Name)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		// //安装链码
		// err = utils.InstallCC()
		// if err != nil {
		// 	fmt.Println(err.Error())
		// 	return
		// }
		// //实例化链码
		// err = utils.InstantiateCC()
		// if err != nil {
		// 	fmt.Println(err.Error())
		// 	return
		// }
		utils.Initialized = !utils.Initialized
	}
}

func main() {
	sdk, err := utils.SetupSDK()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer sdk.Close()

	// org1Start(sdk)

	channelClient, err := utils.CreatChannelClient(sdk)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	utils.Client = channelClient

	// service.InitLedger()

	// service.QueryTest()

	msg, err := service.QueryStuByUID("2211220000")
	if err != nil {
		fmt.Println("查询失败 ", err)
	} else {
		fmt.Println("查询结果 ", msg)
	}

	msg, err = service.CreateStu("2211220039", "wzy", "24")
	if err != nil {
		fmt.Println("创建失败 ", err)
	} else {
		fmt.Println("创建成功 ", msg)
	}

	msg, err = service.QueryStuByUID("2211220039")
	if err != nil {
		fmt.Println("查询失败 ", err)
	} else {
		fmt.Println("查询结果 ", msg)
	}
}
