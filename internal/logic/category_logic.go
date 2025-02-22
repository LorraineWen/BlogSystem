package logic

import (
	"blogsystem/internal/model/dao"
	"blogsystem/internal/model/models"
)

func (*LogicEntity) GetCategorys() (*[]models.Category, error) {
	return dao.Dao.GetCategorys()
}
func (*LogicEntity) GetCategoryNameById(id int) (string, error) {
	return dao.Dao.GetCategoryNameById(id)
}
