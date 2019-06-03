package config

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Konfig struct {
	Context     string
	VirtualHost string
	ServiceName string
	ServicePort int
	GrpcService string
}

func GetConfig(path string) []Konfig {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var konfig []Konfig
	for scanner.Scan() {
		konfig = append(konfig, Explode(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return konfig
}

func Explode(str string) Konfig {
	kv := strings.Split(str, "=")
	if len(kv) < 2 {
		panic("Invalid key config")
	}
	subdomain := strings.TrimSpace(kv[0])
	values := strings.Split(kv[1], "/")
	if len(values) < 2 {
		panic("Invalid value config")
	}
	kubeContext := strings.TrimSpace(values[0])
	serviceValue := strings.Split(values[1], ":")
	serviceName := strings.TrimSpace(serviceValue[0])
	servicePort, _ := strconv.Atoi(serviceValue[1])
	grpcService := ""
	if len(serviceValue) == 3 {
		grpcService = strings.TrimSpace(serviceValue[2])
	}

	return Konfig{
		kubeContext,
		subdomain,
		serviceName,
		servicePort,
		grpcService,
	}
}
