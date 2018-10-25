package article

import (
	"Blog/models"
	log "Blog/utils/log"
	"github.com/astaxie/beego/orm"
)

func GetCommentList(articleId int) (cmtList []*models.VisitorComment) {
	o := orm.NewOrm()
	_, err = o.QueryTable("visitor_comment").Filter("Post__Id", articleId).
		RelatedSel("visitor").OrderBy("-floor").All(&cmtList)
	if err != nil {
		log.Println("GetCommentList（）|cmtlist|err:", err)
	}
	return
}
func PublishComment(comment *models.VisitorComment) {
	o := orm.NewOrm()
	_, err = o.Insert(comment) //插入评论
	if err != nil {
		log.Println("PublishComment()|Insert|err:", err)
	}
	updateCmtNum(comment.Post.Id) //post表评论数加1，博主获得的总评论数加1
	addVisitor(comment.Visitor)   //将访客信息存入数据库
}

func GetFloor(aid int) (floor int) {
	o := orm.NewOrm()
	err = o.Raw("select MAX(floor)+1 from visitor_comment where post_id=?", aid).QueryRow(&floor)
	if floor == 0 && err == nil {
		floor = 1; //当visitor_comment中不存在post_id为aid的文章时，表示当前文章没有评论，floor自动设置为1
	}
	if err != nil {
		log.Println("GetFloor()|Raw()|err:", err)
	}
	return
}
func updateCmtNum(id int) {
	o := orm.NewOrm()
	_, err = o.Raw("update post set comment_nums=comment_nums+1 where id=?", id).Exec()
	if err != nil {
		log.Println("updateCmtNum()|post|err:", err)
	}
	//注意,涉及外键的字段时并不推荐使用Raw查询，因为外键字段的字段名会随着依赖表的表名而改变，比如user_id依赖了User表。
	//出于便捷原因，这里直接使用Raw查询。使用orm高级查询会更利于维护。
	_, err = o.Raw("update user set obtain_comments=obtain_comments+1 where phone in (select user_id from post where id=? )", id).Exec()
	if err != nil {
		log.Println("updateCmtNum()|user|err:", err)
	}
}

func addVisitor(visitor *models.Visitor) {
	o := orm.NewOrm()
	_, err = o.Insert(visitor)
	if err != nil {
		log.Println("PublishComment()|addVisitor()|err:", err)
	}
}
