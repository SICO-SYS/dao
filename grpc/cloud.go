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

// CloudAPIservice RequestRPC
func CloudAPIRequestRPC(cc *grpc.ClientConn, in *pb.CloudAPICall) *pb.CloudAPIBack {
	defer cc.Close()
	c := pb.NewCloudAPIServiceClient(cc)
	r, err := c.RequestRPC(context.Background(), in)
	if err != nil {
		raven.CaptureError(err, nil)
		return &pb.CloudAPIBack{Code: 302}
	}
	return r
}

// CloudTokenService SetRPC
func CloudTokenSetRPC(cc *grpc.ClientConn, in *pb.CloudTokenCall) *pb.CloudTokenBack {
	defer cc.Close()
	c := pb.NewCloudTokenServiceClient(cc)
	r, err := c.SetRPC(context.Background(), in)
	if err != nil {
		raven.CaptureError(err, nil)
		return &pb.CloudTokenBack{Code: 302}
	}
	return r
}

// CloudTokenService GetRPC
func CloudTokenGetRPC(cc *grpc.ClientConn, in *pb.CloudTokenCall) *pb.CloudTokenBack {
	defer cc.Close()
	c := pb.NewCloudTokenServiceClient(cc)
	r, err := c.GetRPC(context.Background(), in)
	if err != nil {
		raven.CaptureError(err, nil)
		return &pb.CloudTokenBack{Code: 302}
	}
	return r
}
