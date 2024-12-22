package rpc

import (
	"context"
	"log"
	"net"

	pb "github.com/a-korkin/shop/internal/common"
	"github.com/a-korkin/shop/internal/core"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type ShopServer struct {
	pb.UnimplementedShopServiceServer
	AppState *core.AppState
}

func NewShopServer(appState *core.AppState) *ShopServer {
	return &ShopServer{
		AppState: appState,
	}
}

func (s *ShopServer) GetItem(ctx context.Context, in *pb.ItemId) (*pb.Item, error) {
	return s.AppState.DbConn.GetItem(in.Id)
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
