package view

import (
	"blogsystem/common"
	"blogsystem/service"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func (*View) Detail(w http.ResponseWriter, r *http.Request) {
	detail := common.Template.Detail
	path := r.URL.Path
	pidStr := strings.TrimPrefix(path, "/p/")
	pidStr = strings.TrimSuffix(pidStr, ".html")
	pid, err := strconv.Atoi(pidStr)
	if err != nil {
		log.Println(err)
		return
	}
	res, err := service.GetPostDetail(pid)
	if err != nil {
		log.Println(err)
		return
	}
	detail.WriteData(w, res)
}
