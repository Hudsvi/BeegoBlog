package work

import (
	"Blog/models"
	log "Blog/utils/log"
	"github.com/astaxie/beego/orm"
)

func GetBlogList(uid string, workId int, page int, limit int) (posts []*models.Post) {
	o := orm.NewOrm()

	//起始位置offSet=(page-1)*limit。例如第1页，offSet=(1-1)*15=0，等同于MySQL的LIMIT offset limit->>limit 0 15
	//代表第1-15条记录。第2页，offSet=(2-1)*15=15，等同于MySQL的LIMIT offset limit->>limit 15 15 代表第1-15条记录。……
	offSet := (page - 1) * limit
	//根据one To many查询Posts基本字段,-id的“-”表示降序
	//值得注意的是，Beego的Limit方法的两个参数含义分别为（limit，offset)，与Mysql的LIMIT恰恰相反。
	if workId == -1 {
		//不进行work_id过滤处理
		_, err = o.QueryTable("post").
			Filter("User__Phone", uid).
			OrderBy("-id").
			Limit(limit, offSet).
			All(&posts, "Id", "Title", "Created", "ReadingNums", "CommentNums")
		if err != nil {
			//错误处理
			log.Println("GetBlogList()|Posts:", err)
		}
	} else {
		_, err = o.QueryTable("post").
			Filter("User__Phone", uid).
			Filter("Work__Id", workId).
			OrderBy("-id").
			Limit(limit, offSet).
			All(&posts, "Id", "Title", "Created", "ReadingNums", "CommentNums")
		if err != nil {
			//错误处理
			log.Println("GetBlogList()|Posts:", err)
		}
	}

	//获取用户Name
	var user models.User
	err = o.QueryTable("user").Filter("Phone", uid).One(&user, "Name")

	for i := 0; i < len(posts); i++ {
		posts[i].User = &user

		//根据博客Id，查询各个博文的Tags
		_, err = o.QueryTable("tag").Filter("Posts__Post__Id", posts[i].Id).All(
			&posts[i].Tags, "Name")
		if err != nil {
			//错误处理
			log.Println("GetBlogList()|Posts|Tags:", err)
		}
	}
	return
}
