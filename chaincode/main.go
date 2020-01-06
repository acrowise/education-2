/**
  @Author : hanxiaodong
*/

package main

import (
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"fmt"
	"github.com/hyperledger/fabric/protos/peer"
)

type EducationChaincode struct {

}

func (t *EducationChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response{

	return shim.Success(nil)
}

func (t *EducationChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response{
	
	fun, args := stub.GetFunctionAndParameters()

	if fun == "addEdu"{
		return t.addEdu(stub, args)	// 增加訊息
	}else if fun == "queryEduByCertNoAndName" {
		return t.queryEduByCertNoAndName(stub, args)	// 依據證書編號及姓名查詢
	}else if fun == "queryEduInfoByEntityID" {
		return t.queryEduInfoByEntityID(stub, args)	// 依據身份證號碼及姓名查詢
	}else if fun == "updateEdu" {
		return t.updateEdu(stub, args)	// 依據證書編號更新訊息
	}else if fun == "delEdu"{
		return t.delEdu(stub, args)	// 依據證書編號刪除訊息
	}

	return shim.Error("指定的函數名稱錯誤!")

}

func main(){
	err := shim.Start(new(EducationChaincode))
	if err != nil{
		fmt.Printf("啟動ducationChaincode時發生錯誤!: %s", err)
	}
}
