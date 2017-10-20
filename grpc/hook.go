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

func HookAuthRPC(cc *grpc.ClientConn, in *pb.HookAuthCall) *pb.HookAuthBack {
	defer cc.Close()
	c := pb.NewHookServiceClient(cc)
	r, err := c.AuthRPC(context.Background(), in)
	if err != nil {
		raven.CaptureError(err, nil)
		return &pb.HookAuthBack{Code: 305}
	}
	return r
}

// HookService CreateRPC
func HookCreateRPC(cc *grpc.ClientConn, in *pb.HookCreateCall) *pb.HookCreateBack {
	defer cc.Close()
	c := pb.NewHookServiceClient(cc)
	r, err := c.CreateRPC(context.Background(), in)
	if err != nil {
		raven.CaptureError(err, nil)
		return &pb.HookCreateBack{Code: 305}
	}
	return r
}

//HookService UpdateNameRPC
func HookUpdateNameRPC(cc *grpc.ClientConn, in *pb.HookUpdateNameCall) *pb.HookUpdateNameBack {
	defer cc.Close()
	c := pb.NewHookServiceClient(cc)
	r, err := c.UpdateNameRPC(context.Background(), in)
	if err != nil {
		raven.CaptureError(err, nil)
		return &pb.HookUpdateNameBack{Code: 305}
	}
	return r
}

// HookService QueryRPC
func HookQueryRPC(cc *grpc.ClientConn, in *pb.HookQueryCall) *pb.HookQueryBack {
	defer cc.Close()
	c := pb.NewHookServiceClient(cc)
	r, err := c.QueryRPC(context.Background(), in)
	if err != nil {
		raven.CaptureError(err, nil)
		return &pb.HookQueryBack{Code: 305}
	}
	return r
}

// HookService ReceiveRPC
func HookReceiveRPC(cc *grpc.ClientConn, in *pb.HookReceiveCall) *pb.HookReceiveBack {
	defer cc.Close()
	c := pb.NewHookServiceClient(cc)
	r, err := c.ReceiveRPC(context.Background(), in)
	if err != nil {
		raven.CaptureError(err, nil)
		return &pb.HookReceiveBack{Code: 305}
	}
	return r
}
