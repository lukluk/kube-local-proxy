package cmd

import (
	"testing"

	cfg "github.com/lukluk/kube-local-proxy/config"
)

func TestGen(t *testing.T) {
	type args struct {
		startingPort int
		konfigs      []cfg.Konfig
	}

	konfigs := make([]cfg.Konfig, 1)
	konfigs[0] = cfg.Konfig{
		"context",
		"vhost",
		"service",
		8080,
		"com.gopay.withdrawal",
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"TestShouldGenerateCommand",
			args{
				2000,
				konfigs,
			},
			"echo context/service&& kubectx context\npod=$(kubectl get pods --field-selector=status.phase=Running | grep service | tail -1 | awk '{print $1}')\nkubectl port-forward pods/$pod 2000:8080 > /tmp/context.service.log & sleep 2; ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Gen(tt.args.startingPort, tt.args.konfigs); got != tt.want {
				t.Errorf("Gen() = %v, want %v", got, tt.want)
			}
		})
	}
}
