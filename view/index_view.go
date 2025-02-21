package view

import (
	"blogsystem/common"
	"blogsystem/service"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func (this *View) IndexView(w http.ResponseWriter, r *http.Request) {
	index := common.Template.Index
	if err := r.ParseForm(); err != nil {
		log.Println(err)
		return
	}
	pageStr := r.Form.Get("page")
	page := 1
	if pageStr != "" {
		page, _ = strconv.Atoi(pageStr)
	}
	pageSize := 10
	hr, err := service.GetBlogCategory(page, pageSize)
	if err != nil {
		log.Println(err)
	}
	index.WriteData(w, hr)
}
func (this *View) CategoryView(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Println(err)
		return
	}
	pageStr := r.Form.Get("page")
	page := 1
	if pageStr != "" {
		page, _ = strconv.Atoi(pageStr)
	}
	pageSize := 10
	//获取路径参数1中的category_id
	categoryTemplate := common.Template.Category
	path := r.URL.Path
	//去掉前缀
	cidString := strings.TrimPrefix(path, "/c/")
	cid, err := strconv.Atoi(cidString)
	if err != nil {
		log.Println(err)
	}
	//对通过语言种类获取的文章进行分页
	categoryResponse, err := service.GetPostsByCategoryId(cid, page, pageSize)
	if err != nil {
		log.Println(err)
	}
	categoryTemplate.WriteData(w, categoryResponse)
}
