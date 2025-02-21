package view

import (
	"blogsystem/common"
	"blogsystem/config"
	"net/http"
)

func (this *View) LoginView(w http.ResponseWriter, r *http.Request) {
	loginTemplate := common.Template.Login
	loginTemplate.WriteData(w, config.Cfg.Viewer)
}
