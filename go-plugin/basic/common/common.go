package common

import (
	"github.com/hashicorp/go-plugin"
	"net/rpc"
)

type Greeter interface {
	Greet() string
}

type GreeterPlugin struct {
	Impl Greeter
}

func (p *GreeterPlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return &GreeterRPCServer{
		Impl: p.Impl,
	}, nil
}

func (p *GreeterPlugin) Client(b *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &GreeterRPCClient{
		client: c,
	}, nil
}



type GreeterRPCServer struct {
	Impl Greeter
}

func (s *GreeterRPCServer) Greet(args interface{}, resp *string) error {
	*resp = s.Impl.Greet()
	return nil
}


type GreeterRPCClient struct {
	client *rpc.Client
}

func (c *GreeterRPCClient) Greet() string {
	var resp string
	err := c.client.Call("Plugin.Greet", new(interface{}), &resp)
	if err != nil {
		panic(err)
	}
	return resp
}
