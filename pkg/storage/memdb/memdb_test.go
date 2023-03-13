package memdb

import (
	"GoNews/pkg/storage"
	Interface "GoNews/pkg/storage"
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	var tests []struct {
		name string
		want *Store
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestStore_AddPost(t *testing.T) {
	type args struct {
		post storage.Post
	}
	var tests []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Store{}
			got, err := s.AddPost(tt.args.post)
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

func TestStore_AddPostuser(t *testing.T) {
	type args struct {
		in0 storage.Authors
	}
	var tests []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Store{}
			got, err := s.AddPostuser(tt.args.in0)
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

func TestStore_DeletePost(t *testing.T) {
	type args struct {
		post storage.Post
	}
	var tests []struct {
		name    string
		args    args
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Store{}
			if err := s.DeletePost(tt.args.post); (err != nil) != tt.wantErr {
				t.Errorf("DeletePost() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStore_DeletePostuser(t *testing.T) {
	type args struct {
		in0 storage.Authors
	}
	var tests []struct {
		name    string
		args    args
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Store{}
			if err := s.DeletePostuser(tt.args.in0); (err != nil) != tt.wantErr {
				t.Errorf("DeletePostuser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStore_Posts(t *testing.T) {
	var tests []struct {
		name    string
		want    []Interface.Post
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Store{}
			got, err := s.Posts()
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

func TestStore_Postsuser(t *testing.T) {
	var tests []struct {
		name    string
		want    []Interface.Authors
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Store{}
			got, err := s.Postsuser()
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

func TestStore_UpdatePost(t *testing.T) {
	type args struct {
		post storage.Post
	}
	var tests []struct {
		name    string
		args    args
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Store{}
			if err := s.UpdatePost(tt.args.post); (err != nil) != tt.wantErr {
				t.Errorf("UpdatePost() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestStore_UpdatePostuser(t *testing.T) {
	type args struct {
		in0 storage.Authors
	}
	var tests []struct {
		name    string
		args    args
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Store{}
			if err := s.UpdatePostuser(tt.args.in0); (err != nil) != tt.wantErr {
				t.Errorf("UpdatePostuser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
