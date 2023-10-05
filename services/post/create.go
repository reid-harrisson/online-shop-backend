package post

import "PockitGolangBoilerplate/models"

func (postService *Service) Create(post *models.Post) {
	postService.DB.Create(post)
}
