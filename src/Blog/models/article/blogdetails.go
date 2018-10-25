package article

import (
	"Blog/models"
	"github.com/astaxie/beego/orm"
	log "Blog/utils/log"
)

var (
	err error
)

func GetBlogDetails(articleID int) (post *models.Post) {
	o := orm.NewOrm()
	post = &models.Post{}
	getBlogInfo(o, articleID, post) //获取博客标题和正文
	getPostTags(o, articleID, post) //获取文章标签
	return
}
func getPostTags(o orm.Ormer, articleID int, post *models.Post) {
	_, err = o.QueryTable("tag").Filter("Posts__Post__Id", articleID).All(&post.Tags, "Id", "Name")
	if err != nil {
		log.Println("GetBlogDetails()|tags|err:", err)
	}
}
func getBlogInfo(o orm.Ormer, articleID int, post *models.Post) {
	err = o.QueryTable("post").Filter("Id", articleID).RelatedSel().One(post)
	if err != nil {
		log.Println("GetBlogDetails()|post|err:", err)
	}
}
