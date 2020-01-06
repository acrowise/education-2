/**
  author: kevin
 */

package main

import (
	"os"
	"fmt"
	"github.com/kongyixueyuan.com/education/sdkInit"
	"github.com/kongyixueyuan.com/education/service"
	"encoding/json"
	"github.com/kongyixueyuan.com/education/web/controller"
	"github.com/kongyixueyuan.com/education/web"
)

const (
	configFile = "config.yaml"
	initialized = false
	EduCC = "educc"
)

func main() {

	initInfo := &sdkInit.InitInfo{

		ChannelID: "kevinkongyixueyuan",
		ChannelConfig: os.Getenv("GOPATH") + "/src/github.com/kongyixueyuan.com/education/fixtures/artifacts/channel.tx",

		OrgAdmin:"Admin",
		OrgName:"Org1",
		OrdererOrgName: "orderer.kevin.kongyixueyuan.com",

		ChaincodeID: EduCC,
		ChaincodeGoPath: os.Getenv("GOPATH"),
		ChaincodePath: "github.com/kongyixueyuan.com/education/chaincode/",
		UserName:"User1",
	}

	sdk, err := sdkInit.SetupSDK(configFile, initialized)
	if err != nil {
		fmt.Printf(err.Error())
		return
	}

	defer sdk.Close()

	err = sdkInit.CreateChannel(sdk, initInfo)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	channelClient, err := sdkInit.InstallAndInstantiateCC(sdk, initInfo)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(channelClient)

	//===========================================//

	serviceSetup := service.ServiceSetup{
		ChaincodeID:EduCC,
		Client:channelClient,
	}

	edu := service.Education{
		Name: "張三",
		Gender: "男",
		Nation: "漢",
		EntityID: "101",
		Place: "台北",
		BirthDay: "1991年01月01日",
		EnrollDate: "2009年9月",
		GraduationDate: "2013年7月",
		SchoolName: "台灣科技大學",
		Major: "資管系",
		QuaType: "普通",
		Length: "二年",
		Mode: "普通全日制",
		Level: "本科",
		Graduation: "畢業",
		CertNo: "111",
		Photo: "/static/photo/11.png",
	}

	edu2 := service.Education{
		Name: "李四",
		Gender: "男",
		Nation: "漢",
		EntityID: "102",
		Place: "台中",
		BirthDay: "1992年02月01日",
		EnrollDate: "2010年9月",
		GraduationDate: "2014年7月",
		SchoolName: "朝陽科技大學",
		Major: "資工系",
		QuaType: "普通",
		Length: "四年",
		Mode: "普通全日制",
		Level: "本科",
		Graduation: "畢業",
		CertNo: "222",
		Photo: "/static/photo/22.png",
	}

	msg, err := serviceSetup.SaveEdu(edu)
	if err != nil {
		fmt.Println(err.Error())
	}else {
		fmt.Println("資訊發布成功,交易編號為: " + msg)
	}

	msg, err = serviceSetup.SaveEdu(edu2)
	if err != nil {
		fmt.Println(err.Error())
	}else {
		fmt.Println("資訊發布成功,交易編號為: " + msg)
	}

	// 依據證書編號與名稱查詢
	result, err := serviceSetup.FindEduByCertNoAndName("222","李四")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		var edu service.Education
		json.Unmarshal(result, &edu)
		fmt.Println("證書編號與名稱查詢成功!：")
		fmt.Println(edu)
	}

	// 依據身份證號查詢
	result, err = serviceSetup.FindEduInfoByEntityID("101")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		var edu service.Education
		json.Unmarshal(result, &edu)
		fmt.Println("依據身份證號查詢成功!：")
		fmt.Println(edu)
	}

	// 修改/增加資訊
	info := service.Education{
		Name: "張三",
		Gender: "男",
		Nation: "漢",
		EntityID: "101",
		Place: "台北",
		BirthDay: "1991年01月01日",
		EnrollDate: "2013年9月",
		GraduationDate: "2019年7月",
		SchoolName: "台灣科技大學",
		Major: "資管系",
		QuaType: "普通",
		Length: "二年",
		Mode: "普通全日制",
		Level: "本科",
		Graduation: "畢業",
		CertNo: "333",
		Photo: "/static/photo/11.png",
	}
	msg, err = serviceSetup.ModifyEdu(info)
	if err != nil {
		fmt.Println(err.Error())
	}else {
		fmt.Println("操作成功,交易編號為: " + msg)
	}

	// 依據身份證號查詢資訊
	result, err = serviceSetup.FindEduInfoByEntityID("101")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		var edu service.Education
		json.Unmarshal(result, &edu)
		fmt.Println("身份證號查詢資訊成功!：")
		fmt.Println(edu)
	}

	// 依據證書編號與名稱查詢
	result, err = serviceSetup.FindEduByCertNoAndName("333","張三")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		var edu service.Education
		json.Unmarshal(result, &edu)
		fmt.Println("證書編號與名稱查詢成功!：")
		fmt.Println(edu)
	}

	/*// 删除資訊
	msg, err = serviceSetup.DelEdu("101")
	if err != nil {
		fmt.Println(err.Error())
	}else {
		fmt.Println("刪除成功,交易編號為: " + msg)
	}

	// 依據身份證號查詢
	result, err = serviceSetup.FindEduInfoByEntityID("101")
	if err != nil {
		fmt.Println(err.Error())
		fmt.Println("依據身份證號查詢失敗! 指定身份證號不存在或已被刪除...")
	} else {
		var edu service.Education
		json.Unmarshal(result, &edu)
		fmt.Println("身份證號查詢成功!：")
		fmt.Println(edu)
	}*/

	//===========================================//

	app := controller.Application{
		Setup: &serviceSetup,
	}
	web.WebStart(app)

}
