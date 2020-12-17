package mcmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"lemon_ssh/logger"
	"lemon_ssh/mlemon"
	"lemon_ssh/model"
	"lemon_ssh/util"
	"net/http"
	"strings"
	"sync"
)

func Exec(u *model.User,m *mlemon.MLemon,cmd string){

	if m.Err != nil {
		var tip string
		if strings.Contains(m.Err.Error(),"none password"){
			tip = "Password Wrong or Keystore file does not exist"
		} else if strings.Contains(m.Err.Error(),"unable to authenticate"){
			tip = "username is not  defined,please use -u username"
		}else{
			tip = m.Err.Error()
		}
		logger.Error("【%s@%s:%d】 connect fail:%v\n",u.User,u.Host,u.Port,util.RedString(tip))
		return
	}

	if result,err := m.Exec(cmd); err != nil {
		logger.Error("【%s@%s:%d】 exec->cmd:%s -- Err:%v",u.User,u.Host,u.Port,cmd,err)
	}else{
		info := util.RedString(fmt.Sprintf("%s@%s:%d",u.User,u.Host,u.Port))
		logger.INFO("%s%s%s - CMD:%s - status:%s - result:\n%s\n",
			util.GreenString("【"),info,util.GreenString("】"),util.YellowString(cmd),util.GreenString("OK"),result)
		util.WriteToFile(util.RecordFile,cmd,util.OpenErr)
	}
}

//文件上传
func ScpFile(u *model.User,m *mlemon.MLemon,remoteDir string,bufSize int,localFile ...string)(err error)  {
	for _,file := range localFile{
	  if err =	m.Scp(file,remoteDir,bufSize); err != nil {
		  break
	  }
	  info := fmt.Sprintf("scp %s %s@%s:%d:%s",file,u.User,u.Host,u.Port,remoteDir)
	  util.WriteToFile(util.RecordFile,info,util.OpenErr)
	}
	return
}

//{"rs":1,"code":0,"address":"中国  广东省 深圳市 联通","ip":"27.38.33.41","isDomain":0}
type Domain struct {
    Rs int `json:"rs"`
    Code int `json:"code"`
    Address string `json:"address"`
    Ip string `json:"ip"`
    IsDomain int `json:"isDomain"`
}

func QueryIp(size int,ips ...string)  {
	var wg sync.WaitGroup
	wg.Add(1)
	pool := make(chan string,size)
	go func(p chan<- string) {
		defer wg.Done()
		for _,ip := range  ips{
			p<-ip
		}
		close(p)
	}(pool)
	var client  =&http.Client{}

	for v := range pool{
		url := fmt.Sprintf("https://ip.cn/api/index?ip=%s&type=1",v)
		wg.Add(1)
		go func(ip,urlAddr string) {
			defer wg.Done()
			if req, err := http.NewRequest("GET", urlAddr, nil); err != nil {
				fmt.Printf("create http.NewRequest fail:%v",err)
			}else{
				req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.36")
				if resp, err := client.Get(urlAddr); err != nil {
					fmt.Printf("exec http get fail:%v\n",err)
				}else{
					if resp.StatusCode != http.StatusOK {
						data, _ := ioutil.ReadAll(resp.Body)
						fmt.Printf("response code is %d, body:%s\n", resp.StatusCode, string(data))
						return
					}

					if resp != nil && resp.Body != nil {
						defer resp.Body.Close()
					}

					if body, err := ioutil.ReadAll(resp.Body); err != nil {
						fmt.Printf("get http body fail:%v\n",err)
					} else{
						var domain = Domain{}
						if err = json.Unmarshal(body,&domain); err != nil {
							fmt.Printf("json.Marshal fail:%v\n",err)
						}
						info := strings.Split(domain.Address," ")
						fmt.Printf("%s,%s,%s\n",domain.Ip,info[0],info[len(info)-1])
					}
				}
			}
		}(v,url)
	}
	wg.Wait()
}