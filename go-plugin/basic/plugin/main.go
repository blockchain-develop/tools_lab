package main

import (
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/go-plugin"
	"github.com/tools_lab/go-plugin/basic/common"
	"os"
)

type GreeterHello struct {
	logger hclog.Logger
}

func (g *GreeterHello) Greet() string {
	g.logger.Debug("message from GreeterHello.Greet")
	return "Hello!"
}


var handshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "BASIC_PLUGIN",
	MagicCookieValue: "hello",
}

func main() {
	logger := hclog.New(&hclog.LoggerOptions{
		Name:                     "greet",
		Level:                    hclog.Trace,
		Output:                   os.Stderr,
		JSONFormat:               true,
	})
	greeter := &GreeterHello{
		logger: logger,
	}
	var pluginMap = map[string]plugin.Plugin {
		"greeter": &common.GreeterPlugin{Impl: greeter},
	}
	logger.Debug("message from plugin", "foo", "bar")
	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig:  handshakeConfig,
		Plugins: pluginMap,
	})
}
