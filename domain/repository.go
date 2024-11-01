package domain

type UserRepository interface {
	GetByUsername(email string) (*User, error)
	CreateUser(user *User) error
}

// untuk function database
type PostsRepository interface {
	GetAllPost() ([]*Posts, error)
	GetPostByID(id int) (*Posts, error)
	CreatePost(post *Posts) error
	DeletePost(id int) error
	UpdatePost(id int, post *Posts) error
}
