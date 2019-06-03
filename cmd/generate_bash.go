package cmd

import (
	"strconv"

	cfg "github.com/lukluk/kube-local-proxy/config"
)

func Gen(startingPort int, konfigs []cfg.Konfig) string {
	inline := ""
	for _, konfig := range konfigs {
		inline = inline + "kubectx " + konfig.Context + "\n" +
			"pod=$(kubectl get pods --field-selector=status.phase=Running | grep " + konfig.ServiceName + " | tail -1 | awk '{print $1}')\n" +
			"kubectl port-forward pods/$pod " + strconv.Itoa(startingPort) + ":" + strconv.Itoa(konfig.ServicePort) + " & \n"
		startingPort++
	}
	return inline
}
