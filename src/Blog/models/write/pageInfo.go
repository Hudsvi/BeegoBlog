package write

import (
	"Blog/models"
	log "Blog/utils/log"
	"github.com/astaxie/beego/orm"
)

//初始化页面的：“筛选列表、标签列表、文集选择、草稿数量”几个模块的数据

//查询筛选列表
func GetParentList() (parentList []*models.Tag) {
	o := orm.NewOrm()
	_, err := o.QueryTable("Tag").Filter("ParentId", 0).All(&parentList, "Id", "Name")
	if err != nil {
		log.Println("GetParentList()|tags|err:", err)
	}
	return
}

//查询标签列表（未筛选前的完整标签列表）
func GetTagList() (tagList []*models.Tag) {
	o := orm.NewOrm()
	_, err := o.QueryTable("Tag").Exclude("ParentId", 0).All(&tagList)
	if err != nil {
		log.Println("GetTagList()|tags|err:", err)
	}
	return
}

//文集列表
func GetWorkList(uid string) (workList []*models.Work) {
	o := orm.NewOrm()
	_, err := o.QueryTable("Work").Filter("User__Phone", uid).All(&workList)
	if err != nil {
		log.Println("GetWorkList()|works|err:", err)
	}
	return
}

//查询草稿箱数目
func GetDraftNums(uid string) (count int) {
	o := orm.NewOrm()
	ct, err := o.QueryTable("draft").Filter("User__Phone", uid).Count()
	count = int(ct)
	if err != nil {
		log.Println("InitialWritePage()|getDraftNum()|count|err:", err)
	}
	return
}
