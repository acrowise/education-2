/**
  @Author : hanxiaodong
*/

package main

/**
姓名：楊xx，性别：男，

民族：漢，身份證號：P123456789

籍貫：xxx ，出生日期：1991年01月01日，		照片，

入學日期：2009年9月，畢(結)業日期：2013年7月，

學校名稱：朝陽科技大學，專業：資管系，

學歷類別：普通，學制：二年，

學習形式：普通全日制，層次：本科，

畢(結)業：畢業，證書編號：111111

 */
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
