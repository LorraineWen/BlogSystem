package service

import (
	"blogsystem/dao"
	"blogsystem/models"
	"blogsystem/utils"
	"errors"
)

func Login(userName string, password string) (*models.LoginResponse, error) {
	password = utils.Md5Crypt(password)
	user := dao.GetUser(userName, password)
	if user == nil {
		return nil, errors.New("账号密码不正确")
	}
	uid := user.Uid
	token, err := utils.Award(&uid)
	if err != nil {
		return nil, errors.New("token error")
	}
	var userInfo models.UserInfo = models.UserInfo{Uid: user.Uid, UserName: user.UserName, Avatar: user.Avatar}
	lr := &models.LoginResponse{
		token,
		userInfo,
	}
	return lr, nil
}
