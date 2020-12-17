柠檬运维工具(lemon):批量执行linux命令及shell脚本

帮助信息:<br>
go run main.go <br>

  -P int<br>
        -P <port> (default 22)<br>
  -c string<br>
        -c <cmd1;cmdN><br>
  -dst string<br>
        -dst 指定远端目录 默认为:/root (default "/root")<br>
  -e    -e default false<br>
  -f string<br>
        -f ip.json.list 指定IP列表文件<br>
  -h string<br>
        -h <host|host1,hostN><br>
  -k string<br>
        -k /root/.ssh/id_rsa<br>
  -l    -l 开启生成IP列表功能<br>
  -lf string<br>
        -lf 指定本地文件<br>
  -n int<br>
        -n <1-10> thread num(n >=1 and n <=10) default 1 (default 1)<br>
  -p string<br>
        -p <password><br>
  -r string<br>
        -r <cmd|history> 指定运行模式 cmd 运行命令 history 历史记录<br>
  -s int<br>
        -s size 指定读取文件缓冲大小,默认为1024 (default 1024)<br>
  -t int<br>
        -t <timeout> default 15s (default 15)<br>
  -u string<br>
        -u <username><br>
  -use<br>
        -use 使用案列<br>

<br><br>
Golang 支持在一个平台下生成另一个平台可执行程序的交叉编译功能。<br>
Mac下编译Linux, Windows平台的64位可执行程序：<br>
<br><br>
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build test.go<br>
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build test.go<br>
Linux下编译Mac, Windows平台的64位可执行程序：<br>
<br>
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build test.go<br>
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build test.go<br>
Windows下编译Mac, Linux平台的64位可执行程序：<br>

SET CGO_ENABLED=0<br>
SET GOOS=darwin3<br>
SET GOARCH=amd64<br>
go build test.go<br>

SET CGO_ENABLED=0<br>
SET GOOS=linux<br>
SET GOARCH=amd64<br>
go build test.go<br>
GOOS：目标可执行程序运行操作系统，支持 darwin，freebsd，linux，windows<br>
GOARCH：目标可执行程序操作系统构架，包括 386，amd64，arm<br>

Golang version 1.5以前版本在首次交叉编译时还需要配置交叉编译环境：

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 ./make.bash
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 ./make.bash