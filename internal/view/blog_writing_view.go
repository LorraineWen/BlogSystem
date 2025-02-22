package view

import (
	"blogsystem/common"
	"blogsystem/internal/handler"
	"net/http"
)

func (*ViewEntity) WritingView(w http.ResponseWriter, r *http.Request) {
	writingTemplate := common.Template.Writing
	response, err := handler.Handler.WritingViewHandler()
	if err != nil {
		writingTemplate.WriteError(w, err)
	}
	writingTemplate.WriteData(w, response)
}
