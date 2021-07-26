# GoShell
Gos是一款专为[**蚁剑**](https://github.com/AntSwordProject/antSword/)编写的WebShell服务工具，使用蚁剑配合Gos内置的Web服务可轻松管理各种缺乏动态脚本环境的服务器

蚁剑中支持的Shell类型：

- CMDLINUX
- CUSTOM


Gos使用Go语言编写，理论上只要编译通过可支持各种系统（我猜的），以下是测试可用的系统：

 - Windows
 - Mac OS
 - Linux


推荐使用`CUSTOM`类型，`CMDLINUX`类型中的数据库管理功能需要系统中内置有数据库的客户端

而`CUSTOM`类型内置了~~各种数据库的客户端~~（暂时未添加，会有的）


让人高兴的是，Gos还简单内置了[**NPS**](https://github.com/ehang-io/nps)的客户端，配合NPS服务可在各种内网环境中自由穿梭

# 使用

```shell
## 启动**CUSTOM**类型的Web服务
./gos -m custom -web-addr 0.0.0.0:9010 -web-pwd ant -web-decoder plain

## 启动**CMDLINUX**类型的Web服务：该服务只支持Linux系统
./gos -m cmd -web-addr 0.0.0.0:9010 -web-pwd ant -web-decoder plain

## 启动**NPC**：配合NPS服务
./gos -m npc -server=1.1.1.1:8024 -vkey=key -type=tcp

```

# 更新

- 2021-07-26: 修复：修改文件时间因为非本地时区而导致的实际修改时间和提交修改时间不一致的问题
- 2021-07-08: 所有Web服务支持编码
- 2021-07-08: 支持蚁剑**CUSTOM**类型
- 2021-07-04: 支持蚁剑**CMDLINUX**类型


# Todo

- [ ] **CUSTOM** 服务内置MYSQL客户端
- [ ] **CUSTOM** 服务内置MSSQL客户端
- [x] **CUSTOM** 服务支持Windows
- [x] **CUSTOM**服务支持Base64编码


# 编译
```shell script
cd cmd/gos
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -trimpath gos.go
```

# 其他

关于Gos的名称，一开始是叫GoShell，后来发现名字太长打起来比较麻烦就改为Gos了