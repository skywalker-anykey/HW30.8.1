package postgres

import (
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"testing"
)

const (
	dbUser = "sandbox"
	dbPass = "sandbox"
	dbHost = "localhost"
	dbPort = "5432"
	dbName = "tasks"
)

func TestNew(t *testing.T) {
	type args struct {
		connect string
	}
	tests := []struct {
		name    string
		args    args
		want    *Storage
		wantErr bool
	}{
		{
			name: "DB Connect",
			args: args{
				connect: fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPass, dbHost, dbPort, dbName),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := New(tt.args.connect)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			return
		})
	}
}

func TestStorage_NewTask(t *testing.T) {
	type fields struct {
		db *pgxpool.Pool
	}
	type args struct {
		t Task
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		// Add test cases.
		{
			name: "NewTask",
			args: args{
				t: Task{
					Title:   "Test title",
					Content: "Test content",
				},
			},
			wantErr: false,
		},
	}
	s, _ := New(fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPass, dbHost, dbPort, dbName))
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := s.NewTask(tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func TestStorage_Tasks(t *testing.T) {
	type fields struct {
		db *pgxpool.Pool
	}
	type args struct {
		taskID   int
		authorID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []Task
		wantErr bool
	}{
		// Add test cases.
		{
			name: "Read All Tasks",
			args: args{
				taskID:   0,
				authorID: 0,
			},
			wantErr: false,
		},
		{
			name: "Read Task id1",
			args: args{
				taskID:   1,
				authorID: 0,
			},
			wantErr: false,
		},
	}
	s, _ := New(fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPass, dbHost, dbPort, dbName))
	for _, tt := range tests {
		fmt.Println("---", tt)
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.Tasks(tt.args.taskID, tt.args.authorID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Tasks() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(got)
		})
	}
}

func TestStorage_DeleteTask(t *testing.T) {
	type fields struct {
		db *pgxpool.Pool
	}
	type args struct {
		id int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// Add test cases.
		{
			name: "Delete id3",
			fields: fields{
				db: &pgxpool.Pool{},
			},
			args: args{
				id: 3,
			},
			wantErr: false,
		},
	}
	s, _ := New(fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPass, dbHost, dbPort, dbName))
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := s.DeleteTask(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStorage_TasksByLabel(t *testing.T) {
	type fields struct {
		db *pgxpool.Pool
	}
	type args struct {
		label string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []Task
		wantErr bool
	}{
		// Add test cases.
		{
			name: "Read By Label",
			args: args{
				label: "new",
			},
			wantErr: false,
		},
	}
	s, _ := New(fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPass, dbHost, dbPort, dbName))
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := s.TasksByLabel(tt.args.label)
			if (err != nil) != tt.wantErr {
				t.Errorf("TasksByLabel() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(got)
		})
	}
}

func TestStorage_UpdateTask(t *testing.T) {
	type fields struct {
		db *pgxpool.Pool
	}
	type args struct {
		t Task
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// Add test cases.
		{
			name: "Update id1",
			args: args{
				t: Task{
					ID:      1,
					Title:   "NewTitle",
					Content: "NewContent",
				},
			},
			wantErr: false,
		},
	}
	s, _ := New(fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPass, dbHost, dbPort, dbName))
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := s.UpdateTask(tt.args.t); (err != nil) != tt.wantErr {
				t.Errorf("UpdateTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
