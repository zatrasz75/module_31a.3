package mongo

import (
	Interface "GoNews/pkg/storage"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

// Storage Хранилище данных.
type Storage struct {
	Db *mongo.Client
}

const (
	databaseName      = "GoNews" // Имя учебной БД
	collectionPost    = "posts"  // Имя коллекции в учебной БД
	collectionAuthors = "authors"
)

// New Конструктор, принимает строку подключения к БД.
func New(ctx context.Context, constr string) (*Storage, error) {
	mongoOpts := options.Client().ApplyURI(constr)
	client, err := mongo.Connect(ctx, mongoOpts)
	if err != nil {
		log.Fatal(err)
	}
	// не забываем закрывать ресурсы
	s := Storage{
		Db: client,
	}
	return &s, nil
}

func (mg *Storage) Postsuser() ([]Interface.Authors, error) {
	collection := mg.Db.Database(databaseName).Collection(collectionAuthors)
	filter := bson.D{}
	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer func(cur *mongo.Cursor, ctx context.Context) {
		err := cur.Close(ctx)
		if err != nil {
			panic(err)
		}
	}(cur, context.Background())

	var data []Interface.Authors
	for cur.Next(context.Background()) {
		var l Interface.Authors
		err := cur.Decode(&l)
		if err != nil {
			return nil, err
		}
		data = append(data, l)
	}
	return data, cur.Err()
}

func (mg *Storage) Posts() ([]Interface.Post, error) {
	collection := mg.Db.Database(databaseName).Collection(collectionPost)
	filter := bson.D{}
	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer func(cur *mongo.Cursor, ctx context.Context) {
		err := cur.Close(ctx)
		if err != nil {
			panic(err)
		}
	}(cur, context.Background())

	var data []Interface.Post
	for cur.Next(context.Background()) {
		var l Interface.Post
		err := cur.Decode(&l)
		if err != nil {
			return nil, err
		}
		data = append(data, l)
	}
	return data, cur.Err()
}

func (mg *Storage) AddPostuser(p Interface.Authors) (int, error) {
	collection := mg.Db.Database(databaseName).Collection(collectionAuthors)
	_, err := collection.InsertOne(context.Background(), p)
	if err != nil {
		panic(err)
	}
	return 0, nil
}

func (mg *Storage) AddPost(p Interface.Post) (int, error) {
	collection := mg.Db.Database(databaseName).Collection(collectionPost)
	_, err := collection.InsertOne(context.Background(), p)
	if err != nil {
		return 0, err
	}
	return 0, nil
}

func (mg *Storage) UpdatePostuser(p Interface.Authors) error {
	collection := mg.Db.Database(databaseName).Collection(collectionAuthors)
	filter := bson.D{{"id", p.Id}}
	update := bson.D{{"$set", bson.D{{"name", p.Name}}}}
	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (mg *Storage) UpdatePost(p Interface.Post) error {
	collection := mg.Db.Database(databaseName).Collection(collectionPost)
	filter := bson.D{{"title", p.Title}}
	update := bson.D{{"$set", bson.D{{"content", p.Content}}}}
	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (mg *Storage) DeletePostuser(p Interface.Authors) error {
	collection := mg.Db.Database(databaseName).Collection(collectionAuthors)
	filter := bson.D{{"name", p.Name}}
	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}
	return nil
}

func (mg *Storage) DeletePost(p Interface.Post) error {
	collection := mg.Db.Database(databaseName).Collection(collectionPost)
	filter := bson.D{{"title", p.Title}}
	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return err
	}
	return nil
}
