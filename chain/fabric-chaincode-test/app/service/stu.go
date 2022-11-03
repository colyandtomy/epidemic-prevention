package service

import (
	"fmt"

	"github.com/hyperledger/fabric-chaincode-test/app/utils"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/pkg/errors"
)

func InitLedger() (string, error) {
	req := channel.Request{
		ChaincodeID: utils.ChaincodeID,
		Fcn:         "initLedger",
	}
	respone, err := utils.Client.Execute(req)
	if err != nil {
		return "", errors.WithMessage(err, "调用链码失败！")
	}
	return string(respone.TransactionID), nil
}

func QueryStuByUID(uid string) (string, error) {
	req := channel.Request{
		ChaincodeID: utils.ChaincodeID,
		Fcn:         "queryStuByUID",
		Args: [][]byte{
			[]byte(uid),
		},
	}

	rsp, err := utils.Client.Query(req)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	return string(rsp.Payload), nil
}

func QueryTest() {
	req := channel.Request{
		ChaincodeID: utils.ChaincodeID,
		Fcn:         "queryTest",
	}

	rsp, err := utils.Client.Query(req)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println("查询结果：", string(rsp.Payload))
}

func CreateStu(uid, name, age string) (string, error) {
	req := channel.Request{
		ChaincodeID: utils.ChaincodeID,
		Fcn:         "createStu",
		Args: [][]byte{
			[]byte(uid),
			[]byte(name),
			[]byte(age),
		},
	}

	rsp, err := utils.Client.Execute(req)
	if err != nil {
		return "", err
	}

	return string(rsp.Payload), nil
}
