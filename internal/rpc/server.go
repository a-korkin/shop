package rpc

import (
	"context"
	pb "github.com/a-korkin/shop/internal/common"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type ShopServer struct {
	pb.UnimplementedShopServiceServer
}

func NewShopServer() *ShopServer {
	return &ShopServer{}
}

func (s *ShopServer) GetItem(ctx context.Context, in *pb.ItemId) (*pb.Item, error) {
	item := pb.Item{
		Id:    1,
		Title: "Item #1",
		Price: 779.33,
	}
	return &item, nil
}

func (srv *ShopServer) Run(port string) {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to create listener: %s", err)
	}
	s := grpc.NewServer()
	pb.RegisterShopServiceServer(s, srv)
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to run grpc server: %s", err)
	}
}
