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

使用案列:<br>

场景1(用户+密码组合->单任务)：
 #对单台主机远程执行pwd命令
 -r cmd -u root -h 192.168.0.200 -P 22 -p abc-123 -c "pwd"

 #对0.200,0.201,0.202远程主机发起执行pwd命令
 -r cmd -u root -h 192.168.0.200,192.168.0.201,192.168.0.202 -P 22 -p abc-123 -c "pwd"

 场景2(用户+密码组合->多任务)：
 #对0.200,0.201,0.202远程主机发起执行pwd命令,开启2个线程执行,-e开启线程，-n指定线程的数量(最大为10)
 -r cmd -u root -h 192.168.0.200,192.168.0.201,192.168.0.202 -P 22 -p abc-123 -c "pwd" -e -n 2


 场景3(用户+信任文件->单任务)：
 #对单台主机远程执行pwd命令
  -r cmd -u root -h 192.168.0.200 -P 22 -k /root/.ssh/id_rsa  -c "pwd"

  #对0.200,0.201,0.202远程主机发起执行pwd命令
  -r cmd -u root -h 192.168.0.200,192.168.0.201,192.168.0.202 -P 22 -k /root/.ssh/id_rsa -c "pwd"

 场景4(用户+信任文件->多任务):
 #对0.200,0.201,0.202远程主机发起执行pwd命令,开启2个线程执行,-e开启线程，-n指定线程的数量(最大为10)
 -r cmd -u root -h 192.168.0.200,192.168.0.201,192.168.0.202 -P 22 -k /root/.ssh/id_rsa -c "pwd" -e -n 2

 场景5(指定IP列表文件->单任务)
 -r cmd -f host.json.list

 场景6(指定IP列表文件->多任务)
 -r cmd -f host.json.list -e -n 3

 场景7(指定IP列表文件+指定信任文件->单任务)
 -r cmd -f host.json.rsa.list  -k /root/.ssh/id_rsa

 场景8(指定IP列表文件+指定信任文件->多任务)
 -r cmd -f host.json.rsa.list  -k /root/.ssh/id_rsa -e -n 3

 场景9(指定IP范围->单任务)
 -r cmd -u root -l -h 192.168.0.70-72  -P 22 -p abc-123 -c "pwd;sleep 5"
 -r cmd -u root -l -h "192.168.0.70-72;192.168.0.50-52"  -P 22 -p abc-123 -c "pwd;sleep 5"

 场景10(指定IP范围->多任务)
 -r cmd -u root -l -h 192.168.0.70-72  -P 22 -p abc-123 -c "pwd;sleep 5" -n 5 -e
 -r cmd -u root -l -h "192.168.0.70-72;192.168.0.50-52"  -P 22 -p abc-123 -c "pwd;sleep 5"

 场景11(指定IP范围+使用信任关系->单任务)
 -r cmd -u root -l -h 192.168.0.70-72 -P 22 -k /root/.ssh/id_rsa -c "pwd;sleep 5"
 -r cmd -u root -l -h "192.168.0.70-72;192.168.0.50-52"  -P 22 -k /root/.ssh/id_rsa -c "pwd;sleep 5"

 场景12(指定IP范围+使用信任关系->多任务)
 -r cmd -u root -l -h 192.168.0.70-72 -P 22 -k /root/.ssh/id_rsa -c "pwd;sleep 5" -n 5 -e
 -r cmd -u root -l -h "192.168.0.70-72;192.168.0.50-52" -P 22 -k /root/.ssh/id_rsa -c "pwd;sleep 5"

host.json.list
{"user": "root","password":"abc-123","host":"192.168.0.50","port":22,"cmd":"pwd;sleep 5","timeout":15}
{"user": "root","password":"abc-123","host":"192.168.0.51","port":22,"cmd":"date;sleep 5","timeout":15}
{"user": "root","password":"abc-123","host":"192.168.0.52","port":22,"cmd":"who;sleep 5","timeout":15}

host.json.rsa.list
{"user": "root","host":"192.168.0.105","port":22,"cmd":"pwd;sleep 5","timeout":15}
{"user": "root","host":"192.168.0.50","port":22,"cmd":"date;sleep 5","timeout":15}
{"user": "root","host":"192.168.0.51","port":22,"cmd":"who;sleep 5","timeout":15}



