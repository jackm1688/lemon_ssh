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