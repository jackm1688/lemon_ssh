package mlemon

import (
	"bytes"
	"fmt"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"lemon_ssh/model"
	"lemon_ssh/mssh"
	"os"
	"path"
	"strings"
)

type MLemon struct {
	session *ssh.Session
	scp     *sftp.Client
	Err     error
}

func NewMLemon(user *model.User) *MLemon {
	ssh := &mssh.MSSH{User: user}
	sftp := &mssh.MSFTP{User: user}

	mLemon := &MLemon{}
	var err error
	session,sshClient, err := ssh.Connect()
	if err != nil {
		mLemon.Err = err
	}else {
		scp, err := sftp.Connect(sshClient)
		mLemon.Err = err
		mLemon.scp = scp
	}

	mLemon.session = session
	return mLemon
}

//执行单个命令
func (this *MLemon) Exec(cmd string) (result string, err error) {
	if this.Err != nil {
		return "", err
	}

	defer func() {
		if this.session != nil {
			this.session.Close()
		}
	}()

	var sessionBuf bytes.Buffer
	this.session.Stdout = &sessionBuf
	this.session.Run(cmd)
	return strings.TrimSpace(fmt.Sprint(this.session.Stdout)), nil
}

//上传文件
func (this *MLemon) Scp(localFile string, remoteDir string, bufSize int) (err error) {

	srcFile, err := os.Open(localFile)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	remoteFileName := path.Base(localFile)

	if this.scp == nil {
		err = fmt.Errorf("sftp.Client nil pointer dereference")
		return err
	}

	dstFile, err := this.scp.Create(path.Join(remoteDir, remoteFileName))
	if err != nil {
		return
	}
	defer dstFile.Close()
	buf := make([]byte, bufSize)
	for {
		n, _ := srcFile.Read(buf)
		if n == 0 {
			break
		}
		_,_= dstFile.Write(buf[0:n])
	}
	return nil
}
