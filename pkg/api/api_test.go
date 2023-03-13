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
	tests := []struct {
		name   string
		fields fields
		want   *mux.Router
	}{
		// TODO: Add test cases.
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
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
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

func TestAPI_addPostHandlerUser(t *testing.T) {
	type fields struct {
		db     storage.Interface
		router *mux.Router
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &API{
				db:     tt.fields.db,
				router: tt.fields.router,
			}
			api.addPostHandlerUser(tt.args.w, tt.args.r)
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
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
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

func TestAPI_deletePostHandlerUser(t *testing.T) {
	type fields struct {
		db     storage.Interface
		router *mux.Router
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &API{
				db:     tt.fields.db,
				router: tt.fields.router,
			}
			api.deletePostHandlerUser(tt.args.w, tt.args.r)
		})
	}
}

func TestAPI_endpoints(t *testing.T) {
	type fields struct {
		db     storage.Interface
		router *mux.Router
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
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
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
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

func TestAPI_postsHandlerUser(t *testing.T) {
	type fields struct {
		db     storage.Interface
		router *mux.Router
	}
	type args struct {
		w   http.ResponseWriter
		in1 *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &API{
				db:     tt.fields.db,
				router: tt.fields.router,
			}
			api.postsHandlerUser(tt.args.w, tt.args.in1)
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
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
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

func TestAPI_updatePostHandlerUser(t *testing.T) {
	type fields struct {
		db     storage.Interface
		router *mux.Router
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			api := &API{
				db:     tt.fields.db,
				router: tt.fields.router,
			}
			api.updatePostHandlerUser(tt.args.w, tt.args.r)
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		db storage.Interface
	}
	tests := []struct {
		name string
		args args
		want *API
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
