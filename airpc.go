package airpc

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// NewServer new rpc server
func NewServer(host string, service interface{}) error {
	err := rpc.Register(service)
	if err != nil {
		return err
	}
	rpc.HandleHTTP()
	listener, err := net.Listen("tcp", host)
	if err != nil {
		return err
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept err: %v", err)
			continue
		}
		go jsonrpc.ServeConn(conn)
	}
}

// NewClient new rpc client
func NewClient(host string) (*rpc.Client, error) {
	conn, err := net.Dial("tcp", host)
	if err != nil {
		return nil, err
	}
	return jsonrpc.NewClient(conn), nil
}
