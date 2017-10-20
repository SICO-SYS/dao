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

// OrchestrationService CreateRPC
func OrchestrationCreateRPC(cc *grpc.ClientConn, in *pb.OrchestrationCreateCall) *pb.OrchestrationCreateBack {
	defer cc.Close()
	c := pb.NewOrchestrationServiceClient(cc)
	r, err := c.CreateRPC(context.Background(), in)
	if err != nil {
		raven.CaptureError(err, nil)
		return &pb.OrchestrationCreateBack{Code: 306}
	}
	return r
}

// OrchestrationService CheckRPC
func OrchestrationCheckRPC(cc *grpc.ClientConn, in *pb.OrchestrationCheckCall) *pb.OrchestrationCheckBack {
	defer cc.Close()
	c := pb.NewOrchestrationServiceClient(cc)
	r, err := c.CheckRPC(context.Background(), in)
	if err != nil {
		raven.CaptureError(err, nil)
		return &pb.OrchestrationCheckBack{Code: 306}
	}
	return r
}
