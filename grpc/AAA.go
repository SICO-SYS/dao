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

// AAATokenService GenerateRPC
func AAATokenGenerateRPC(cc *grpc.ClientConn, in *pb.AAAGenerateTokenCall) *pb.AAAGenerateTokenBack {
	defer cc.Close()
	c := pb.NewAAATokenServiceClient(cc)
	r, err := c.GenerateRPC(context.Background(), in)
	if err != nil {
		raven.CaptureError(err, nil)
		return &pb.AAAGenerateTokenBack{Code: 301}
	}
	return r
}

// AAATokenService AuthenticationRPC
func AAATokenAuthenticationRPC(cc *grpc.ClientConn, in *pb.AAATokenCall) *pb.AAATokenBack {
	defer cc.Close()
	c := pb.NewAAATokenServiceClient(cc)
	r, err := c.AuthenticationRPC(context.Background(), in)
	if err != nil {
		raven.CaptureError(err, nil)
		return &pb.AAATokenBack{Code: 301}
	}
	return r
}

// AAATokenService AuthorizationRPC
func AAATokenAuthorizationRPC(cc *grpc.ClientConn, in *pb.AAAServiceCall) *pb.AAAServiceBack {
	defer cc.Close()
	c := pb.NewAAATokenServiceClient(cc)
	r, err := c.AuthorizationRPC(context.Background(), in)
	if err != nil {
		raven.CaptureError(err, nil)
		return &pb.AAAServiceBack{Code: 301}
	}
	return r
}

// AAATokenService AuthorizationRPC
func AAATokenAccountingRPC(cc *grpc.ClientConn, in *pb.AAAEventCall) *pb.AAAEventBack {
	defer cc.Close()
	c := pb.NewAAATokenServiceClient(cc)
	r, err := c.AccountingRPC(context.Background(), in)
	if err != nil {
		raven.CaptureError(err, nil)
		return &pb.AAAEventBack{Code: 301}
	}
	return r
}
