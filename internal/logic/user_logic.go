package logic

import (
	"blogsystem/common"
	"blogsystem/internal/model/dao"
	"blogsystem/internal/model/models"
	"blogsystem/pkg/utils"
	"errors"
	"net/http"
)

func (*LogicEntity) Login(w http.ResponseWriter, r *http.Request) {
	param := common.GetRequestJsonParam(r)
	username := param["username"].(string)
	passwd := param["passwd"].(string)
	passwd = utils.Md5Crypt(passwd, "mszlu")
	loginReq := new(models.LoginReq)
	loginReq.Name = username
	loginReq.Passwd = passwd
	user, dbError := dao.Dao.Login(loginReq)
	if dbError != nil {
		if dbError.IsNilError {
			dbError.Err = errors.New("账号密码不正确")
		}
		common.Error(w, dbError.Err)
		return
	}
	uid := user.Uid
	token, _ := utils.Award(&uid)
	loginRes := &models.LoginResp{Token: token, UserInfo: models.UserRes{user.Uid, user.UserName, user.Avatar}}
	common.ReturnSuccess(w, loginRes)
}
func (*LogicEntity) GetUserNameById(id int) (string, error) {
	return dao.Dao.GetUserNameById(id)
}
