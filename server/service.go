package server

import (
	"fmt"
	"log"
	"strconv"

	"github.com/google/tcpproxy"
	cfg "github.com/lukluk/kube-local-proxy/config"
)

const httpPort = "80"
const tcpPort = "443"
const localhost = "127.0.0.1"

type Server struct {
	Proxy        tcpproxy.Proxy
	Konfigs      []cfg.Konfig
	startingPort int
}

func NewServer(startingPort int, konfigs []cfg.Konfig) *Server {
	return &Server{
		tcpproxy.Proxy{},
		konfigs,
		startingPort,
	}
}
func (s *Server) Start() { //Http start
	localPort := s.startingPort
	for _, konfig := range s.Konfigs {
		if konfig.GrpcService != "" {
			s.registerTcp(konfig.VirtualHost, localPort)
		} else {
			fmt.Println("http")
			s.registerHttp(konfig.VirtualHost, localPort)
		}
		localPort++
	}
	log.Fatal(s.Proxy.Run())
}

func (s *Server) registerHttp(vhost string, localPort int) {
	fmt.Println(":"+httpPort, vhost, localhost+":"+strconv.Itoa(localPort))
	s.Proxy.AddHTTPHostRoute(":"+httpPort, vhost, tcpproxy.To(localhost+":"+strconv.Itoa(localPort)))
}

func buildTCPTarget(startingPort int, konfigs []cfg.Konfig) tcpproxy.Target {
	return NewTCPServerRouter(
		startingPort,
		konfigs,
	)
}
func (s *Server) registerTcp(vhost string, localPort int) {
	s.Proxy.AddRoute(":"+tcpPort, buildTCPTarget(s.startingPort, s.Konfigs))
}
