package controllers

import (
	"clean/domain"
	"clean/helper"
	"net/http"

	"github.com/labstack/echo/v4"
)

type PostsController struct {
	postsUsecase domain.PostsUsecase
}

func NewPostsController(a domain.PostsUsecase) *PostsController {
	return &PostsController{postsUsecase: a}
}

func (c *PostsController) GetAllPost(ctx echo.Context) error {
	posts, err := c.postsUsecase.GetAllPost()
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Gagal mengambil post data",
			"error":   err.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "sukses",
		"data":    posts,
	})
}

func (c *PostsController) GetPostById(ctx echo.Context) error {
	id, err := helper.GetIDParam(ctx)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "Format ID tidak valid",
			"error":   err.Error(),
		})
	}
	post, err := c.postsUsecase.GetPostByID(id)
	if err != nil {
		if err.Error() == "post data tidak ditemukan" {
			return ctx.JSON(http.StatusNotFound, map[string]string{
				"message": "Pos tidak ditemukan",
			})
		}
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Gagal mengambil post data",
			"error":   err.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "sukses",
		"data":    post,
	})
}

func (c *PostsController) CreatePost(ctx echo.Context) error {
	post := new(domain.Posts)
	if err := ctx.Bind(post); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "Gagal mengikat data",
			"error":   err.Error(),
		})
	}
	if err := c.postsUsecase.CreatePost(post); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Gagal membuat pos",
			"error":   err.Error(),
		})
	}
	return ctx.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Pos berhasil dibuat",
		"data":    post,
	})
}

func (c *PostsController) UpdatePost(ctx echo.Context) error {
	id, err := helper.GetIDParam(ctx)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "Format ID tidak valid",
			"error":   err.Error(),
		})
	}
	post := new(domain.Posts)
	if err := ctx.Bind(post); err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "Gagal mengikat data",
			"error":   err.Error(),
		})
	}
	if err := c.postsUsecase.UpdatePost(id, post); err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Gagal memperbarui pos",
			"error":   err.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "Pos berhasil diperbarui",
		"data":    post,
	})
}

func (c *PostsController) DeletePost(ctx echo.Context) error {
	id, err := helper.GetIDParam(ctx)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]string{
			"message": "Format ID tidak valid",
			"error":   err.Error(),
		})
	}
	if err := c.postsUsecase.DeletePost(id); err != nil {
		if err.Error() == "post data tidak ditemukan" { // Ganti dengan cara pengecekan yang sesuai
			return ctx.JSON(http.StatusNotFound, map[string]string{
				"message": "Pos tidak ditemukan",
			})
		}
		return ctx.JSON(http.StatusInternalServerError, map[string]string{
			"message": "Gagal menghapus post data",
			"error":   err.Error(),
		})
	}
	return ctx.JSON(http.StatusOK, map[string]string{
		"message": "Pos berhasil dihapus",
	})
}
