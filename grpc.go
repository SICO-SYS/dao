package dao

import (
	"google.golang.org/grpc"

	. "github.com/SiCo-DevOps/log"
)

func RpcConn(bsns string) *grpc.ClientConn {
	address := bsns + ":6666"
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		LogErrMsg(5, "dao.RpcConn."+bsns)
	}

	return conn
}
