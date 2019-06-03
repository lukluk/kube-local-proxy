package server

import (
	"fmt"
	"net"
	"strconv"
	"strings"

	cfg "github.com/lukluk/kube-virtualhost/config"
)

type TCPServerRouter struct {
	Konfigs      []cfg.Konfig
	StartingPort int
}

func NewTCPServerRouter(startingPort int, konfigs []cfg.Konfig) *TCPServerRouter {
	return &TCPServerRouter{
		konfigs,
		startingPort,
	}
}
func (r *TCPServerRouter) HandleConn(conn net.Conn) {
	fmt.Println("Handling new connection...")
	// Close connection when this function ends

	buf := make([]byte, 200)
	var bufs [][]byte
	var nrs []int
	method := ""
	for {
		nr, _ := conn.Read(buf)
		nrs = append(nrs, nr)
		bufs = append(bufs, buf)
		method = parseMethod(buf)
		fmt.Println(method)
		if method != "" {
			break
		}
	}

	if method != "" {
		startingPort := r.StartingPort
		for _, konfig := range r.Konfigs {
			if konfig.GrpcService != "" && strings.Contains(method, konfig.GrpcService) {
				dialProxy := DialProxy{
					Addr: localhost + ":" + strconv.Itoa(startingPort),
				}
				dialProxy.HandleConn(conn, bufs, nrs)
				return
			}
			startingPort++
		}
	}

}
func parseMethod(b []byte) string {
	out := ""
	method := ""
	i := 0
	for c := range b {
		chr := string(b[c])

		if isLetter(b[c]) {
			out += chr
			if strings.Contains(out, "path+/") {
				method += chr
				if chr == "/" && i > 0 {
					return method
				}
				i++
			}
		}

	}
	return ""
}

func isLetter(r byte) bool {
	if r == '@' || r == '+' || r == '/' || r == ':' || r == '.' {
		return true
	}
	if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {
		return false
	}

	return true
}
