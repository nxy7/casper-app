package storage

type OnchainStorage interface {
	GetUserNfts() string
	GetAllNfts() string
}

type Storage interface {
}
