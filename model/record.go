package model

type LogRecord struct {
	Id int64 `json:"id"`
	CreateAt string `json:"createAt"`
	Content string `json:"content"`
	Md5Value string `json:"md5Value"`
}
