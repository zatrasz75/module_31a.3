package postgres

import (
	"GoNews/pkg/storage"
	Interface "GoNews/pkg/storage"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		ctx    context.Context
		constr string
	}
	var tests []struct {
		name    string
		args    args
		want    *Storage
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.ctx, tt.args.constr)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorage_AddPost(t *testing.T) {
	type fields struct {
		Db *pgxpool.Pool
	}
	type args struct {
		p storage.Post
	}
	var tests []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pg := &Storage{
				Db: tt.fields.Db,
			}
			if err := pg.AddPost(tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("AddPost() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStorage_DeletePost(t *testing.T) {
	type fields struct {
		Db *pgxpool.Pool
	}
	type args struct {
		p storage.Post
	}
	var tests []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pg := &Storage{
				Db: tt.fields.Db,
			}
			if err := pg.DeletePost(tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("DeletePost() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStorage_UpdatePost(t *testing.T) {
	type fields struct {
		Db *pgxpool.Pool
	}
	type args struct {
		p storage.Post
	}
	var tests []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pg := &Storage{
				Db: tt.fields.Db,
			}
			if err := pg.UpdatePost(tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("UpdatePost() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStorage_Posts(t *testing.T) {
	type fields struct {
		Db *pgxpool.Pool
	}
	var tests []struct {
		name    string
		fields  fields
		want    []Interface.Post
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pg := &Storage{
				Db: tt.fields.Db,
			}
			got, err := pg.Posts()
			if (err != nil) != tt.wantErr {
				t.Errorf("Posts() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Posts() got = %v, want %v", got, tt.want)
			}
		})
	}
}
