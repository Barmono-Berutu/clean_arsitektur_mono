package domain

type UserUsecase interface {
	Login(email string, password string) (string, error)
	Register(user *User) error
}

// berfungsi untuk function logika error atau bind data
type PostsUsecase interface {
	GetAllPost() ([]*Posts, error)
	GetPostByID(id int) (*Posts, error)
	CreatePost(post *Posts) error
	DeletePost(id int) error
	UpdatePost(id int, post *Posts) error
}
