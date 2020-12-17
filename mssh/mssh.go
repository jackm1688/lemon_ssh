package mssh

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"lemon_ssh/logger"
	"lemon_ssh/model"
	"net"
	"time"
)

type MSSH struct {
	User *model.User
}


func (this *MSSH) Connect() (session *ssh.Session,client *ssh.Client,err error) {

	//get auth method
	auth := make([]ssh.AuthMethod, 0)
	if this.User.Key == "" {
		auth = append(auth, ssh.Password(this.User.Password))
	} else {
		pemBytes, err := ioutil.ReadFile(this.User.Key)
		if err != nil {
			logger.Error("Cannot open key file:%v\n", err)
			return nil,nil,err
		}
		var signer ssh.Signer
		if this.User.Password == "" {
			signer, err = ssh.ParsePrivateKey(pemBytes)
		} else {
			signer, err = ssh.ParsePrivateKeyWithPassphrase(pemBytes, []byte(this.User.Password))
		}

		if err != nil {
			return nil,nil,err
		}
		auth = append(auth, ssh.PublicKeys(signer))
	}

	var config  ssh.Config
	if len(this.User.CipherList) == 0 {
		config = ssh.Config{
			Ciphers: []string{"aes128-ctr", "aes192-ctr", "aes256-ctr", "aes128-gcm@openssh.com", "arcfour256", "arcfour128", "aes128-cbc", "3des-cbc", "aes192-cbc", "aes256-cbc"},
		}
	} else {
		config = ssh.Config{Ciphers: this.User.CipherList}
	}

	clientConfig := &ssh.ClientConfig{
		Config:  config,
		User:    this.User.User,
		Auth:    auth,
		Timeout: this.User.Timeout * time.Second,
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	}

	//connect to ssh
	addr := fmt.Sprintf("%s:%d", this.User.Host, this.User.Port)
	if client, err = ssh.Dial("tcp", addr, clientConfig); err != nil {
		return
	}

	//create session
	if session, err = client.NewSession(); err != nil {
		return
	}

	modes := ssh.TerminalModes{
		ssh.ECHO:          0,     // disable echoing
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}

	if err := session.RequestPty("xterm", 80, 40, modes); err != nil {
		return nil,nil,err
	}
	return
}