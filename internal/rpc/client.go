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
	// defer func() {
	// 	if err := conn.Close(); err != nil {
	// 		log.Fatalf("failed to close grpc client connection: %s", err)
	// 	}
	// }()
	return pb.NewShopServiceClient(conn), err
}
