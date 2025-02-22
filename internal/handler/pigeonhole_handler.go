package handler

import (
	"blogsystem/config"
	"blogsystem/internal/logic"
	"blogsystem/internal/model/models"
)

func (*HandlerEntity) PigeonholeViewHandler() (*models.PigeonholeData, error) {
	categorys, err := logic.Logic.GetCategorys()
	if err != nil {
		return nil, err
	}
	lines, err := logic.Logic.PostByMonth()
	if err != nil {
		return nil, err
	}
	data := &models.PigeonholeData{
		config.Cfg.Viewer,
		config.Cfg.System,
		*categorys,
		lines,
	}
	return data, nil
}
