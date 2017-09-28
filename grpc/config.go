/*

LICENSE:  MIT
Author:   sine
Email:    sinerwr@gmail.com

*/

package rpc

import (
	"github.com/getsentry/raven-go"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/SiCo-Ops/Pb"
)

// B ConfigService PushRPC
func ConfigPushRPC(cc *grpc.ClientConn, in *pb.ConfigPushCall) *pb.ConfigPushBack {
	defer cc.Close()
	c := pb.NewConfigServiceClient(cc)
	r, err := c.PushRPC(context.Background(), in)
	if err != nil {
		raven.CaptureError(err, nil)
		return &pb.ConfigPushBack{Code: 304}
	}
	return r
}

// B ConfigService PullRPC
func ConfigPullRPC(cc *grpc.ClientConn, in *pb.ConfigPullCall) *pb.ConfigPullBack {
	defer cc.Close()
	c := pb.NewConfigServiceClient(cc)
	r, err := c.PullRPC(context.Background(), in)
	if err != nil {
		raven.CaptureError(err, nil)
		return &pb.ConfigPullBack{Code: 304}
	}
	return r
}
