package view

import (
	"blogsystem/common"
	"blogsystem/internal/handler"
	"net/http"
)

// 归档页面，根据日期对博客进行归档
func (*ViewEntity) PigeonholeView(w http.ResponseWriter, r *http.Request) {
	pigeonholeTemplate := common.Template.Pigeonhole
	response, err := handler.Handler.PigeonholeViewHandler()
	if err != nil {
		pigeonholeTemplate.WriteError(w, err)
	}
	pigeonholeTemplate.WriteData(w, *response)
}
