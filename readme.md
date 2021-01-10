## 软件测试覆盖率<br>
1、定义：覆盖率是用来度量测试完整性的一个手段，同时也是测试技术有效性的一个度量。<br>
  2、计算：覆盖率=（至少被执行一次的item数）/item的总数<br>
  3、特点<br>
     1）通过覆盖率数据，可以检测我们的测试是否充分<br>
     2）分析出测试的弱点在哪方面<br>
     3）指导我们设计能够增加覆盖率的测试用例，有效提高测试质量，但是测试用例设计不能一味追求覆盖率，因为测试成本随覆盖率的增加而增加。<br>


## 软件测试覆盖率分类 <br>
覆盖率按照测试方法大体上可以划分为三大类，即白盒覆盖（white-Box Coverage）、灰盒覆盖（Gray-Box coverage）和黑盒覆盖（Black-Box Coverage）。

### 白盒覆盖率（white-Box Coverage） <br>
白盒覆盖率中使用的最常见的就是逻辑覆盖率（Logical Coverage ），也叫代码覆盖率（Code Coverage）或者结构化覆盖率（Structural Coverage），<br>
我们常见的逻辑覆盖包括：语句覆盖、判定覆盖、条件覆盖、判定条件覆盖、条件组合覆盖、路径覆盖。<br>

1、语句覆盖（Statement Coverage）

     1）定义：在测试时，运行被测程序后，程序中被执行的可执行语句的比率。
     2）计算公式：语句覆盖率=（至少被执行一次的语句数量）/（可执行的语句总数）
     3）100%语句覆盖率含义：在测试时，首先设计若干个测试用例，然后运行被测程序，使程序中的每个可执行语句至少执行一次。
     4）特点：语句覆盖可以检验每个可执行语句，但是即使语句覆盖率达到了100%，也会有缺陷发现不了，所以覆盖率只是我们度量的手段。
2、判定覆盖（Decision Coverage）/分支覆盖率（Branch Coverage）

     1）定义：在测试时，运行被测程序后，程序中所有判断语句的取真分支和取假分支被执行到的比率。
     2）计算公式：判定覆盖率=（判定结果被评价的次数）/（判定结果的总数）
     3）100%条件覆盖率含义：在测试时，首先设计若干个测试用例，然后运行测试程序，使得程序中每个判断的取真分支和取假分支至少经历一次，即判断的真假值均曾被满足。
     4）特点
       （1）若判定覆盖达到100%，则语句覆盖必为100%。
       （2）即使判定覆盖率达到了100%，也会有缺陷发现不了。
3、条件覆盖（Condition Coverage）

     1）定义：在测试时，运行被测程序后，程序中所有判断语句中每个条件的可能取值（真值和假值）出现过的比率。
     2）计算公式：条件覆盖率=（条件操作数值至少被评价一次的数量）/（条件操作数值的总数）
     3）100%条件覆盖率含义：在测试时，首先设计若干个测试用例，然后运行被测试程序，要使每个判断中每个条件的可能取值至少满足一次。
     4）特点：覆盖条件的测试用例不一定覆盖判定。
4、判定-条件覆盖（Decision Condition Coverage）/分支条件覆盖（Branch Condition Coverage）

     1)定义：在测试时，运行被测程序后，程序中所有判断语句中每个条件的可能取值（真值和假值）和每个判断本身的判定结果（为真为假）出现的比率。
     2）计算公式：判定-条件覆盖率=（条件操作数值或判定结果至少被评价一次的数量）/（条件操作数值的总数+判定结果的总数）
     3）100%判定-条件覆盖率含义：设计足够的测试用例，使得判断中每个条件的所有可能取值至少执行一次，同时每个判断本身的所有可能结果至少执行一次。换言之，即是要求各个判断的所有的可能的取值组合至少执行一次。
     4）特点
       （1）判定-条件覆盖率实际上就是判定覆盖率和条件覆盖率的组合。
       （2）采用判定-条件覆盖，逻辑表达式中的错误不一定能够查得出来。
5、条件组合覆盖（Condition combination coverage）

     1)定义：在测试时，运行被测程序后，所有语句中原子条件所有的可能的取值结果组合出现过的比率。
     2）计算公式：条件组合覆盖率=（至少被执行一次的条件组合）/（总的可能的条件组合数）
     3）100%条件组合覆盖率含义：设计足够的测试用例，使得判断中条件的各种可能组合至少出现过一次。
     4）特点：若条件组合覆盖率为100%，则语句覆盖率、判定覆盖率、条件覆盖率和判定-条件覆盖率必为100%。
6、路径覆盖（Path Coverage）

     1)定义：在测试时，运行被测程序后，程序中所有可能的路径被执行的比率。
     2）计算公式：路径覆盖率=（至少被执行一次的路径数）/（总的路径数）
     3）100%路径覆盖率含义：设计足够的测试用例，要求覆盖程序中所有可能的路径。
     4）特点
       （1）路径覆盖比判定条件覆盖更强，但是不能包含判定条件覆盖。
       （2）若路径覆盖率为100%，则语句覆盖率、判定覆盖率必为100%。
小结：逻辑覆盖率可以作为软件测试的一个度量，但是，即使达到了100%的逻辑覆盖率，仍然无法保证程序的正确性。

### 灰盒覆盖率（Gray-Box Coverage）  <br>

函数覆盖和接口覆盖可以归为灰盒测试的范畴。

1、函数覆盖

     1)定义：它表示在测试中，有哪些函数被测试到了，其被测试到的频率有多大，这些函数在系统所有函数中占的比例有多大。
     2）计算公式：函数覆盖=（至少被执行一次的函数数量）/（系统中函数的总数）
     3）特点：是针对一个系统或者子系统测试的。
2、接口覆盖（Interface Coverage）/入口点覆盖（Entry-Point Coverage）

     1)定义：要求通过设计一定的用例使得系统的每个接口被测试到。
     2）计算公式：接口覆盖=（至少被执行一次的接口数量）/（系统中接口的总数）         
### 黑盒覆盖率（Black-Box Coverage）  <br>

在实际测试中，与黑盒相关的覆盖率比较少，主要是功能覆盖率（Function Coverage），其中最常见的是需求覆盖。

需求覆盖

     1)定义：它表示在测试中，有哪些函数被测试到了，其被测试到的频率有多大，这些函数在系统所有函数中占的比例有多大通过设计一定的测试用例，要求每个需求点都被测试到。
     2）计算公式：需求覆盖=（被验证到的需求数量）/（总的需求总数）   
<br>
=============================================================================================<br>
golang的测试框架及测试分类<br>
a.单元测试(程序员自己来做)<br>
   1.单元测试<br>
      a.对于各个分支进行测试，如果不符合预期则失败<br>
      b.使用testing.T这个对象进行单元测试控制<br>
   2.压力测试(基准测试)<br>
     a.主要用来做性能测试<br>
     b.go test自动会执行所有的基准测试，并且打印执行耗时统计<br>
b.功能测试(专门测试人没来做)<br>

c.自动测试框架<br>
  1.testing<br>
     a.testing包提供了自动化测试相关的框架<br>
     b.支持单元测试和压力测试<br>

     import(<br>
       "testing"<br>
     )<br>

2.go中的测试规范<br>
  a.用来测试的代码必须以_test.go<br>
  b.单元测试的函数名必须以Test开头，并且只有一个参数,参数类型是*Testing.T<br>
  c.基准测试或压力测试必须以Benchmark开头，并且只有一个参数，参数类型是*Testing.B<br>


### 执行单元测试命令 <br>
go test -v -run TestAdd <br>
=== RUN   TestAdd <br>
    demo_test.go:10: sum is 4 <br>
--- PASS: TestAdd (0.00s) <br>
PASS <br>
ok      go-demo-test/demo       0.599s <br>

<br>

### 执行性能测试 <br>

go test -v -bench BenchmarkAdd .<br>
=== RUN   TestAdd<br>
    demo_test.go:10: sum is 4<br>
--- PASS: TestAdd (0.00s)<br>
goos: darwin<br>
goarch: amd64<br>
pkg: go-demo-test/demo<br>
BenchmarkAdd<br>
BenchmarkAdd-4          1000000000               0.418 ns/op<br>
PASS<br>
ok      go-demo-test/demo       0.573s<br>

### 测试所有的单元的
go test -v  ...


### 覆盖率生成
Golang单元测试覆盖率的生成也简单到令人发指。两步：<br>

  执行go test时指定-coverprofile参数收集覆盖率数据；<br>
  执行go tool cover生成文本、html等可视化格式的覆盖率报告。<br>
<br>

go test -v -coverprofile=cover.out /Users/mengfanzhen/go/src/lemon_ssh/m
main.go  mlemon/  model/   mssh/  
bogon:lemon_ssh JackMeng$ go test -v -coverprofile=cover.out /Users/mengfanzhen/go/src/lemon_ssh/mssh/m
mcmd/         msftp.go      mssh.go       mssh_test.go  
bogon:lemon_ssh JackMeng$ go test -v -coverprofile=cover.out /Users/mengfanzhen/go/src/lemon_ssh/mssh/
=== RUN   TestSSH
--- PASS: TestSSH (0.00s)
PASS
coverage: 0.0% of statements
ok      lemon_ssh/mssh  0.426s  coverage: 0.0% of statements


##  生成html格式的覆盖率报告
$ go tool cover -html=cover.out -o coverage.html
$ ll coverage.html





运维工具(lemon):批量执行linux命令及shell脚本

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