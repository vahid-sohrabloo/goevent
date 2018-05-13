package example

import "google.golang.org/grpc"

type GrpcEvent struct {
	StreamServerInterceptor []func(list *[]grpc.StreamServerInterceptor) `eventName:"GetStreamServerInterceptor"`
	UnaryServerInterceptor  []func(list *[]grpc.UnaryServerInterceptor)  `eventName:"GetUnaryInterceptor"`
	Init                    []func(grpcServer *grpc.Server)              `eventName:"Init"`
	Start                   []func(grpcServer grpc.Server)               `eventName:"Start"`
}
