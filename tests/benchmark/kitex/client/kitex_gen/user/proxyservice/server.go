// Code generated by Kitex v0.7.1. DO NOT EDIT.
package proxyservice

import (
	server "github.com/cloudwego/kitex/server"
	user "github.com/kitex-contrib/codec-dubbo/tests/benchmark/kitex/client/kitex_gen/user"
)

// NewServer creates a server.Server with the given handler and options.
func NewServer(handler user.ProxyService, opts ...server.Option) server.Server {
	var options []server.Option

	options = append(options, opts...)
	svr := server.NewServer(options...)
	if err := svr.RegisterService(serviceInfo(), handler); err != nil {
		panic(err)
	}
	return svr
}