package db

type DB interface {
	ConnectDB() error
	DisConnectDB()
}
