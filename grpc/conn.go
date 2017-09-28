/*

LICENSE:  MIT
Author:   sine
Email:    sinerwr@gmail.com

*/

package rpc

import (
	"github.com/getsentry/raven-go"
	"google.golang.org/grpc"
)

// Deprecated After 1.0.0
func RPCConn(address string) *grpc.ClientConn {
	defer func() {
		recover()
	}()
	conn, err := grpc.Dial(address)
	if err != nil {
		raven.CaptureError(err, nil)
	}

	return conn
}

func Conn(host, port string) (*grpc.ClientConn, error) {
	if port == "" {
		port = "6666"
	}
	address := host + ":" + port
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	return conn, err
}
