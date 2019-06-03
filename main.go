package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/lukluk/kube-local-proxy/cmd"
	"github.com/lukluk/kube-local-proxy/config"
	"github.com/lukluk/kube-local-proxy/server"
)

const defaultConfigPath = ".klp.cfg"

func main() {
	if len(os.Args) < 2 {
		return
	}
	arg := os.Args[1]
	usr, _ := user.Current()
	configPath := usr.HomeDir + "/" + defaultConfigPath
	konfigs := config.GetConfig(configPath)
	startingPort := 3011

	if arg == "start" {
		if konfigs == nil {
			fmt.Println("config not found, please write to $HOME/.klp.cfg")
			return
		}
		fmt.Println("kube-local-proxy ready!")
		s := server.NewServer(startingPort, konfigs)
		s.Start()
	} else if arg == "gen" {
		fmt.Println(cmd.Gen(startingPort, konfigs))
	}
}
