package mongo

import (
	"GoNews/pkg/storage"
	Interface "GoNews/pkg/storage"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	type args struct {
		ctx    context.Context
		constr string
	}
	tests := []struct {
		name    string
		args    args
		want    *Storage
		wantErr bool
	}{
		// TODO: Add test cases.
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
		Db *mongo.Client
	}
	type args struct {
		p storage.Post
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mg := &Storage{
				Db: tt.fields.Db,
			}
			got, err := mg.AddPost(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddPost() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AddPost() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorage_AddPostuser(t *testing.T) {
	type fields struct {
		Db *mongo.Client
	}
	type args struct {
		p storage.Authors
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mg := &Storage{
				Db: tt.fields.Db,
			}
			got, err := mg.AddPostuser(tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddPostuser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AddPostuser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorage_DeletePost(t *testing.T) {
	type fields struct {
		Db *mongo.Client
	}
	type args struct {
		p storage.Post
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mg := &Storage{
				Db: tt.fields.Db,
			}
			if err := mg.DeletePost(tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("DeletePost() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStorage_DeletePostuser(t *testing.T) {
	type fields struct {
		Db *mongo.Client
	}
	type args struct {
		p storage.Authors
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mg := &Storage{
				Db: tt.fields.Db,
			}
			if err := mg.DeletePostuser(tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("DeletePostuser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStorage_Posts(t *testing.T) {
	type fields struct {
		Db *mongo.Client
	}
	tests := []struct {
		name    string
		fields  fields
		want    []Interface.Post
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mg := &Storage{
				Db: tt.fields.Db,
			}
			got, err := mg.Posts()
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

func TestStorage_Postsuser(t *testing.T) {
	type fields struct {
		Db *mongo.Client
	}
	tests := []struct {
		name    string
		fields  fields
		want    []Interface.Authors
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mg := &Storage{
				Db: tt.fields.Db,
			}
			got, err := mg.Postsuser()
			if (err != nil) != tt.wantErr {
				t.Errorf("Postsuser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Postsuser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStorage_UpdatePost(t *testing.T) {
	type fields struct {
		Db *mongo.Client
	}
	type args struct {
		p storage.Post
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mg := &Storage{
				Db: tt.fields.Db,
			}
			if err := mg.UpdatePost(tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("UpdatePost() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStorage_UpdatePostuser(t *testing.T) {
	type fields struct {
		Db *mongo.Client
	}
	type args struct {
		p storage.Authors
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mg := &Storage{
				Db: tt.fields.Db,
			}
			if err := mg.UpdatePostuser(tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("UpdatePostuser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
