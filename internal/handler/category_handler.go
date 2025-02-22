package handler

import (
	"blogsystem/config"
	"blogsystem/internal/logic"
	"blogsystem/internal/model/models"
	"strconv"
	"strings"
)

func (*HandlerEntity) CategoryViewHandler(path string, page string) (*models.CategoryData, error) {
	id := strings.TrimPrefix(path, "/c/")
	cId, _ := strconv.Atoi(id)
	currentPage, _ := strconv.Atoi(page)
	cName, err := logic.Logic.GetCategoryNameById(cId)
	if err != nil {
		return nil, err
	}
	categorys, err := logic.Logic.GetCategorys()
	if err != nil {
		return nil, err
	}
	post, total, err := logic.Logic.PostPageByCategory(currentPage, 10, cId)
	if err != nil {
		return nil, err
	}
	pagesAll := ((total - 1) / 10) + 1
	pages := []int{}
	for i := 1; i <= pagesAll; i++ {
		pages = append(pages, i)
	}
	hd := models.HomeData{
		config.Cfg.Viewer,
		*categorys,
		*post,
		total,
		currentPage,
		pages,
		currentPage != pagesAll,
	}
	var categoryData = &models.CategoryData{
		hd,
		cName,
	}
	return categoryData, nil
}
