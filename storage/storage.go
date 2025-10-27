package storage

type Storable interface {
	Save() error
	Load() error
}