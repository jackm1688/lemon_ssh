package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"
)

//{"rs":1,"code":0,"address":"中国  广东省 深圳市 联通","ip":"27.38.33.41","isDomain":0}
type Domain struct {
	Rs       int    `json:"rs"`
	Code     int    `json:"code"`
	Address  string `json:"address"`
	Ip       string `json:"ip"`
	IsDomain int    `json:"isDomain"`
}

func QueryIp(size int, ips ...string) {
	var wg sync.WaitGroup
	wg.Add(1)
	pool := make(chan string, size)
	go func(p chan<- string) {
		defer wg.Done()
		for _, ip := range ips {
			p <- ip
		}
		close(p)
	}(pool)
	var client = &http.Client{}

	for v := range pool {
		url := fmt.Sprintf("https://ip.cn/api/index?ip=%s&type=1", v)
		wg.Add(1)
		go func(ip, urlAddr string) {
			defer wg.Done()
			req, err := http.NewRequest("GET", urlAddr, nil)
			if err != nil {
				fmt.Printf("create http.NewRequest fail:%v", err)
				return
			}
			req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.36")
			resp, err := client.Get(urlAddr)
			if err != nil {
				fmt.Printf("exec http get fail:%v\n", err)
				return
			}
			if resp.StatusCode != http.StatusOK {
				data, _ := ioutil.ReadAll(resp.Body)
				fmt.Printf("response code is %d, body:%s\n", resp.StatusCode, string(data))
				return
			}
			if resp != nil && resp.Body != nil {
				defer resp.Body.Close()
			}

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Printf("get http body fail:%v\n", err)
				return
			}

			var domain = Domain{}
			err = json.Unmarshal(body, &domain)
			if err != nil {
				fmt.Printf("json.Marshal fail:%v\n", err)
				return
			}
			info := strings.Split(domain.Address, " ")
			fmt.Printf("%s,%s,%s\n", domain.Ip, info[0], info[len(info)-1])
		}(v, url)
	}
	wg.Wait()
}

func main() {
	args := os.Args[1:]
	if len(args) >200{
		fmt.Println("最大200个")
		return
	}
	QueryIp(5, args...)
}
