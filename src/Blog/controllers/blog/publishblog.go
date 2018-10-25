package blog

import (
	"Blog/models"
	"Blog/models/write"
	"github.com/astaxie/beego"
	"log"
	"strconv"
)

type PublishBlogController struct {
	beego.Controller
}

func (c *PublishBlogController) Get() {
	c.Abort("404")
}

func (c *PublishBlogController) Post() {
	var blogTitle string
	var blogContent string
	var workId int
	var draft *models.Draft
	var post *models.Post
	var action = c.GetString("action", "")

	//存草稿
	if action == "save" {
		draft = &models.Draft{
			Work: &models.Work{},
			User: &models.User{},
			Tags: []*models.Tag{},
		}

		blogTitle = c.GetString("blogTitle", "")
		blogContent = c.GetString("blogContent", "")
		workId, _ = c.GetInt("workId", -1)
		if workId == -1 {
			c.Abort("404")
		}
		c.Ctx.Request.ParseForm()
		ts := c.Ctx.Request.Form["tags"] //包含tags参数的多个值
		if ts != nil && len(ts) != 0 {
			for i := 0; i < len(ts); i++ {
				id, _ := strconv.Atoi(ts[i])
				//特别要注意切片的赋值方式，不能用数组tags[i]=...这种初始化。
				draft.Tags = append(draft.Tags, &models.Tag{Id: id})
			}
		} else {
			c.Abort("404")
		}
		draft.Title = blogTitle
		draft.Content = blogContent
		draft.Work.Id = workId
		draft.User.Phone = "13594777290"

		write.SaveDraft(draft)
	} else // 发表
	if action == "publish" {
		post = &models.Post{
			User: &models.User{},
			Work: &models.Work{},
			Tags: []*models.Tag{},
		}
		log.Println(post)
	}

}
