package ssh

import (
	"log"
	"os"

	"golang.org/x/crypto/ssh"
)

type CmdJob struct {
	IP       string
	Port     string
	User     string
	Password string
	Command  string
	WorkDir  string

	Client *ssh.Client
}

func (cj *CmdJob) Dial() error {

	if cj.Client != nil {
		return nil
	}

	config := &ssh.ClientConfig{
		User: cj.User,
		Auth: []ssh.AuthMethod{
			ssh.Password(cj.Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         0,
	}

	client, err := ssh.Dial("tcp", cj.IP+":"+cj.Port, config)
	if err != nil {
		return err
	}

	cj.Client = client
	return nil
}

func (cj *CmdJob) Run() error {

	err := cj.Dial()
	if err != nil {
		return err
	}

	session, err := cj.Client.NewSession()
	if err != nil {
		return err
	}
	defer session.Close()

	session.Stdin = os.Stdin
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr

	// if cmd == "" && sh != "" {
	// 	scriptFile := sh[strings.LastIndex(sh, "/")+1:]
	//
	// 	dir2 := dir
	// 	if !strings.HasSuffix(dir, "/") {
	// 		dir2 = dir + "/"
	// 	}
	//
	// 	from = sh
	// 	to = user + ":" + password + ":" + ip + ":" + dir2 + scriptFile
	// 	cmd = "source " + dir2 + scriptFile
	// 	sshCp()
	// }

	err = session.Run(cj.Command)
	if err != nil {
		if waitMsg, ok := err.(*ssh.ExitError); ok {
			os.Exit(waitMsg.ExitStatus())
		} else {
			log.Print("Run cmd failed: ", err)
			os.Exit(-1)
		}
	}

	return nil
}

func (cj *CmdJob) Clean() {
	if cj.Client != nil {
		_ = cj.Client.Close()
	}
}
