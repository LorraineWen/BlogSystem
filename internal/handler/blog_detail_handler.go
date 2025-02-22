package handler

import (
	"blogsystem/config"
	"blogsystem/internal/logic"
	"blogsystem/internal/model/models"
	"blogsystem/pkg/utils"
	"html/template"
	"strconv"
	"strings"
)

func (*HandlerEntity) BlogDetailHandler(path string) (*models.PostReseponse, error) {
	id := strings.TrimPrefix(path, "/p/")
	id = strings.TrimSuffix(id, ".html")
	pid, _ := strconv.Atoi(id)
	post, err := logic.Logic.GetBlogById(pid)
	if err != nil {
		return nil, err
	}
	var pm models.PostMore
	pm.UserName, err = logic.Logic.GetUserNameById(post.UserId)
	if err != nil {
		return nil, err
	}
	pm.Pid = post.Pid
	pm.ViewCount = post.ViewCount
	pm.CategoryId = post.CategoryId
	pm.CategoryName, err = logic.Logic.GetCategoryNameById(post.CategoryId)
	if err != nil {
		return nil, err
	}
	pm.Content = template.HTML(post.Content)
	pm.Title = post.Title
	pm.Slug = post.Slug
	pm.CreateAt = utils.Format(post.CreateAt)
	postResponse := models.PostReseponse{
		config.Cfg.Viewer,
		config.Cfg.System,
		pm,
	}
	return &postResponse, nil
}
