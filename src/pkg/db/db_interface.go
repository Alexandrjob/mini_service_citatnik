package db

type DataBase interface {
	Conn() (interface{}, error)
	Close() error
}
