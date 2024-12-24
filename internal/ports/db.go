package ports

import (
	pb "github.com/a-korkin/shop/internal/common"
	"google.golang.org/grpc"
)

type DbConnect interface {
	GetItem(id int32) (*pb.Item, error)
	CreateItem(in *pb.ItemDto) (*pb.Item, error)
	GetItems(in *pb.PageParams) (*pb.ItemList, error)
	DropItem(in *pb.ItemId) (*pb.Empty, error)
	UpdItem(in *pb.Item) (*pb.Item, error)

	CreateUser(in *pb.UserDto) (*pb.User, error)
	UpdUser(in *pb.User) (*pb.User, error)
	GetUser(in *pb.UserId) (*pb.User, error)
	GetUsers(*pb.PageParams, grpc.ServerStreamingServer[pb.User]) error
	DropUser(in *pb.UserId) (*pb.Empty, error)
}
