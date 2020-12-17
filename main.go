package main

import (
	"flag"
	"fmt"
	"lemon_ssh/mlemon"
	"lemon_ssh/model"
	"lemon_ssh/mssh/mcmd"
	"lemon_ssh/util"
	"os"
	"path"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	HISTORY_DIR       = "./history"
	HISOTRY_FILE_NAME = "record.db"
	MAXSIZE           = 2000
)

var (
	run      string
	user     string
	host     string
	password string
	cmd      string
	port     int
	num      int
	enable   bool
	timeout  int
	keyFile  string
	use      bool
	file     string
	list     bool
	lFile    string
	rtd      string
	size     int
)

var RecordList []*model.LogRecord

func init() {
	crateHistoryDir()
	RecordList = util.ReadFromFile(util.RecordFile, util.OpenErr)
}

func main() {

	defer func() {
		if util.OpenErr != nil {
			_ = util.RecordFile.Close()
		}
	}()

	ParseArgs()
	flag.Parse()

	if use {
		fmt.Println(util.Tip)
		return
	}

	if run == "" {
		flag.Usage()
		return
	}
	if enable == true {
		//设置缓冲池大小
		if num <= 0 {
			num = 1
		}
		//设置缓冲池大小
		if num > 15 {
			num = 3
		}
	}
	switch run {
	case "cmd":
		cmdCase()
	case "history":
		getHistory()
	default:
		fmt.Println("未知的选项")
	}
}

func getHistory() {
	if util.RecordFile != nil {
		for i, v := range RecordList {
			fmt.Printf("%4d\t%s -> %s \n", i, v.CreateAt, v.Content)
		}
	}
}

func cmdCase() {
	if file != "" && len(file) >= 0 {
		userList, err := util.ReadFromJSONFile(file)
		if err != nil {
			fmt.Printf("读取文件出错:%v\n", err)
			return
		}
		if len(userList) == 0 {
			fmt.Println("IP列表为空")
			return
		}
		userExec(userList)
	} else if list {
		hosts := make([]string, 0)
		list := strings.Split(host, ";")
		for _, v := range list {
			item := strings.Split(v, "-")
			ips := strings.Split(item[0], ".")
			var start, end int64
			var err error
			if start, err = strconv.ParseInt(ips[len(ips)-1], 10, 10); err != nil {
				fmt.Printf("获取起始数失败:%v,value:%v\n", err, item)
				continue
			}

			if end, err = strconv.ParseInt(item[1], 10, 10); err != nil {
				fmt.Printf("parse fail:%v, valeu:%v\n", err, item)
				continue
			}

			for i := start; i <= end; i++ {
				ips[len(ips)-1] = fmt.Sprintf("%d", i)
				ip := strings.Join(ips, ".")
				hosts = append(hosts, ip)
			}
		}
		if len(hosts) == 0 {
			fmt.Printf("IP列表为空")
			return
		}
		exec(hosts...)
	} else {
		hosts := strings.Split(host, ",")
		exec(hosts...)
	}
}

func userExec(users []model.User) {
	if enable {
		var wg sync.WaitGroup
		taskPool := make(chan model.User, num)
		wg.Add(1)
		//producer
		go func(pool chan<- model.User) {
			defer wg.Done()
			for _, v := range users {
				pool <- v
			}
			close(taskPool)
		}(taskPool)

		for v := range taskPool {
			wg.Add(1)
			go func(u model.User) {
				defer wg.Done()
				m := MakeMLemon(&u)
				scpFile(&u, m) //上传文件
				mcmd.Exec(&u, m, u.Cmd)
				fmt.Println()
				runtime.Gosched()
			}(v)
		}
		wg.Wait()
	} else {
		for _, u := range users {
			m := MakeMLemon(&u)
			scpFile(&u, m) //上传文件
			mcmd.Exec(&u, m, u.Cmd)
			fmt.Println()
		}
	}
}

func exec(hosts ...string) {
	if enable {
		var wg sync.WaitGroup
		hostPool := make(chan string, num)
		wg.Add(1)
		//producer
		go func(pool chan<- string) {
			defer wg.Done()
			for _, v := range hosts {
				pool <- v
			}
			close(pool)
		}(hostPool)

		for v := range hostPool {
			wg.Add(1)
			go func(ip string) {
				defer wg.Done()
				u, m := MakeMLemonAndUser(user, ip, port, password,
					keyFile, time.Duration(timeout))
				scpFile(u, m)
				mcmd.Exec(u, m, cmd)
				fmt.Println()
				runtime.Gosched()
			}(v)
		}
		wg.Wait()
	} else {
		for _, v := range hosts {
			u, m := MakeMLemonAndUser(user, v, port, password,
				keyFile, time.Duration(timeout))
			scpFile(u, m)        //上传文件
			mcmd.Exec(u, m, cmd) //执行命令
			fmt.Println()
		}
	}
}

func MakeMLemonAndUser(user string, host string, port int,
	password string, key string, timeout time.Duration) (*model.User, *mlemon.MLemon) {
	u := &model.User{
		User:     user,
		Password: password,
		Host:     host,
		Port:     port,
		Timeout:  timeout,
		Key:      key,
	}
	return u, mlemon.NewMLemon(u)
}

func MakeMLemon(user *model.User) *mlemon.MLemon {
	return mlemon.NewMLemon(user)
}

func scpFile(u *model.User, m *mlemon.MLemon) {
	if lFile != "" {
		if size <= 0 || size >= 4096 {
			size = 1024
		}
		var files []string = make([]string, 0)
		if strings.Contains(lFile, ";") {
			files = strings.Split(lFile, ";")
		} else {
			files = append(files, lFile)
		}
		_ = mcmd.ScpFile(u, m, rtd, size, files...)
	}
}

func ParseArgs() {

	flag.StringVar(&run, "r", "", "-r <cmd|history> 指定运行模式 cmd 运行命令 history 历史记录")
	flag.StringVar(&user, "u", "", "-u <username>")
	flag.StringVar(&host, "h", "", "-h <host|host1,hostN>")
	flag.StringVar(&cmd, "c", "", "-c <cmd1;cmdN>")
	flag.StringVar(&password, "p", "", "-p <password>")
	flag.IntVar(&port, "P", 22, "-P <port>")
	flag.IntVar(&num, "n", 1, "-n <1-10> thread num(n >=1 and n <=10) default 1")
	flag.BoolVar(&enable, "e", false, "-e default false")
	flag.IntVar(&timeout, "t", 15, "-t <timeout> default 15s")
	flag.StringVar(&keyFile, "k", "", "-k /root/.ssh/id_rsa")
	flag.BoolVar(&use, "use", false, "-use 使用案列")
	flag.StringVar(&file, "f", "", "-f ip.json.list 指定IP列表文件")
	flag.BoolVar(&list, "l", false, "-l 开启生成IP列表功能")
	flag.StringVar(&lFile, "lf", "", "-lf 指定本地文件")
	flag.StringVar(&rtd, "dst", "/root", "-dst 指定远端目录 默认为:/root")
	flag.IntVar(&size, "s", 1024, "-s size 指定读取文件缓冲大小,默认为1024")

}

func crateHistoryDir() {
	info, _ := os.Stat(HISTORY_DIR)
	if !info.IsDir() {
		_ = os.Mkdir(HISTORY_DIR, 0775)
	} else {
		historyFile := path.Join(HISTORY_DIR, HISOTRY_FILE_NAME)
		util.RecordFile, util.OpenErr = os.OpenFile(historyFile, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0755)
		if len(RecordList) == MAXSIZE {
			bak := path.Join(HISTORY_DIR, fmt.Sprintf("%s_%s", HISOTRY_FILE_NAME, time.Now().Format("2006-01-02")))
			os.Rename(historyFile, bak)
			util.RecordFile, util.OpenErr = os.OpenFile(historyFile, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0755)
		}
	}
}
