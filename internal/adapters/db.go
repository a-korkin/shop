package adapters

import (
	"database/sql"
	"io"

	pb "github.com/a-korkin/shop/internal/common"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
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

func (dbConn *DbConnect) CreateUser(in *pb.UserDto) (*pb.User, error) {
	rows, err := dbConn.Db.Query(
		`
insert into public.users(last_name, first_name) 
values ($1, $2) returning id, last_name, first_name`, in.LastName, in.FirstName)
	if err != nil {
		return nil, err
	}
	user := pb.User{}
	if rows.Next() {
		if err := rows.Scan(&user.Id, &user.LastName, &user.FirstName); err != nil {
			return nil, err
		}
	}
	return &user, nil
}

func (dbConn *DbConnect) UpdUser(in *pb.User) (*pb.User, error) {
	rows, err := dbConn.Db.Query(`
update public.users
set last_name = $2,
	first_name = $3
where id = $1
returning last_name, first_name`, in.Id, in.LastName, in.FirstName)
	if err != nil {
		return nil, err
	}
	if rows.Next() {
		if err := rows.Scan(&in.LastName, &in.FirstName); err != nil {
			return nil, err
		}
	}
	return in, nil
}

func (dbConn *DbConnect) GetUser(in *pb.UserId) (*pb.User, error) {
	rows, err := dbConn.Db.Query(`
select last_name, first_name
from public.users
where id = $1`, in.Id)
	if err != nil {
		return nil, err
	}
	user := pb.User{}
	if rows.Next() {
		if err = rows.Scan(&user.LastName, &user.FirstName); err != nil {
			return nil, err
		}
	}
	return &user, nil
}

func (dbConn *DbConnect) GetUsers(params *pb.PageParams, stream grpc.ServerStreamingServer[pb.User]) error {
	rows, err := dbConn.Db.Query(`
select id, last_name, first_name
from public.users
offset $1
limit $2`, params.Offset, params.Limit)
	if err != nil {
		return err
	}
	for rows.Next() {
		user := pb.User{}
		if err = rows.Scan(&user.Id, &user.LastName, &user.FirstName); err != nil {
			return err
		}
		err = stream.Send(&user)
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
	}
	return nil
}

func (dbConn *DbConnect) DropUser(in *pb.UserId) (*pb.Empty, error) {
	_, err := dbConn.Db.Exec("delete from public.users where id = $1", in.Id)
	if err != nil {
		return nil, err
	}
	return &pb.Empty{}, nil
}

func (dbConn *DbConnect) Buy(in *pb.PurchaseDto) (*pb.Purchase, error) {
	rows, err := dbConn.Db.Query(`
insert into public.purchases(user_id, item_id, count_items)
value($1, $2, $3)
returning id, user_id, item_id, time_of_purchase, count_items
`, in.UserId, in.ItemId, in.CountItems)
	if err != nil {
		return nil, err
	}
	out := pb.Purchase{}
	if rows.Next() {
		err = rows.Scan(&out.Id, &out.UserId,
			&out.ItemId, &out.TimeOfPurchase, &out.CountItems)
		if err != nil {
			return nil, err
		}
	}
	return &out, nil
}
