package dao

import (
	"blogsystem/models"
	"log"
)

func GetCateGory() ([]models.Category, error) {
	rows, err := DB.Query("select * from blog_category")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var cateGories []models.Category
	for rows.Next() {
		var category models.Category
		err := rows.Scan(&category.Cid, &category.Name, &category.CreateAt, &category.UpdateAt)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		cateGories = append(cateGories, category)
	}
	return cateGories, nil
}
func GetPostPageByCategoryId(cId int) string {
	row := DB.QueryRow("select  name from blog_category where cid=?", cId)
	if row.Err() != nil {
		log.Println(row.Err())
	}
	var categoryName string
	err := row.Scan(&categoryName)
	if err != nil {
		log.Println(err)
	}
	return categoryName
}
