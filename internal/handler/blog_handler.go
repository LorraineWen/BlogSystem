package handler

import (
	"blogsystem/config"
	"blogsystem/internal/logic"
	"blogsystem/internal/model/models"
	"strconv"
	"strings"
)

func (*HandlerEntity) BlogViewHandler(path string, page string) (*models.HomeData, error) {
	if page == "" {
		page = "1"
	}
	categorys, err := logic.Logic.GetCategorys()
	if err != nil {
		return nil, err
	}
	currentPage, _ := strconv.Atoi(page)
	slug := strings.TrimPrefix(path, "/")
	var post []models.PostMore
	var total int
	if slug != "" {
		//请求的是自定义的路径 文章的slug
		//查询文章的时候 需要按照slug查询
		post, total, err = logic.Logic.PostPageBySlug(currentPage, 10, slug)
		if err != nil {
			return nil, err
		}
	} else {
		post, total, err = logic.Logic.PostPage(currentPage, 10)
		if err != nil {
			return nil, err
		}
	}
	pagesAll := ((total - 1) / 10) + 1
	pages := []int{}
	for i := 1; i <= pagesAll; i++ {
		pages = append(pages, i)
	}
	hd := models.HomeData{
		config.Cfg.Viewer,
		*categorys,
		post,
		total,
		currentPage,
		pages,
		currentPage != pagesAll,
	}
	return &hd, nil
}
