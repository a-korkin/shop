package ports

import (
	pb "github.com/a-korkin/shop/internal/common"
)

type DbConnect interface {
	GetItem(id int32) (*pb.Item, error)
	CreateItem(in *pb.ItemDto) (*pb.Item, error)
	GetItems(in *pb.PageParams) (*pb.ItemList, error)
}
