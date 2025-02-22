package dao

import (
	"blogsystem/internal/model/models"
	"errors"
)

func (this *DaoEntity) GetBlogCount() (int, error) {
	row := this.DB.QueryRow("select count(1) from blog_post")
	var total int
	err := row.Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}
func (this *DaoEntity) GetPostCountBySlug(slug string) (int, error) {
	row := this.DB.QueryRow("select count(1) from blog_post where slug=?", slug)
	var total int
	err := row.Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}
func (this *DaoEntity) GetPostCountCategory(categoryId int) (int, error) {
	row := this.DB.QueryRow("select count(1) from blog_post where category_id=?", categoryId)
	var total int
	err := row.Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (this *DaoEntity) GetPostPageCategory(page int, pageSize int, categoryId int) (*[]models.Post, error) {
	page = (page - 1) * pageSize
	ret, err := this.DB.Query("select * from blog_post where category_id=? limit ?,?", categoryId, page, pageSize)
	if err != nil {
		return nil, err
	}
	var posts []models.Post
	for ret.Next() {
		var post models.Post
		err = ret.Scan(
			&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return &posts, nil
}

func (this *DaoEntity) GetPostPageBySlug(page int, pageSize int, slug string) (*[]models.Post, error) {
	page = (page - 1) * pageSize
	ret, err := this.DB.Query("select * from blog_post where slug = ? limit ?,?", slug, page, pageSize)
	if err != nil {
		return nil, err
	}
	var posts []models.Post
	for ret.Next() {
		var post models.Post
		err = ret.Scan(
			&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return &posts, nil
}

func (this *DaoEntity) GetPostPage(page int, pageSize int) (*[]models.Post, error) {
	page = (page - 1) * pageSize
	ret, err := this.DB.Query("select * from blog_post limit ?,?", page, pageSize)
	if err != nil {
		return nil, err
	}
	var posts []models.Post
	for ret.Next() {
		var post models.Post
		err = ret.Scan(
			&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return &posts, nil
}

func (this *DaoEntity) GetPostAll() (*[]models.Post, error) {
	ret, err := this.DB.Query("select * from blog_post")
	if err != nil {
		return nil, err
	}
	var posts []models.Post
	for ret.Next() {
		var post models.Post
		_ = ret.Scan(
			&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt)
		posts = append(posts, post)
	}
	return &posts, nil
}
func (this *DaoEntity) UpdatePost(post *models.Post) error {
	_, err := this.DB.Exec("update  blog_post set title=?,"+
		"content=?,"+
		"markdown=?,"+
		"category_id=?,"+
		"type=?,"+
		"slug=?,"+
		"update_at=? where user_id=? and pid=?",
		post.Title,
		post.Content,
		post.Markdown,
		post.CategoryId,
		post.Type,
		post.Slug,
		post.UpdateAt,
		post.UserId,
		post.Pid)
	if err != nil {
		return err
	}
	return nil
}
func (this *DaoEntity) GetPostById(id int) (*models.Post, error) {
	row := this.DB.QueryRow("select * from blog_post where pid=?", id)
	if row.Err() != nil {
		return nil, row.Err()
	}
	post := new(models.Post)
	err := row.Scan(
		&post.Pid,
		&post.Title,
		&post.Content,
		&post.Markdown,
		&post.CategoryId,
		&post.UserId,
		&post.ViewCount,
		&post.Type,
		&post.Slug,
		&post.CreateAt,
		&post.UpdateAt)
	if err != nil {
		return nil, err
	}
	return post, nil
}
func (this *DaoEntity) SavePost(post *models.Post) error {
	ret, err := this.DB.Exec("insert into blog_post (title,content,markdown,category_id,user_id,view_count,type,slug,create_at,update_at) values(?,?,?,?,?,?,?,?,?,?)",
		post.Title,
		post.Content,
		post.Markdown,
		post.CategoryId,
		post.UserId,
		post.ViewCount,
		post.Type,
		post.Slug,
		post.CreateAt,
		post.UpdateAt)
	if err != nil {
		return err
	}
	id, err := ret.LastInsertId()
	if err != nil {
		return err
	}
	post.Pid = int(id)
	return nil
}

func (this *DaoEntity) PostSearch(search string) (*[]models.Post, error) {
	ret, err := this.DB.Query("select * from blog_post where title like ?", "%"+search+"%")
	if err != nil {
		return nil, err
	}
	var posts []models.Post
	if ret == nil {
		return &posts, errors.New("没有任何结果")
	}
	for ret.Next() {
		var post models.Post
		err = ret.Scan(
			&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return &posts, nil
}
