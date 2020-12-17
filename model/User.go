package model

import "time"

type User struct {
	User string  `json:"user"`
	Password string  `json:"password"`
	Host string  `json:"host"`
	Key string `json:"key"`
	Port int `json:"port"`
	CipherList []string `json:"cipherList"`
	Timeout time.Duration `json:"timeout"`
	IsTrustFile bool `json:"isTrustFile"`
	Cmd string `json:"cmd"`
}
