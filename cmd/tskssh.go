package main

import (
	"flag"
	"os"
	"path/filepath"

	"github.com/tskdsb/tskssh/types"
	"github.com/tskdsb/tskssh/util"
)

func main() {

	ip := flag.String("ip", "", "host ip")
	port := flag.String("port", "22", "host port")
	user := flag.String("user", "root", "user name")
	password := flag.String("password", "", "password for user")
	command := flag.String("cmd", "", "command to run")
	file := flag.String("file", "", "script file to transport")
	workDir := flag.String("dir", "", "path to save file")
	flag.Parse()

	var host = types.SSHHost{
		IP:       *ip,
		Port:     *port,
		User:     *user,
		Password: *password,
	}

	err := host.Dial()
	if err != nil {
		util.LogErr(err)
		os.Exit(1)
	}
	defer func() {
		util.LogErr(host.Client.Close())
	}()

	if *file != "" {
		err := host.ReceiveFile(*file, filepath.Join(*workDir, filepath.Base(*file)))
		if err != nil {
			util.LogErr(err)
			os.Exit(1)
		}
	}

	if *command != "" {
		err = host.RunCommand(*command)
		if err != nil {
			util.LogErr(err)
			os.Exit(1)
		}
	}

	// if err != nil {
	// 	if waitMsg, ok := err.(*ssh.ExitError); ok {
	// 		os.Exit(waitMsg.ExitStatus())
	// 	} else {
	// 		log.Print("run cmd failed: ", err)
	// 		os.Exit(-1)
	// 	}
	// }
}
