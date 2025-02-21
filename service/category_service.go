package service

import (
	"blogsystem/config"
	"blogsystem/dao"
	"blogsystem/models"
	"html/template"
)

func GetBlogCategory(page int, pageSize int) (*models.HomeData, error) {
	categories, err := dao.GetCateGory()
	if err != nil {
		return nil, err
	}
	posts, err := dao.GetBlog(page, pageSize)
	var postMores []models.PostMore
	for _, post := range posts {
		categoryName := dao.GetCategoryNameById(post.CategoryId)
		userName := dao.GetUserNameById(post.UserId)
		content := []rune(post.Content)
		if len(content) > 100 {
			content = content[:100]
		}
		postMore := models.PostMore{
			post.Pid,
			post.Title,
			post.Slug,
			template.HTML(content),
			post.CategoryId,
			categoryName,
			post.UserId,
			userName,
			post.ViewCount,
			post.Type,
			models.DateDay(post.CreateAt),
			models.DateDay(post.UpdateAt),
		}
		postMores = append(postMores, postMore)
	}
	total := dao.GetPostTotal()
	pageCount := (total-1)/10 + 1
	var pages []int
	for i := 0; i < pageCount; i++ {
		pages = append(pages, i+1)
	}
	var hr = &models.HomeData{
		config.Cfg.Viewer,
		categories,
		postMores,
		total,
		page,
		pages,
		page != pageCount,
	}
	return hr, nil
}
func GetPostsByCategoryId(cid int, page int, pageSize int) (*models.CategoryResponse, error) {
	categories, err := dao.GetCateGory()
	if err != nil {
		return nil, err
	}
	posts, err := dao.GetBlogByCategoryId(cid, page, pageSize)
	var postMores []models.PostMore
	for _, post := range posts {
		categoryName := dao.GetCategoryNameById(post.CategoryId)
		userName := dao.GetUserNameById(post.UserId)
		content := []rune(post.Content)
		if len(content) > 100 {
			content = content[:100]
		}
		postMore := models.PostMore{
			post.Pid,
			post.Title,
			post.Slug,
			template.HTML(content),
			post.CategoryId,
			categoryName,
			post.UserId,
			userName,
			post.ViewCount,
			post.Type,
			models.DateDay(post.CreateAt),
			models.DateDay(post.UpdateAt),
		}
		postMores = append(postMores, postMore)
	}
	total := dao.GetPostTotalByCategoryId(cid)
	pageCount := (total-1)/10 + 1
	var pages []int
	for i := 0; i < pageCount; i++ {
		pages = append(pages, i+1)
	}
	var hr = &models.HomeData{
		config.Cfg.Viewer,
		categories,
		postMores,
		total,
		page,
		pages,
		page != pageCount,
	}
	categoryResponse := &models.CategoryResponse{CategoryName: dao.GetCategoryNameById(cid), HomeData: *hr}
	return categoryResponse, nil
}
