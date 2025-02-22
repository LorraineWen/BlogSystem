package view

import (
	"blogsystem/common"
	"blogsystem/internal/handler"
	"net/http"
)

func (*ViewEntity) CategoryView(w http.ResponseWriter, r *http.Request) {
	categoryTempldate := common.Template.Category
	if err := r.ParseForm(); err != nil {
		categoryTempldate.WriteError(w, err)
	}
	path := r.URL.Path
	page := r.Form.Get("page")
	if page == "" {
		page = "1"
	}
	response, err := handler.Handler.CategoryViewHandler(path, page)
	if err != nil {
		categoryTempldate.WriteError(w, err)
	}
	categoryTempldate.WriteData(w, *response)
}
