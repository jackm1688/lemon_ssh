package mssh

import (
	"errors"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"lemon_ssh/model"
)

type MSFTP struct {
	User *model.User
}

func (this *MSFTP) Connect(sshClient  *ssh.Client) (sftpClient *sftp.Client, err error) {

	if sshClient == nil {
		return nil,errors.New("ssh client is nil pointer")
	}
	//create sftpClient
	if sftpClient, err = sftp.NewClient(sshClient); err != nil {
		return nil, err
	}
	return sftpClient, nil
}
