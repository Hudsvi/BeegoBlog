package home

import (
	"github.com/astaxie/beego/orm"
	log "Blog/utils/log"
)

//用户发表博客的数量、获得评论的数量....
type BlogNums struct {
	BlogNums        interface{}
	PublishComments interface{}
	ObtainComments  interface{}
	StarNums        interface{}
}

func GetBlogNums(uid string) (nums *BlogNums) {
	nums = &BlogNums{}
	o := orm.NewOrm()
	getBlogNums(nums, o, uid)
	maps := getOthers(o, uid)
	nums.PublishComments = maps[0]["publish_comments"]
	nums.ObtainComments = maps[0]["obtain_comments"]
	nums.StarNums = maps[0]["star_nums"]
	return
}
func getOthers(o orm.Ormer, uid string) []orm.Params {
	//查询其它3项的数量信息
	var maps []orm.Params
	_, err = o.Raw("select publish_comments,obtain_comments,star_nums from user where phone=?", uid).Values(&maps)
	if err != nil {
		log.Println("GetBlogNums()|Other|err:", err)
	}
	return maps
}
func getBlogNums(nums *BlogNums, o orm.Ormer, uid string) {
	//查询总博客数量
	nums.BlogNums, err = o.QueryTable("post").Filter("User__Phone", uid).Count()
	if err != nil {
		log.Println("GetBlogNums()|BlogNums|err:", err)
	}
}
