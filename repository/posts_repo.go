package repository

import (
	"clean/domain"

	"gorm.io/gorm"
)

// logika database
type PostsRepositoryImpl struct {
	DB *gorm.DB
}

func NewPostsRepository(db *gorm.DB) domain.PostsRepository {
	return &PostsRepositoryImpl{DB: db}
}

func (r *PostsRepositoryImpl) GetAllPost() ([]*domain.Posts, error) {
	var posts []*domain.Posts
	err := r.DB.Find(&posts).Error
	return posts, err
}

func (r *PostsRepositoryImpl) GetPostByID(id int) (*domain.Posts, error) {
	var posts *domain.Posts
	err := r.DB.Where("id = ?", id).First(&posts).Error
	return posts, err
}
func (r *PostsRepositoryImpl) CreatePost(posts *domain.Posts) error {
	return r.DB.Create(posts).Error
}
func (r *PostsRepositoryImpl) DeletePost(id int) error {
	return r.DB.Delete(&domain.User{}, id).Error
}
func (r *PostsRepositoryImpl) UpdatePost(id int, posts *domain.Posts) error {
	return r.DB.Where("id =?", id).Updates(posts).Error
}
