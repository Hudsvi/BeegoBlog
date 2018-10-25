package home

import (
	"Blog/models"
	"github.com/astaxie/beego/orm"
	log "Blog/utils/log"
)

var (
	err error
)

func GetUserInfo(uid string) (user *models.User) {
	o:= orm.NewOrm()
	user = &models.User{}
	getProfile(o, uid, user)
	getUserTags(o, uid, user)
	return
}
func getUserTags(o orm.Ormer, uid string, user *models.User) {
	//根据many To many查询Tags
	_, err = o.QueryTable("tag").Filter("Users__User__Phone", uid).All(&user.Tags,
		"Id", "Name")
	if err != nil {
		//错误处理
		log.Println("GetUserInfo()|Tags:", err)
	}
}
func getProfile(o orm.Ormer, uid string, user *models.User) {
	//查询user的相关字段，并根据one对one关联查询Profile
	err = o.QueryTable("user").Filter("Phone", uid).RelatedSel("profile").One(user, "Phone",
		"Name", "Area", "Signature")
	if err != nil {
		//错误处理
		log.Println("GetUserInfo()|Profile:", err)
	}
}
