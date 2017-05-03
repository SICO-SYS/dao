package dao

import (
	"google.golang.org/grpc"

	. "github.com/SiCo-DevOps/log"
)

func RpcConn(address string) *grpc.ClientConn {
	defer func() {
		recover()
		if recover() != nil {
			LogErrMsg(50, "dao.RpcConn")
		}
	}()
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		LogErrMsg(5, "dao.RpcConn."+address)
	}

	return conn
}
