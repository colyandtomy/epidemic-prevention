# fabric-chaincode-test

## 简介

本项目是一个测试项目，该项目使用了fabric官方的测试网络，并部署了自己改写的stu链码。stu链码实现了简单的学生信息管理（目前仅实现了增加学生信息和查询学生信息）。项目通过fabric-sdk-go实现了与测试网络的连接和链码调用及查询。

## 部署说明

1. 将fabric-chaincode-test.tar.gz解压
    ```bash
    fabric-chaincode-test.tar.gz
    ```
2. 将解压后得到的目录移动到${GOPATH}/src/hyperledger/下（注意：也可部署到其他位置，但需要该代码和配置文件，建议按照本文要求部署）
3. 进入fabric-chaincode-test/network/目录，通过byfn.sh脚本部署fabric网络和部署链码：

    ```bash
    ./byfn.sh up
    ```
    （注：本测试项目使用的是fabric官方提供的测试网络。如果你在之前曾经启动了官方的测试网络，请执行./byfn.sh restart将原先的网络关闭并重新启动网络和部署链码）

4. 进入app目录，执行go run main.go
5. 如果看到终端显示如下内容，则表示测试项目运行成功：

    ```bash
    Fabric SDK初始化成功
    通道客户端创建成功，可以利用此客户端调用链码进行查询或执行事务.
    查询结果  {"uid":"2211220000","name":"AAA","age":"20"}
    创建成功 
    查询结果  {"uid":"2211220039","name":"wzy","age":"24"}
    ```
