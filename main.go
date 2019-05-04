package main

import (
	"flag"
	"log"

	"github.com/tskdsb/tskssh/ssh"
)

func main() {
	var a ssh.CmdJob
	flag.StringVar(&a.IP, "ip", "", "host ip")
	flag.StringVar(&a.Port, "port", "22", "host port")
	flag.StringVar(&a.User, "user", "root", "user name")
	flag.StringVar(&a.Password, "password", "", "password for user")
	flag.StringVar(&a.Command, "cmd", "", "command to run")
	flag.StringVar(&a.WorkDir, "workDir", "", "which dir to run command")

	flag.Parse()
	err := a.Run()
	if err != nil {
		log.Printf("%s\n", err)
	}
	a.Clean()
}
