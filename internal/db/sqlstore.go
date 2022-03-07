package db

type SQLStoreI interface {
	User() UserRepository
}
