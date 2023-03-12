package api

import (
	"GoNews/pkg/storage"
	"github.com/gorilla/mux"
	"net/http"
	"reflect"
	"testing"
)

func TestAPI_Router(t *testing.T) {
	type fields struct {
		db     storage.Interface
		router *mux.Router
	}
	var tests []struct {
		name   string
		fields fields
		want   *mux.Router
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &API{
				db:     tt.fields.db,
				router: tt.fields.router,
			}
			if got := api.Router(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Router() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAPI_addPostHandler(t *testing.T) {
	type fields struct {
		db     storage.Interface
		router *mux.Router
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	var tests []struct {
		name   string
		fields fields
		args   args
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &API{
				db:     tt.fields.db,
				router: tt.fields.router,
			}
			api.addPostHandler(tt.args.w, tt.args.r)
		})
	}
}

func TestAPI_deletePostHandler(t *testing.T) {
	type fields struct {
		db     storage.Interface
		router *mux.Router
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	var tests []struct {
		name   string
		fields fields
		args   args
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &API{
				db:     tt.fields.db,
				router: tt.fields.router,
			}
			api.deletePostHandler(tt.args.w, tt.args.r)
		})
	}
}

func TestAPI_endpoints(t *testing.T) {
	type fields struct {
		db     storage.Interface
		router *mux.Router
	}
	var tests []struct {
		name   string
		fields fields
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &API{
				db:     tt.fields.db,
				router: tt.fields.router,
			}
			api.endpoints()
		})
	}
}

func TestAPI_postsHandler(t *testing.T) {
	type fields struct {
		db     storage.Interface
		router *mux.Router
	}
	type args struct {
		w   http.ResponseWriter
		in1 *http.Request
	}
	var tests []struct {
		name   string
		fields fields
		args   args
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &API{
				db:     tt.fields.db,
				router: tt.fields.router,
			}
			api.postsHandler(tt.args.w, tt.args.in1)
		})
	}
}

func TestAPI_updatePostHandler(t *testing.T) {
	type fields struct {
		db     storage.Interface
		router *mux.Router
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	var tests []struct {
		name   string
		fields fields
		args   args
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &API{
				db:     tt.fields.db,
				router: tt.fields.router,
			}
			api.updatePostHandler(tt.args.w, tt.args.r)
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		db storage.Interface
	}
	var tests []struct {
		name string
		args args
		want *API
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
