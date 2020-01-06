/**
  @Prject: goProjects
  @Dev_Software: GoLand
  @File : upload
  @Time : 2018/10/24 11:58 
  @Author : hanxiaodong
*/

package controller

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"crypto/rand"
	"path/filepath"
	"os"
	"mime"
	"log"
)


func (app *Application) UploadFile(w http.ResponseWriter, r *http.Request)  {

	start := "{"
	content := ""
	end := "}"

	file, _, err := r.FormFile("file")
	if err != nil {
		content = "\"error\":1,\"result\":{\"msg\":\"指定無效的檔案\",\"path\":\"\"}"
		w.Write([]byte(start + content + end))
		return
	}
	defer file.Close()

	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		content = "\"error\":1,\"result\":{\"msg\":\"無法讀取檔案內容\",\"path\":\"\"}"
		w.Write([]byte(start + content + end))
		return
	}

	filetype := http.DetectContentType(fileBytes)
	//log.Println("filetype = " + filetype)
	switch filetype {
	case "image/jpeg", "image/jpg":
	case "image/gif", "image/png":
	case "application/pdf":
		break
	default:
		content = "\"error\":1,\"result\":{\"msg\":\"檔案類型錯誤\",\"path\":\"\"}"
		w.Write([]byte(start + content + end))
		return
	}

	fileName := randToken(12)	// 指定檔案名稱
	fileEndings, err := mime.ExtensionsByType(filetype)	// 取得檔案副檔名
	//log.Println("fileEndings = " + fileEndings[0])
	// 指定檔案儲存路徑
	newPath := filepath.Join("web", "static", "photo", fileName + fileEndings[0])
	//fmt.Printf("FileType: %s, File: %s\n", filetype, newPath)

	newFile, err := os.Create(newPath)
	if err != nil {
		log.Println("產生檔案失敗：" + err.Error())
		content = "\"error\":1,\"result\":{\"msg\":\"產生檔案失敗\",\"path\":\"\"}"
		w.Write([]byte(start + content + end))
		return
	}
	defer newFile.Close()

	if _, err := newFile.Write(fileBytes); err != nil || newFile.Close() != nil {
		log.Println("寫入檔案失敗：" + err.Error())
		content = "\"error\":1,\"result\":{\"msg\":\"儲存檔案內容失敗\",\"path\":\"\"}"
		w.Write([]byte(start + content + end))
		return
	}

	path := "/static/photo/" + fileName + fileEndings[0]
	content = "\"error\":0,\"result\":{\"fileType\":\"image/png\",\"path\":\"" + path + "\",\"fileName\":\"ce73ac68d0d93de80d925b5a.png\"}"
	w.Write([]byte(start + content + end))
	return
}

func randToken(len int) string {
	b := make([]byte, len)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

