package view

import (
	"blogsystem/common"
	"blogsystem/service"
	"log"
	"net/http"
	"strconv"
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
