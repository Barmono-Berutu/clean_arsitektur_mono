package services

import (
	"clean/domain"
)

type PostsServiceImpl struct {
	postRepo domain.PostsRepository
}

func NewPostsService(postRepo domain.PostsRepository) domain.PostsUsecase {
	return &PostsServiceImpl{postRepo: postRepo}
}

func (s *PostsServiceImpl) GetAllPost() ([]*domain.Posts, error) {
	return s.postRepo.GetAllPost()
}

func (s *PostsServiceImpl) GetPostByID(id int) (*domain.Posts, error) {
	return s.postRepo.GetPostByID(id)
}

func (s *PostsServiceImpl) CreatePost(post *domain.Posts) error {
	return s.postRepo.CreatePost(post)
}

func (s *PostsServiceImpl) DeletePost(id int) error {
	return s.postRepo.DeletePost(id)
}

func (s *PostsServiceImpl) UpdatePost(id int, post *domain.Posts) error {
	return s.postRepo.UpdatePost(id, post)
}
