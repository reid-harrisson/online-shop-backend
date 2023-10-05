package post

import "PockitGolangBoilerplate/models"

func (postService *Service) Delete(post *models.Post) {
	postService.DB.Delete(post)
}
