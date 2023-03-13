package main

import (
	"GoNews/pkg/api"
	"GoNews/pkg/storage"
	"GoNews/pkg/storage/memdb"
	"GoNews/pkg/storage/mongo"
	"GoNews/pkg/storage/postgres"
	"context"
	mongo2 "go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"time"
)

// Сервер GoNews.
type server struct {
	db  storage.Interface
	api *api.API
}

func main() {
	// Создаём объект сервера.
	var srv server
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	// Создаём объекты баз данных.
	//
	// БД в памяти.
	db := memdb.New()

	// Реляционная БД PostgresSQL.
	db2, err := postgres.New(ctx, "postgres://postgres:postgrespw@localhost:49153/posts")
	if err != nil {
		log.Fatal(err)
	}
	defer db2.Db.Close()

	// Документная БД MongoDB.
	db3, err := mongo.New(ctx, "mongodb://localhost:27017/")
	if err != nil {
		log.Fatal(err)
	}
	defer func(Db *mongo2.Client, ctx context.Context) {
		var err = Db.Disconnect(ctx)
		if err != nil {
			panic(err)
		}
	}(db3.Db, ctx)

	_, _, _ = db, db2, db3

	// Инициализируем хранилище сервера конкретной БД.
	srv.db = db3

	// Создаём объект API и регистрируем обработчики.
	srv.api = api.New(srv.db)

	// Запускаем веб-сервер на порту 8080 на всех интерфейсах.
	// Предаём серверу маршрутизатор запросов,
	// поэтому сервер будет все запросы отправлять на маршрутизатор.
	// Маршрутизатор будет выбирать нужный обработчик.
	err = http.ListenAndServe(":8080", srv.api.Router())
	if err != nil {
		log.Fatal(err)
	}
}
