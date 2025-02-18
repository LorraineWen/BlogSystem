package router

import (
	"blogsystem/api"
	"blogsystem/view"
	"net/http"
)

func Router() {
	//返回页面的路由
	http.HandleFunc("/", view.HtmlView.IndexView)

	//返回静态资源的路由
	//当请求resource路径时，映射到public/resource目录下的资源，因为template中的js是直接请求/resource
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))

	//返回json数据的路由
	http.HandleFunc("/updateblog", api.HandleApi.UpdateBlog)
}
