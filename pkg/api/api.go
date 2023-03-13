package api

import (
	"GoNews/pkg/storage"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

// API Программный интерфейс сервера GoNews
type API struct {
	db     storage.Interface
	router *mux.Router
}

// New Конструктор объекта API
func New(db storage.Interface) *API {
	api := API{
		db: db,
	}
	api.router = mux.NewRouter()
	api.endpoints()
	return &api
}

// Регистрация обработчиков API.
func (api *API) endpoints() {
	api.router.HandleFunc("/posts", api.postsHandler).Methods(http.MethodGet, http.MethodOptions)
	api.router.HandleFunc("/posts", api.addPostHandler).Methods(http.MethodPost, http.MethodOptions)
	api.router.HandleFunc("/posts", api.updatePostHandler).Methods(http.MethodPut, http.MethodOptions)
	api.router.HandleFunc("/posts", api.deletePostHandler).Methods(http.MethodDelete, http.MethodOptions)

	api.router.HandleFunc("/posts", api.postsHandlerUser).Methods(http.MethodGet, http.MethodOptions)
	api.router.HandleFunc("/posts", api.addPostHandlerUser).Methods(http.MethodPost, http.MethodOptions)
	api.router.HandleFunc("/posts", api.updatePostHandlerUser).Methods(http.MethodPut, http.MethodOptions)
	api.router.HandleFunc("/posts", api.deletePostHandlerUser).Methods(http.MethodDelete, http.MethodOptions)

	//api.router.HandleFunc("/posts", api.addPostHandlerUserM).Methods(http.MethodPost, http.MethodOptions)
}

// Router Получение маршрутизатора запросов.
// Требуется для передачи маршрутизатора веб-серверу.
func (api *API) Router() *mux.Router {
	return api.router
}

// Получение всех публикаций.
func (api *API) postsHandlerUser(w http.ResponseWriter, _ *http.Request) {
	posts, err := api.db.Postsuser()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	bytes, err := json.Marshal(posts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = w.Write(bytes)
	if err != nil {
		return
	}
}

// Получение всех публикаций.
func (api *API) postsHandler(w http.ResponseWriter, _ *http.Request) {
	posts, err := api.db.Posts()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	bytes, err := json.Marshal(posts)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = w.Write(bytes)
	if err != nil {
		return
	}
}

// Добавление публикации.
func (api *API) addPostHandlerUser(w http.ResponseWriter, r *http.Request) {
	var p storage.Authors
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = api.db.AddPostuser(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// Добавление публикации.
func (api *API) addPostHandler(w http.ResponseWriter, r *http.Request) {
	var p storage.Post
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = api.db.AddPost(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// Обновление публикации.
func (api *API) updatePostHandlerUser(w http.ResponseWriter, r *http.Request) {
	var p storage.Authors
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = api.db.UpdatePostuser(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// Обновление публикации.
func (api *API) updatePostHandler(w http.ResponseWriter, r *http.Request) {
	var p storage.Post
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = api.db.UpdatePost(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// Удаление публикации.
func (api *API) deletePostHandlerUser(w http.ResponseWriter, r *http.Request) {
	var p storage.Authors
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = api.db.DeletePostuser(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// Удаление публикации.
func (api *API) deletePostHandler(w http.ResponseWriter, r *http.Request) {
	var p storage.Post
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = api.db.DeletePost(p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
