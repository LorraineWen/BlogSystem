package logic

import (
	"blogsystem/common"
	"blogsystem/config"
	"blogsystem/internal/model/dao"
	"blogsystem/internal/model/models"
	"blogsystem/pkg/utils"
	"errors"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"html/template"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func (*LogicEntity) PostByMonth() (*map[string][]models.Post, error) {
	posts, err := dao.Dao.GetPostAll()
	if err != nil {
		return nil, err
	}
	lines := make(map[string][]models.Post)
	for _, post := range *posts {
		month := post.CreateAt.Format("2006-01")
		lines[month] = append(lines[month], post)
	}
	return &lines, nil
}
func (*LogicEntity) PostPageByCategory(page int, pageSize int, categoryId int) (*[]models.PostMore, int, error) {
	posts, err := dao.Dao.GetPostPageCategory(page, pageSize, categoryId)
	if err != nil {
		return nil, 0, err
	}
	total, err := dao.Dao.GetPostCountCategory(categoryId)
	if err != nil {
		return nil, 0, err
	}
	var postMores []models.PostMore
	for _, post := range *posts {
		var pm models.PostMore
		pm.Pid = post.Pid
		pm.ViewCount = post.ViewCount
		pm.CategoryId = post.CategoryId
		pm.CategoryName, err = dao.Dao.GetCategoryNameById(post.CategoryId)
		if err != nil {
			return nil, 0, err
		}
		content := []rune(post.Content)
		if len(content) > 100 {
			content = []rune(post.Content)[:100]
		}
		pm.Content = template.HTML(content)
		pm.Title = post.Title
		pm.Slug = post.Slug
		pm.CreateAt = utils.Format(post.CreateAt)
		pm.UserName, err = dao.Dao.GetUserNameById(post.UserId)
		if err != nil {
			return nil, 0, err
		}
		postMores = append(postMores, pm)
	}
	return &postMores, total, nil
}

func (*LogicEntity) PostPageBySlug(page int, pageSize int, slug string) ([]models.PostMore, int, error) {
	posts, err := dao.Dao.GetPostPageBySlug(page, pageSize, slug)
	if err != nil {
		return nil, 0, err
	}
	total, err := dao.Dao.GetPostCountBySlug(slug)
	if err != nil {
		return nil, 0, err
	}
	var postMores []models.PostMore
	for _, post := range *posts {
		var pm models.PostMore
		pm.Pid = post.Pid
		pm.ViewCount = post.ViewCount
		pm.CategoryId = post.CategoryId
		pm.CategoryName, err = dao.Dao.GetCategoryNameById(post.CategoryId)
		if err != nil {
			return nil, 0, err
		}
		content := []rune(post.Content)
		if len(content) > 100 {
			content = []rune(post.Content)[:100]
		}
		pm.Content = template.HTML(content)
		pm.Title = post.Title
		pm.Slug = post.Slug
		pm.CreateAt = utils.Format(post.CreateAt)
		pm.UserName, err = dao.Dao.GetUserNameById(post.UserId)
		if err != nil {
			return nil, 0, err
		}
		postMores = append(postMores, pm)
	}
	return postMores, total, nil
}
func (*LogicEntity) PostPage(page int, pageSize int) ([]models.PostMore, int, error) {
	posts, err := dao.Dao.GetPostPage(page, pageSize)
	if err != nil {
		return nil, 0, err
	}
	total, err := dao.Dao.GetBlogCount()
	if err != nil {
		return nil, 0, err
	}
	var postMores []models.PostMore
	for _, post := range *posts {
		var pm models.PostMore
		pm.Pid = post.Pid
		pm.ViewCount = post.ViewCount
		pm.CategoryId = post.CategoryId
		pm.CategoryName, err = dao.Dao.GetCategoryNameById(post.CategoryId)
		if err != nil {
			return nil, 0, err
		}
		content := []rune(post.Content)
		if len(content) > 100 {
			content = []rune(post.Content)[:100]
		}
		pm.Content = template.HTML(content)
		pm.Title = post.Title
		pm.Slug = post.Slug
		pm.CreateAt = utils.Format(post.CreateAt)
		pm.UserName, err = dao.Dao.GetUserNameById(post.UserId)
		if err != nil {
			return nil, 0, err
		}
		postMores = append(postMores, pm)
	}
	return postMores, total, nil
}
func (*LogicEntity) GetPost(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	id := strings.TrimPrefix(path, "/api/v1/post/")
	pid, err := strconv.Atoi(id)
	if err != nil {
		common.Error(w, err)
		return
	}
	post, err := dao.Dao.GetPostById(pid)
	if err != nil {
		common.Error(w, err)
		return
	}
	common.ReturnSuccess(w, post)
}
func (*LogicEntity) AddOrUpdate(w http.ResponseWriter, r *http.Request) {
	//先判断是否为POST还是PUT POST为新增 PUT为更新
	method := r.Method
	switch method {
	case http.MethodPost:
		//获取登录用户信息
		token := r.Header.Get("Authorization")
		_, claims, err := utils.ParseToken(token)
		if err != nil {
			common.Error(w, errors.New("登录已过期"))
			return
		}
		uid := claims.Uid
		//获取参数
		param := common.GetRequestJsonParam(r)
		categoryId := param["categoryId"].(string)
		content := param["content"].(string)
		markdown := param["markdown"].(string)
		slug := param["slug"].(string)
		title := param["title"].(string)
		articleType := float64(0)
		if param["type"] != nil {
			articleType = param["type"].(float64)
		}
		post := new(models.Post)
		post.Title = title
		post.UserId = uid
		post.ViewCount = 0
		cId, _ := strconv.Atoi(categoryId)
		post.CategoryId = cId
		post.Markdown = markdown
		post.Slug = slug
		post.Type = int(articleType)
		post.Content = content
		post.CreateAt = time.Now()
		post.UpdateAt = time.Now()
		if err := dao.Dao.SavePost(post); err != nil {
			common.Error(w, errors.New("数据库错误"))
			return
		}
		common.ReturnSuccess(w, post)
		return
	case http.MethodPut:
		//获取登录用户信息
		token := r.Header.Get("Authorization")
		_, _, err := utils.ParseToken(token)
		if err != nil {
			common.Error(w, errors.New("登录已过期"))
			return
		}
		//获取参数
		param := common.GetRequestJsonParam(r)
		userId := param["userId"].(float64)
		categoryId := param["categoryId"].(float64)
		content := param["content"].(string)
		markdown := param["markdown"].(string)
		slug := param["slug"].(string)
		title := param["title"].(string)
		articleType := float64(0)
		if param["type"] != nil {
			articleType = param["type"].(float64)
		}
		pid := param["pid"].(float64)
		post := new(models.Post)
		post.Pid = int(pid)
		post.Title = title
		post.UserId = int(userId)
		post.CategoryId = int(categoryId)
		post.Markdown = markdown
		post.Slug = slug
		post.Type = int(articleType)
		post.Content = content
		post.CreateAt = time.Now()
		post.UpdateAt = time.Now()
		if err := dao.Dao.UpdatePost(post); err != nil {
			common.Error(w, errors.New("数据库错误"))
			return
		}
		common.ReturnSuccess(w, post)
		return

	}
}

func (*LogicEntity) PostSearch(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	searchStr := r.Form.Get("val")
	posts, err := dao.Dao.PostSearch(searchStr)
	if err != nil {
		common.Error(w, err)
		return
	}
	var searchResp []models.SearchResp
	for _, post := range *posts {
		var sr models.SearchResp
		sr.Pid = post.Pid
		sr.Title = post.Title
		searchResp = append(searchResp, sr)
	}
	common.ReturnSuccess(w, searchResp)
}

// 图片上传会使用七牛云存储，需要修改writing.js和config.toml，bucket的配置
func (*LogicEntity) UploadImage(w http.ResponseWriter, r *http.Request) {
	//自定义凭证有效期（示例2小时，Expires 单位为秒，为上传凭证的有效时间）
	bucket := "lorrainewen-blogsystem"
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	putPolicy.Expires = 7200 //示例2小时有效期
	mac := qbox.NewMac(config.Cfg.System.QiniuAccessKey, config.Cfg.System.QiniuSecretKey)
	upToken := putPolicy.UploadToken(mac)
	common.ReturnSuccess(w, upToken)
}
func (LogicEntity) GetBlogById(id int) (*models.Post, error) {
	return dao.Dao.GetPostById(id)
}
