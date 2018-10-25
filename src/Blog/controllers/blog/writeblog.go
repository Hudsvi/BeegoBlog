package blog

import (
	"Blog/models"
	"Blog/models/write"
	"github.com/astaxie/beego"
)

type WriteBlogController struct {
	beego.Controller
}

func (c *WriteBlogController) Get() {
	//临时身份验证
	if c.GetString("passwd", "") != "admin" {
		c.Abort("404")
	} else {
		var draftNums int            //草稿数量
		var parentList []*models.Tag //最高级分类标签，用于筛选
		var tagList []*models.Tag    //一个最高级标签，有多个子标签
		var workList []*models.Work  //文集

		c.Layout = "layout.html"
		c.TplName = "write/content.html"
		c.LayoutSections = make(map[string]string)
		c.LayoutSections["Head"] = "write/head.html"

		draftNums = write.GetDraftNums("13594777290")
		parentList = write.GetParentList()
		tagList = write.GetTagList()
		workList = write.GetWorkList("13594777290")

		c.Data["BlogTitle"] = "写博客"
		c.Data["draftNums"] = draftNums
		c.Data["parentList"] = parentList
		c.Data["tagList"] = tagList
		c.Data["workList"] = workList
	}

}
