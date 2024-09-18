// Package memdb заглушка для тестов работы с приложения с БД
package memdb

import (
	"pgsql/pkg/storage/postgres"
)

type DB []postgres.Task

func (db DB) Tasks(int, int) ([]postgres.Task, error) {
	return db, nil
}

func (db DB) NewTask(t postgres.Task) (int, error) {
	return 0, nil
}

func (db DB) DeleteTask(id int) error {
	return nil
}

func (db DB) TasksByLabel(label string) ([]postgres.Task, error) {
	return db, nil
}

func (db DB) UpdateTask(t postgres.Task) error {
	return nil
}
