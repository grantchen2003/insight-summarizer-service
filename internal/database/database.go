package database

import "sync"

type FileComponentPayload struct {
	FileComponentId int
	Summary         string
}

type FileComponentSummary struct {
	Id              int
	FileComponentId int
	Summary         string
}

type Database interface {
	Connect() error
	Close() error
	BatchSaveFileComponentSummaries([]FileComponentPayload) ([]FileComponentSummary, error)
}

var (
	singletonInstance Database
	once              sync.Once
)

func GetSingletonInstance() Database {
	once.Do(func() {
		singletonInstance = &PostgreSql{}
	})

	return singletonInstance
}
