package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/lukluk/kube-virtualhost/cmd"
	"github.com/lukluk/kube-virtualhost/config"
	"github.com/lukluk/kube-virtualhost/server"
)

const defaultConfigPath = ".kvhost.cfg"

func main() {
	if len(os.Args) < 2 {
		fmt.Println("kvh ready")
		return
	}
	arg := os.Args[1]
	usr, _ := user.Current()
	configPath := usr.HomeDir + "/" + defaultConfigPath
	konfigs := config.GetConfig(configPath)
	startingPort := 3011

	if arg == "start" {
		fmt.Println(konfigs)
		s := server.NewServer(startingPort, konfigs)
		s.Start()
	} else if arg == "gen" {
		fmt.Println(cmd.Gen(startingPort, konfigs))
	}
}
