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

func RPCConn(address string) *grpc.ClientConn {
	defer func() {
		recover()
		if recover() != nil {
			raven.CaptureMessage("dao.grpc.RPCConn", nil)
		}
	}()
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		raven.CaptureError(err, nil)
	}

	return conn
}
