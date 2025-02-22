package handler

import (
	"blogsystem/config"
	"blogsystem/internal/logic"
)

func (*HandlerEntity) WritingViewHandler() (map[string]interface{}, error) {
	categorys, err := logic.Logic.GetCategorys()
	if err != nil {
		return nil, err
	}
	m := make(map[string]interface{})
	m["categorys"] = categorys
	m["CdnURL"] = config.Cfg.System.CdnURL
	m["Title"] = config.Cfg.Viewer.Title
	return m, nil
}
