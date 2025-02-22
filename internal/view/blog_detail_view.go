package view

import (
	"blogsystem/common"
	"blogsystem/internal/handler"
	"net/http"
)

func (*ViewEntity) BlogDetailView(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	response, err := handler.Handler.BlogDetailHandler(path)
	if err != nil {
		common.Error(w, err)
		return
	}
	common.Template.Detail.WriteData(w, *response)
}
