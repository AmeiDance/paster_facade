// Code generated by Kitex v0.0.4. DO NOT EDIT.
package facade

import (
	"github.com/ameidance/paster_facade/model/vo/kitex_gen/paster/facade"
	"github.com/cloudwego/kitex/server"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler facade.Facade, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)

	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}