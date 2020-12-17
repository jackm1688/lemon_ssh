package util

import(
	"fmt"
	"os"
	"testing"
)

func TestWriteToFile(t *testing.T) {

	file,_ := os.OpenFile("./record.db",os.O_CREATE|os.O_RDWR|os.O_APPEND,0660)
	/*record := &model.LogRecord{
		Id:       1,
		CreateAt: "2020-11-13 20:38:02",//time.Now().Format("2006-01-02 15:05:04"),
		Content:  "kubctl get pod",
	}
	WriteToFile(file,record,nil)*/

	list := ReadFromFile(file,nil)
	for _,v  := range list{
		fmt.Printf("%#v\n",v)
	}
}
