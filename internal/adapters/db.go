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
		"select id, title, price from public.items where id = $1", id)
	if err != nil {
		return nil, err
	}
	if rows.Next() {
		item := pb.Item{}
		if err := rows.Scan(&item.Id, &item.Title, &item.Price); err != nil {
			return nil, err
		}
		return &item, nil
	}
	return nil, nil
}
func (dbConn *DbConnect) CreateItem(in *pb.ItemDto) (*pb.Item, error) {
	rows, err := dbConn.Db.Query(
		"insert into public.items(title, price) values($1, $2) returning id",
		in.Title, in.Price)
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
