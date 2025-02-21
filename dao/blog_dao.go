package dao

import (
	"blogsystem/models"
	"log"
)

func GetBlog(page, pageSize int) ([]models.Post, error) {
	page = (page - 1) * pageSize
	rows, err := DB.Query("select * from blog_post limit ?,?", page, pageSize)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(
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
			&post.UpdateAt,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}
func GetBlogByCategoryId(cid, page, pageSize int) ([]models.Post, error) {
	page = (page - 1) * pageSize
	rows, err := DB.Query("select * from blog_post where category_id = ? limit ?,?", cid, page, pageSize)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	var posts []models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(
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
			&post.UpdateAt,
		)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}
func GetPostTotal() int {
	row := DB.QueryRow("select count(1) from blog_post")
	var count int
	row.Scan(&count)
	return count
}
func GetPostTotalByCategoryId(cid int) int {
	row := DB.QueryRow("select count(1) from blog_post where category_id = ?", cid)
	var count int
	row.Scan(&count)
	return count
}
func GetPostDetailByid(pid int) (models.Post, error) {
	rows, err := DB.Query("select * from blog_post where pid = ?", pid)
	var post models.Post
	if err != nil {
		log.Fatal(err)
		return post, err
	}
	for rows.Next() {
		err := rows.Scan(
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
			&post.UpdateAt,
		)
		if err != nil {
			return post, err
		}
	}
	return post, nil
}
