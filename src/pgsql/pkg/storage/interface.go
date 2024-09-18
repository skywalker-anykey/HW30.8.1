package storage

import "pgsql/pkg/storage/postgres"

type Interface interface {
	Tasks(int, int) ([]postgres.Task, error)
	NewTask(postgres.Task) (int, error)
	DeleteTask(int) error
	TasksByLabel(string) ([]postgres.Task, error)
	UpdateTask(postgres.Task) error
}
