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

// TemplateService CreateRPC
func TemplateCreateRPC(cc *grpc.ClientConn, in *pb.AssetTemplateCall) *pb.AssetTemplateBack {
	defer cc.Close()
	c := pb.NewTemplateServiceClient(cc)
	r, err := c.CreateRPC(context.Background(), in)
	if err != nil {
		raven.CaptureError(err, nil)
		return &pb.AssetTemplateBack{Code: 303}
	}
	return r
}

// AssetService SynchronizeRPC
func AssetSynchronizeRPC(cc *grpc.ClientConn, in *pb.AssetSynchronizeCall) *pb.AssetSynchronizeBack {
	defer cc.Close()
	c := pb.NewAssetServiceClient(cc)
	r, err := c.SynchronizeRPC(context.Background(), in)
	if err != nil {
		raven.CaptureError(err, nil)
		return &pb.AssetSynchronizeBack{Code: 303}
	}
	return r
}

// AssetService CustomRPC
func AssetCustomRPC(cc *grpc.ClientConn, in *pb.AssetCustomizeCall) *pb.AssetCustomizeBack {
	defer cc.Close()
	c := pb.NewAssetServiceClient(cc)
	r, err := c.CustomRPC(context.Background(), in)
	if err != nil {
		raven.CaptureError(err, nil)
		return &pb.AssetCustomizeBack{Code: 303}
	}
	return r
}
