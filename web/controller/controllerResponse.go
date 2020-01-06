/**
  @Author : hanxiaodong
*/

package controller

import (
	"net/http"
	"path/filepath"
	"html/template"
	"fmt"
)

func ShowView(w http.ResponseWriter, r *http.Request, templateName string, data interface{})  {

	// 指定視圖所在路徑
	pagePath := filepath.Join("web", "tpl", templateName)

	resultTemplate, err := template.ParseFiles(pagePath)
	if err != nil {
		fmt.Printf("建立模板實例錯誤!: %v", err)
		return
	}

	err = resultTemplate.Execute(w, data)
	if err != nil {
		fmt.Printf("模板中融合資料時發生錯誤!: %v", err)
		//fmt.Fprintf(w, "顯示在用戶端瀏覽器中錯誤!")
		return
	}

}
