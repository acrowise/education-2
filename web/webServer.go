/**
  @Author : hanxiaodong
*/

package web

import (
	"net/http"
	"fmt"
	"github.com/kongyixueyuan.com/education/web/controller"
)


// 啟用Web服務並指定路由資訊
func WebStart(app controller.Application)  {

	fs:= http.FileServer(http.Dir("web/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// 指定路由資訊(請求匹配)
	http.HandleFunc("/", app.LoginView)
	http.HandleFunc("/login", app.Login)
	http.HandleFunc("/loginout", app.LoginOut)

	http.HandleFunc("/index", app.Index)
	http.HandleFunc("/help", app.Help)

	http.HandleFunc("/addEduInfo", app.AddEduShow)	// 顯示增加資訊頁面
	http.HandleFunc("/addEdu", app.AddEdu)	// 提交資計請求

	http.HandleFunc("/queryPage", app.QueryPage)	// 轉至證書編號與姓名查詢頁面
	http.HandleFunc("/query", app.FindCertByNoAndName)	// 依據證書編號與姓名查詢

	http.HandleFunc("/queryPage2", app.QueryPage2)	// 轉至身份證號查詢頁面
	http.HandleFunc("/query2", app.FindByID)	// 依據身份證號查詢


	http.HandleFunc("/modifyPage", app.ModifyShow)	// 修改資訊頁面
	http.HandleFunc("/modify", app.Modify)	//  修改資訊

	http.HandleFunc("/upload", app.UploadFile)

	fmt.Println("啟用Web服務,Listen port: 9000")
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("Web服務啟用失敗!: %v", err)
	}

}



