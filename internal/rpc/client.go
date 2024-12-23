package rpc

import (
	pb "github.com/a-korkin/shop/internal/common"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewGrpcClient(addr string) (pb.ShopServiceClient, error) {
	conn, err := grpc.NewClient(
		addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	return pb.NewShopServiceClient(conn), err
}
