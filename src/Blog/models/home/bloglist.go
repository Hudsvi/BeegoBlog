package home

import (
	"Blog/models"
	"github.com/astaxie/beego/orm"
	log "Blog/utils/log"
)

func GetBlogInfo(uid string) (post []*models.Post) {
	o := orm.NewOrm()
	//根据one To many查询Posts基本字段,-Created的“-”表示降序
	_, err = o.QueryTable("post").
		Filter("User__Phone", uid).
		OrderBy("-id").
		Limit(15).
		All(&post, "Id", "Title", "Created","ReadingNums","CommentNums")
	if err != nil {
		//错误处理
		log.Println("GetUserInfo()|Posts:", err)
	}
	//User只用于保存用户Name字段，故没采用RelatedSel()关联查询
	var user models.User
	err = o.QueryTable("user").Filter("Phone", uid).One(&user, "Name")

	for i := 0; i < len(post); i++ {
		post[i].User=&user
		//根据博客Id，查询各个博文的Tags
		_, err = o.QueryTable("tag").Filter("Posts__Post__Id", post[i].Id).All(
			&post[i].Tags, "Id", "Name")
		if err != nil {
			//错误处理
			log.Println("GetUserInfo()|Posts|Tags:", err)
		}

	}
	return
}
