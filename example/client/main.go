package main

import (
	"fmt"

	"github.com/9d77v/iec104"
	"github.com/9d77v/iec104/example/client/config"
	"github.com/9d77v/iec104/example/client/worker"
)

func main() {
	address := fmt.Sprintf("%s:%d", "192.168.76.16", 2404)
	subAddress := ""
	if config.SubServerHost != "" && config.SubServerPort != 0 {
		subAddress = fmt.Sprintf("%s:%d", config.SubServerHost, config.ServerPort)
	}
	client := iec104.NewClient(address, config.Logger, subAddress)
	client.Run(worker.Task)
}
