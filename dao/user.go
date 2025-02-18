package dao

import "log"

func GetUserNameById(userId int) string {
	row := DB.QueryRow("select user_name from blog_user where uid=?", userId)
	var name string
	err := row.Scan(&name)
	if err != nil {
		log.Println(err)
	}
	return name
}
