package mock

import (
	"github.com/go-iot-platform/go-micro/registry"
	"github.com/go-iot-platform/go-micro/server"
)

type MockHandler struct {
	Id   string
	Opts server.HandlerOptions
	Hdlr interface{}
}

func (m *MockHandler) Name() string {
	return m.Id
}

func (m *MockHandler) Handler() interface{} {
	return m.Hdlr
}

func (m *MockHandler) Endpoints() []*registry.Endpoint {
	return []*registry.Endpoint{}
}

func (m *MockHandler) Options() server.HandlerOptions {
	return m.Opts
}