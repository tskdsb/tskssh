package types

import (
	"fmt"
	"os"

	"github.com/tskdsb/tskssh/util"
	"golang.org/x/crypto/ssh"
)

type SSHHost struct {
	IP       string
	Port     string
	User     string
	Password string
	Client   *ssh.Client
}

func (h *SSHHost) Dial() error {
	config := &ssh.ClientConfig{
		User: h.User,
		Auth: []ssh.AuthMethod{
			ssh.Password(h.Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		// Timeout:         0,
	}

	client, err := ssh.Dial("tcp", h.IP+":"+h.Port, config)
	if err != nil {
		return err
	}

	h.Client = client
	return nil
}

func (h *SSHHost) RunCommand(cmd string) error {
	session, err := h.Client.NewSession()
	if err != nil {
		return err
	}

	session.Stdin = os.Stdin
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr

	return session.Run(cmd)
}

func (h *SSHHost) ReceiveFile(localFileName string, remotePath string) error {
	session, err := h.Client.NewSession()
	if err != nil {
		return err
	}

	localFile, err := os.Open(localFileName)
	if err != nil {
		return err
	}
	defer func() {
		util.LogErr(localFile.Close())
	}()

	session.Stdin = localFile

	return session.Run(fmt.Sprintf("cat >%s", remotePath))
}

func (h *SSHHost) SendFile(localFileName string, remotePath string) error {
	session, err := h.Client.NewSession()
	if err != nil {
		return err
	}

	localFile, err := os.Create(localFileName)
	if err != nil {
		return err
	}
	defer func() {
		util.LogErr(localFile.Close())
	}()

	session.Stdout = localFile

	return session.Run(fmt.Sprintf("cat %s", remotePath))
}
