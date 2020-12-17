package util

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"io"
	"io/ioutil"
	"lemon_ssh/model"
	"os"
	"strings"
)


var Tip  = `
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
`

func GreenString(text string) string{
	return color.GreenString("%s",text)
}

func RedString(text string) string{
	return color.RedString("%s",text)
}

func YellowString(text string) string{
	return color.YellowString("%s",text)
}

func BlueString(text string) string{
	return color.BlueString("%s",text)
}

func ReadFile(filename string)(data []byte,err error){
	return ioutil.ReadFile(filename)
}

func ReadFromJSONFile(filename string)(users []model.User,err error){
	file,err := os.Open(filename)
	defer file.Close()
	if err != nil {
		return
	}

	reader := bufio.NewReader(file)
	for{
		line,err := reader.ReadBytes('\n')
		var user model.User
		if len(strings.TrimSpace(string(line))) > 0 {
			err = json.Unmarshal(line,&user)
			if user.Host != "" && user.User != ""{
				users = append(users,user)
			}
		}

		if err != nil{
			if err == io.EOF{
				break
			}

			fmt.Printf("read fail:%v,line:%v\n",err,string(line))
			break
		}
	}
	return
}
