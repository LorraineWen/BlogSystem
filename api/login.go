package api

import (
	"blogsystem/common"
	"blogsystem/service"
	"net/http"
)

func (*Api) Login(w http.ResponseWriter, r *http.Request) {
	params := common.GetRequestJsonParam(r)
	userName := params["username"].(string)
	password := params["passwd"].(string)
	data, err := service.Login(userName, password)
	if err != nil {
		common.Error(w, err)
		return
	}
	common.Success(w, data)
}
