package view

import (
	"blogsystem/common"
	"blogsystem/internal/handler"
	"net/http"
)

func (*ViewEntity) BlogView(w http.ResponseWriter, r *http.Request) {
	indexTemplate := common.Template.Index
	if err := r.ParseForm(); err != nil {
		indexTemplate.WriteError(w, err)
	}
	path := r.URL.Path
	page := r.Form.Get("page")
	response, err := handler.Handler.BlogViewHandler(path, page)
	if err != nil {
		indexTemplate.WriteError(w, err)
	}
	indexTemplate.WriteData(w, *response)
}
