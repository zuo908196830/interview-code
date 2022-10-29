# interview-code

### 服务端

在/tcp-server/conf/conf.json配置ip+port，以及是否开启Nagle，开启为true，关闭为false

在/tcp-server/utils/globalobj.go 中第34行，如果是linux环境使用"../conf/conf.json"，如果是windows环境使用"./conf/conf.json"

直接go run main.go即可

### 客户端

直接go run main,go