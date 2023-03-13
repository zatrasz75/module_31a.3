package memdb

import (
	Interface "GoNews/pkg/storage"
)

// Store Хранилище данных.
type Store struct{}

// New Конструктор объекта хранилища.
func New() *Store {
	return new(Store)
}

func (s *Store) Posts() ([]Interface.Post, error) {
	return posts, nil
}

func (s *Store) AddPost(post Interface.Post) (int, error) {
	return 0, nil
}
func (s *Store) UpdatePost(post Interface.Post) error {
	return nil
}
func (s *Store) DeletePost(post Interface.Post) error {
	return nil
}

func (s *Store) Postsuser() ([]Interface.Authors, error) {
	return nil, nil
}

func (s *Store) AddPostuser(Interface.Authors) (int, error) {
	return 0, nil
}

func (s *Store) UpdatePostuser(Interface.Authors) error {
	return nil
}

func (s *Store) DeletePostuser(Interface.Authors) error {
	return nil
}

var name = []Interface.Authors{
	{
		Id:   1,
		Name: "Михаил Владимирович",
	},
}

var posts = []Interface.Post{
	{
		ID:      1,
		Title:   "Эффективный переход",
		Content: "Go - это новый язык. Хотя он заимствует идеи из существующих языков, он обладает необычными свойствами, которые делают эффективные программы Go отличными по характеру от программ, написанных на его родственниках. Прямой перевод программы на C++ или Java в Go вряд ли приведет к удовлетворительному результату — Java-программы написаны на Java, а не на Go. С другой стороны, размышление о проблеме с точки зрения Go могло бы привести к успешной, но совершенно иной программе. Другими словами, чтобы писать Go хорошо, важно понимать его свойства и идиомы. Также важно знать установленные соглашения для программирования в Go, такие как присвоение имен, форматирование, построение программ и так далее, чтобы другим программистам Go было легко понять программы, которые вы пишете.",
	},
	{
		ID:      2,
		Title:   "Модель памяти Go",
		Content: "Модель памяти Go определяет условия, при которых при чтении переменной в одной подпрограмме может быть гарантировано соблюдение значений, полученных при записи в ту же переменную в другой подпрограмме.",
	},
}
