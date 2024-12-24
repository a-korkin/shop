package rpc

import (
	"context"
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

func (s *ShopServer) CreateItem(ctx context.Context, in *pb.ItemDto) (*pb.Item, error) {
	return s.AppState.DbConn.CreateItem(in)
}

func (s *ShopServer) GetItems(ctx context.Context, in *pb.PageParams) (*pb.ItemList, error) {
	return s.AppState.DbConn.GetItems(in)
}

func (s *ShopServer) DropItem(ctx context.Context, in *pb.ItemId) (*pb.Empty, error) {
	return s.AppState.DbConn.DropItem(in)
}

func (s *ShopServer) UpdItem(ctx context.Context, in *pb.Item) (*pb.Item, error) {
	return s.AppState.DbConn.UpdItem(in)
}

func (s *ShopServer) CreateUser(ctx context.Context, in *pb.UserDto) (*pb.User, error) {
	return s.AppState.DbConn.CreateUser(in)
}

func (s *ShopServer) UpdUser(ctx context.Context, in *pb.User) (*pb.User, error) {
	return nil, nil
}

func (s *ShopServer) GetUser(ctx context.Context, in *pb.UserId) (*pb.User, error) {
	return nil, nil
}

func (s *ShopServer) GetUsers(*pb.PageParams, grpc.ServerStreamingServer[pb.User]) error {
	return nil
}

func (s *ShopServer) DropUser(ctx context.Context, in *pb.UserId) (*pb.Empty, error) {
	return nil, nil
}

func (srv *ShopServer) Run(port string) error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}
	s := grpc.NewServer()
	pb.RegisterShopServiceServer(s,
		srv)
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		return err
	}
	return nil
}
