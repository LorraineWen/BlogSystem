package router

import (
	"blogsystem/internal/logic"
	"blogsystem/internal/view"
	"net/http"
)

func Routers() {
	//返回页面的路由
	http.HandleFunc("/", view.View.BlogView)
	http.HandleFunc("/login", view.View.LoginView)
	http.HandleFunc("/c/", view.View.CategoryView)
	http.HandleFunc("/pigeonhole", view.View.PigeonholeView)
	http.HandleFunc("/p/", view.View.BlogDetailView)
	http.HandleFunc("/writing", view.View.WritingView)
	//返回数据的路由
	http.HandleFunc("/api/v1/login", logic.Logic.Login)
	http.HandleFunc("/api/v1/qiniu/token", logic.Logic.UploadImage)
	http.HandleFunc("/api/v1/post", logic.Logic.AddOrUpdate)
	http.HandleFunc("/api/v1/post/", logic.Logic.GetPost)
	http.HandleFunc("/api/v1/post/search", logic.Logic.PostSearch)
	//设置资源目录为工作目录下的public/resource，当访问/resource时自动跳转这个目录
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))
}
