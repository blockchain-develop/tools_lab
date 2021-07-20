package main

import (
	"fmt"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"github.com/tools_lab/go-plugin/basic/common"
	"os"
	"os/exec"
)

var handshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "BASIC_PLUGIN",
	MagicCookieValue: "hello",
}

func main() {
	logger := hclog.New(&hclog.LoggerOptions{
		Name:                     "plugin",
		Level:                    hclog.Debug,
		Output:                   os.Stdout,
	})
	var pluginMap = map[string]plugin.Plugin{
		"greeter": &common.GreeterPlugin{},
	}
	client := plugin.NewClient(&plugin.ClientConfig{
		HandshakeConfig:  handshakeConfig,
		Plugins:          pluginMap,
		Cmd:              exec.Command("./plugin/greeter"),
		Logger:           logger,
	})
	defer client.Kill()
	//
	rpcClient, err := client.Client()
	if err != nil {
		panic(err)
	}
	raw, err := rpcClient.Dispense("greeter")
	if err != nil {
		panic(err)
	}
	greeter := raw.(common.Greeter)
	message := greeter.Greet()
	fmt.Println(message)
}
