package storage

// Post - публикация.
type Post struct {
	ID          int
	Title       string
	Content     string
	AuthorID    int
	AuthorName  string
	Created_at  int
	PublishedAt int
}

type Authors struct {
	Id   int
	Name string
}

// Interface задаёт контракт на работу с БД.
type Interface interface {
	Posts() ([]Post, error)    // Получение всех публикаций
	AddPost(Post) (int, error) // Создание новой публикации
	UpdatePost(Post) error     // Обновление публикации
	DeletePost(Post) error     // Удаление публикации по ID

	Postsuser() ([]Authors, error)    // Получение всех публикаций
	AddPostuser(Authors) (int, error) // Создание новой публикации
	UpdatePostuser(Authors) error     // Обновление публикации
	DeletePostuser(Authors) error     // Удаление публикации по ID
}
