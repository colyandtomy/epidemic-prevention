package utils

import (
	"os"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt"
)

var (
	ChannelID        = "mychannel"
	ChannelConfig    = "/home/wzy/go/src/github.com/hyperledger/fabric-chaincode-test/network/channel-artifacts/channel.tx"
	OrdererID        = "orderer.example.com"
	OrgAdmin         = "Admin"
	OrdererOrgName   = "OrdererOrg"
	Org1Name         = "Org1"
	Org2Name         = "Org2"
	ChaincodeGoPath  = os.Getenv("GOPATH")
	ChaincodePath    = "/home/go/src/github.com/hyperledger/fabric-chaincode-test/chaincode/go"
	UserName         = "User1"
	ChaincodeID      = "stu"
	ConfigFile       = "./config/config.yaml"
	Initialized      = false
	Installed        = false
	ChaincodeVersion = "1.0"
	ResMgmtClient    *resmgmt.Client
	OrgResMgmt       *resmgmt.Client
	Client           *channel.Client
)
