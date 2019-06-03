package server

import (
	"fmt"
	"io"
	"net"
	"strconv"
	"strings"

	cfg "github.com/lukluk/kube-local-proxy/config"
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
	var bufs [][]byte
	var nrs []int
	method := ""
	for {
		buf := buildBuffer(conn)
		nr, _ := conn.Read(buf)
		nrs = append(nrs, nr)
		bufs = append(bufs, buf)
		method = parseMethod(buf)
		method = strings.Replace(method, "/", "", -1)
		fmt.Println(method)
		if method != "" {
			break
		}
	}

	if method != "" {
		startingPort := r.StartingPort
		for _, konfig := range r.Konfigs {
			if konfig.GrpcService != "" && method == konfig.GrpcService {
				dialProxy := DialProxy{
					Addr: localhost + ":" + strconv.Itoa(startingPort),
				}
				dialProxy.HandleConn(conn, bufs, nrs)
				return
			}
			startingPort++
		}
		fallbackResponse(conn)
		return
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

func buildBuffer(conn io.Reader) []byte {
	var buf []byte
	if buf == nil {
		size := 32 * 1024
		if l, ok := conn.(*io.LimitedReader); ok && int64(size) > l.N {
			if l.N < 1 {
				size = 1
			} else {
				size = int(l.N)
			}
		}
		buf = make([]byte, size)
	}
	return buf
}

func fallbackResponse(w io.Writer) {
	io.WriteString(w, "PROXY UNKNOWN\r\n")
}
