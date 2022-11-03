// /*
// Copyright IBM Corp. 2016 All Rights Reserved.

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// 		 http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// */

// package main

// //WARNING - this chaincode's ID is hard-coded in chaincode_example04 to illustrate one way of
// //calling chaincode from a chaincode. If this example is modified, chaincode_example04.go has
// //to be modified as well with the new ID of chaincode_example02.
// //chaincode_example05 show's how chaincode ID can be passed in as a parameter instead of
// //hard-coding.

// import (
// 	"fmt"
// 	"strconv"

// 	"github.com/hyperledger/fabric/core/chaincode/shim"
// 	pb "github.com/hyperledger/fabric/protos/peer"
// )

// // SimpleChaincode example simple Chaincode implementation
// type SimpleChaincode struct {
// }

// func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
// 	fmt.Println("ex02 Init")
// 	_, args := stub.GetFunctionAndParameters()
// 	var A, B string    // Entities
// 	var Aval, Bval int // Asset holdings
// 	var err error

// 	if len(args) != 4 {
// 		return shim.Error("Incorrect number of arguments. Expecting 4!!")
// 	}

// 	// Initialize the chaincode
// 	A = args[0]
// 	Aval, err = strconv.Atoi(args[1])
// 	if err != nil {
// 		return shim.Error("Expecting integer value for asset holding")
// 	}
// 	B = args[2]
// 	Bval, err = strconv.Atoi(args[3])
// 	if err != nil {
// 		return shim.Error("Expecting integer value for asset holding")
// 	}
// 	fmt.Printf("Aval = %d, Bval = %d\n", Aval, Bval)

// 	// Write the state to the ledger
// 	err = stub.PutState(A, []byte(strconv.Itoa(Aval)))
// 	if err != nil {
// 		return shim.Error(err.Error())
// 	}

// 	err = stub.PutState(B, []byte(strconv.Itoa(Bval)))
// 	if err != nil {
// 		return shim.Error(err.Error())
// 	}

// 	return shim.Success(nil)
// }

// func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
// 	fmt.Println("ex02 Invoke")
// 	function, args := stub.GetFunctionAndParameters()
// 	if function == "invoke" {
// 		// Make payment of X units from A to B
// 		return t.invoke(stub, args)
// 	} else if function == "delete" {
// 		// Deletes an entity from its state
// 		return t.delete(stub, args)
// 	} else if function == "query" {
// 		// the old "Query" is now implemtned in invoke
// 		return t.query(stub, args)
// 	}

// 	return shim.Error("Invalid invoke function name. Expecting \"invoke\" \"delete\" \"query\"")
// }

// // Transaction makes payment of X units from A to B
// func (t *SimpleChaincode) invoke(stub shim.ChaincodeStubInterface, args []string) pb.Response {
// 	var A, B string    // Entities
// 	var Aval, Bval int // Asset holdings
// 	var X int          // Transaction value
// 	var err error

// 	if len(args) != 3 {
// 		return shim.Error("Incorrect number of arguments. Expecting 3")
// 	}

// 	A = args[0]
// 	B = args[1]

// 	// Get the state from the ledger
// 	// TODO: will be nice to have a GetAllState call to ledger
// 	Avalbytes, err := stub.GetState(A)
// 	if err != nil {
// 		return shim.Error("Failed to get state")
// 	}
// 	if Avalbytes == nil {
// 		return shim.Error("Entity not found")
// 	}
// 	Aval, _ = strconv.Atoi(string(Avalbytes))

// 	Bvalbytes, err := stub.GetState(B)
// 	if err != nil {
// 		return shim.Error("Failed to get state")
// 	}
// 	if Bvalbytes == nil {
// 		return shim.Error("Entity not found")
// 	}
// 	Bval, _ = strconv.Atoi(string(Bvalbytes))

// 	// Perform the execution
// 	X, err = strconv.Atoi(args[2])
// 	if err != nil {
// 		return shim.Error("Invalid transaction amount, expecting a integer value")
// 	}
// 	Aval = Aval - X
// 	Bval = Bval + X
// 	fmt.Printf("Aval = %d, Bval = %d\n", Aval, Bval)

// 	// Write the state back to the ledger
// 	err = stub.PutState(A, []byte(strconv.Itoa(Aval)))
// 	if err != nil {
// 		return shim.Error(err.Error())
// 	}

// 	err = stub.PutState(B, []byte(strconv.Itoa(Bval)))
// 	if err != nil {
// 		return shim.Error(err.Error())
// 	}

// 	return shim.Success(nil)
// }

// // Deletes an entity from state
// func (t *SimpleChaincode) delete(stub shim.ChaincodeStubInterface, args []string) pb.Response {
// 	if len(args) != 1 {
// 		return shim.Error("Incorrect number of arguments. Expecting 1")
// 	}

// 	A := args[0]

// 	// Delete the key from the state in ledger
// 	err := stub.DelState(A)
// 	if err != nil {
// 		return shim.Error("Failed to delete state")
// 	}

// 	return shim.Success(nil)
// }

// // query callback representing the query of a chaincode
// func (t *SimpleChaincode) query(stub shim.ChaincodeStubInterface, args []string) pb.Response {
// 	var A string // Entities
// 	var err error

// 	if len(args) != 1 {
// 		return shim.Error("Incorrect number of arguments. Expecting name of the person to query")
// 	}

// 	A = args[0]

// 	// Get the state from the ledger
// 	Avalbytes, err := stub.GetState(A)
// 	if err != nil {
// 		jsonResp := "{\"Error\":\"Failed to get state for " + A + "\"}"
// 		return shim.Error(jsonResp)
// 	}

// 	if Avalbytes == nil {
// 		jsonResp := "{\"Error\":\"Nil amount for " + A + "\"}"
// 		return shim.Error(jsonResp)
// 	}

// 	jsonResp := "{\"Name\":\"" + A + "\",\"Amount\":\"" + string(Avalbytes) + "\"}"
// 	fmt.Printf("Query Response:%s\n", jsonResp)
// 	return shim.Success(Avalbytes)
// }

// func main() {
// 	err := shim.Start(new(SimpleChaincode))
// 	if err != nil {
// 		fmt.Printf("Error starting Simple chaincode: %s", err)
// 	}
// }

package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

type SmartContract struct{}

type Stu struct {
	UID  string `json:"uid"`
	Name string `json:"name"`
	Age  string `json:"age"`
}

// 初始化。仅在智能合约被实例化时执行
func (s *SmartContract) Init(stub shim.ChaincodeStubInterface) sc.Response {
	stus := []Stu{
		{UID: "2211220000", Name: "AAA", Age: "20"},
		{UID: "2211220001", Name: "BBB", Age: "21"},
		{UID: "2211220002", Name: "CCC", Age: "22"},
		{UID: "2211220003", Name: "DDD", Age: "23"},
		{UID: "2211220004", Name: "EEE", Age: "24"},
	}

	for i := 0; i < len(stus); i++ {
		stuAsBytes, _ := json.Marshal(stus[i])
		stub.PutState(stus[i].UID, stuAsBytes)
	}

	return shim.Success(nil)
}

func (s *SmartContract) Invoke(stub shim.ChaincodeStubInterface) sc.Response {
	function, args := stub.GetFunctionAndParameters()
	if function == "initLedger" {
		return s.initLedger(stub)
	} else if function == "queryStuByUID" {
		return s.queryStuByUID(stub, args)
	} else if function == "createStu" {
		return s.createStu(stub, args)
	} else if function == "queryTest" {
		return s.queryTest(stub)
	} else if function == "queryAllStu" {
		return s.queryAllStu(stub)
	}

	return shim.Error("未找到相应的函数")
}

func (s *SmartContract) initLedger(stub shim.ChaincodeStubInterface) sc.Response {
	stus := []Stu{
		{UID: "2211220000", Name: "AAA", Age: "20"},
		{UID: "2211220001", Name: "BBB", Age: "21"},
		{UID: "2211220002", Name: "CCC", Age: "22"},
		{UID: "2211220003", Name: "DDD", Age: "23"},
		{UID: "2211220004", Name: "EEE", Age: "24"},
	}

	for i := 0; i < len(stus); i++ {
		stuAsBytes, _ := json.Marshal(stus[i])
		stub.PutState(stus[i].UID, stuAsBytes)
	}

	return shim.Success(nil)
}

func (s *SmartContract) createStu(stub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 3 {
		return shim.Error("请输入3个参数")
	}

	var stu = Stu{UID: args[0], Name: args[1], Age: args[2]}

	stuAsBytes, _ := json.Marshal(stu)
	stub.PutState(args[0], stuAsBytes)

	return shim.Success(nil)
}

func (s *SmartContract) queryStuByUID(stub shim.ChaincodeStubInterface, args []string) sc.Response {
	if len(args) != 1 {
		return shim.Error("请输入1个参数")
	}

	stuAsBytes, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(stuAsBytes)
}

func (s *SmartContract) queryAllStu(stub shim.ChaincodeStubInterface) sc.Response {

	return shim.Success(nil)
}

func (s *SmartContract) queryTest(stub shim.ChaincodeStubInterface) sc.Response {
	testB, err := json.Marshal(Stu{UID: "221122xxxx", Name: "Test", Age: "0"})
	if err != nil {
		return shim.Error("Error")
	}

	return shim.Success(testB)
}

func main() {
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}
