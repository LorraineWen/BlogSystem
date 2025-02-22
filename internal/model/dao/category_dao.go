package dao

import (
	"blogsystem/internal/model/models"
	"log"
)

func (this *DaoEntity) GetCategorys() (*[]models.Category, error) {
	ret, err := this.DB.Query("select * from blog_category")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var cs []models.Category
	for ret.Next() {
		var cat models.Category
		_ = ret.Scan(&cat.Cid, &cat.Name, &cat.CreateAt, &cat.UpdateAt)
		cs = append(cs, cat)
	}
	return &cs, nil
}
func (this *DaoEntity) GetCategoryNameById(id int) (string, error) {
	row := this.DB.QueryRow("select name from blog_category where cid=?", id)
	var name string
	err := row.Scan(&name)
	if err != nil {
		return "", err
	}
	return name, nil
}
