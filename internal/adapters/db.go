package adapters

import (
	"database/sql"
	pb "github.com/a-korkin/shop/internal/common"
	_ "github.com/lib/pq"
)

type DbConnect struct {
	Db *sql.DB
}

func NewDBConnect(connectionStr string) (*DbConnect, error) {
	db, err := sql.Open("postgres", connectionStr)
	if err != nil {
		return nil, err
	}
	dbConnect := DbConnect{
		Db: db,
	}
	return &dbConnect, nil
}

func (dbConn *DbConnect) GetItem(id int32) (*pb.Item, error) {
	rows, err := dbConn.Db.Query(
		"select id, title, price, category from public.items where id = $1", id)
	if err != nil {
		return nil, err
	}
	if rows.Next() {
		item := pb.Item{}
		if err := rows.Scan(&item.Id, &item.Title, &item.Price, &item.Category); err != nil {
			return nil, err
		}
		return &item, nil
	}
	return nil, nil
}

func (dbConn *DbConnect) CreateItem(in *pb.ItemDto) (*pb.Item, error) {
	rows, err := dbConn.Db.Query(
		`
insert into public.items(title, price, category) 
values($1, $2, $3) 
returning id`, in.Title, in.Price, in.Category)
	if err != nil {
		return nil, err
	}
	if rows.Next() {
		item := pb.Item{
			Title: in.Title,
			Price: in.Price,
		}
		if err := rows.Scan(&item.Id); err != nil {
			return nil, err
		}
		return &item, nil
	}

	return nil, nil
}

func (dbConn *DbConnect) GetItems(in *pb.PageParams) (*pb.ItemList, error) {
	rows, err := dbConn.Db.Query(`
select id, title, price, category
from public.items
offset $1
limit $2`, in.Offset, in.Limit)
	if err != nil {
		return nil, err
	}
	list := make([]*pb.Item, 0)
	for rows.Next() {
		item := pb.Item{}
		err := rows.Scan(&item.Id, &item.Title, &item.Price, &item.Category)
		if err != nil {
			return nil, err
		}
		list = append(list, &item)
	}
	return &pb.ItemList{Items: list}, nil
}

func (dbConn *DbConnect) DropItem(in *pb.ItemId) (*pb.Empty, error) {
	_, err := dbConn.Db.Exec("delete from public.items where id = $1", in.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Empty{}, nil
}

func (dbConn *DbConnect) UpdItem(in *pb.Item) (*pb.Item, error) {
	rows, err := dbConn.Db.Query(`
update public.items
set title = $2,
	price = $3
	category = $4
where id = $1`, in.Id, in.Title, in.Price, in.Category)
	if err != nil {
		return nil, err
	}
	if rows.Next() {
		return in, nil
	}
	return nil, nil
}
