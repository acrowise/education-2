/**
  @Author : hanxiaodong
*/
package service

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"fmt"
	"time"
	"github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
)

type Education struct {
	ObjectType	string	`json:"docType"`
	Name	string	`json:"Name"`		// 姓名
	Gender	string	`json:"Gender"`		// 性别
	Nation	string	`json:"Nation"`		// 民族
	EntityID	string	`json:"EntityID"`		// 身份證號
	Place	string	`json:"Place"`		// 籍貫
	BirthDay	string	`json:"BirthDay"`		// 出生日期
	EnrollDate	string	`json:"EnrollDate"`		// 入學日期
	GraduationDate	string	`json:"GraduationDate"`	// 畢(結)業日期
	SchoolName	string	`json:"SchoolName"`	// 學校名稱
	Major	string	`json:"Major"`	// 專業
	QuaType	string	`json:"QuaType"`	// 學歷類別
	Length	string	`json:"Length"`	// 學制
	Mode	string	`json:"Mode"`	// 學習形式
	Level	string	`json:"Level"`	// 層次
	Graduation	string	`json:"Graduation"`	// 畢(結)業
	CertNo	string	`json:"CertNo"`	// 證書編號
	Photo	string	`json:"Photo"`	// 照片
	Historys	[]HistoryItem	// 目前edu的歷史記錄
}

type HistoryItem struct {
	TxId	string
	Education	Education
}

type ServiceSetup struct {
	ChaincodeID	string
	Client	*channel.Client
}

func regitserEvent(client *channel.Client, chaincodeID, eventID string) (fab.Registration, <-chan *fab.CCEvent) {

	reg, notifier, err := client.RegisterChaincodeEvent(chaincodeID, eventID)
	if err != nil {
		fmt.Println("註冊智能合約失敗!: %s", err)
	}
	return reg, notifier
}

func eventResult(notifier <-chan *fab.CCEvent, eventID string) error {
	select {
	case ccEvent := <-notifier:
		fmt.Printf("接收到智能合約事件: %v\n", ccEvent)
	case <-time.After(time.Second * 20):
		return fmt.Errorf("不能指定的事件ID接收到相對應的智能合約事件(%s)", eventID)
	}
	return nil
}
