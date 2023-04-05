package datebase

import "database/sql"

type DBConn struct {
	Conn sql.Conn
}

func Connect(datasource string) *DBConn {
	return &DBConn{
		Conn: sql.
	}
}
