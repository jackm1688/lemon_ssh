package util

import (
	"bufio"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"lemon_ssh/model"
	"os"
	"time"
)

var(
	RecordFile *os.File
	OpenErr error
)


const prefix = "Lemon3E05915819D43882AA14BF3057527D199463B2E1E416897DC79F49870C2035C0345B89B45538B40199744E4DC501E622EC834175C91B22D679DF25850A903F10719B44227558C30C2F006EDA4297A75413395A13B0267D5B165099017BBC3626803010091EB5E12AED0FB501D21D382412268F625DF988867B50361E685121B5E007A9234100F116CEB9670EC081724638E8817F8DE1382E344421058CC79E9F72251395C0E99BBF17312561837C791E38DD3C2F226C037514E43C87C3D82C15E16C3F81C0A6C12EF8D756A00E9F40B3951E028041FA24D74CB0D2425814256398E1AB1A755FA6AA970E41F9E34803F0516E5877916108090D8F283806982DA1B89F29B51DAC7A1352CC55B9A8E666D2A82A47471F652B0118BA9B52FEE5E1361742B1D8836D503F723B36636BA"

func WriteToFile(file *os.File, context string, err error) {
	var record *model.LogRecord = &model.LogRecord{
		CreateAt: time.Now().Format("2006-01-02 15:04:05"),
		Content:  context,
	}

	if err == nil {
		data, _ := json.Marshal(record)
		md5Str := fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s,%s", prefix, string(data)))))
		record.Md5Value = md5Str
		data, _ = json.Marshal(record)
		writer := bufio.NewWriter(file)
		_,_ =fmt.Fprintln(writer, string(data))
		_ = writer.Flush()
	}
}

func ReadFromFile(file *os.File, err error) (recordList []*model.LogRecord) {
	if err == nil {
		reader := bufio.NewReader(file)
		for {
			if data, err := reader.ReadBytes('\n'); err != io.EOF{
				var record model.LogRecord
				if err := json.Unmarshal(data,&record); err == nil {
					recordList = append(recordList,&record)
				}
			}else {
				break
			}
		}
	}
	return
}

func CheckFile(filename string) bool  {
	_,err := os.Stat(filename)
	return os.IsExist(err)
}

func EquMD5(record *model.LogRecord) bool {

	srcMd5 := record.Md5Value
	record.Md5Value = ""
	data, _ := json.Marshal(record)
	dstMd5 := fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s,%s", prefix, string(data)))))

	if srcMd5 == dstMd5 {
		return true
	}
	return false
}
