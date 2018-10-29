package blog

import (
	"Blog/models"
	"Blog/models/write"
	"github.com/astaxie/beego"
	"strconv"
	"time"
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
	var action = c.GetString("action", "")                 //判断是保存草稿还是
	var created = time.Now().Format("2006-01-02") //创建时间,String()方法是完整格式化输出
	blogTitle = c.GetString("blogTitle", "")
	blogContent = c.GetString("myContent", "")
	workId, _ = c.GetInt("workId", -1)
	if workId == -1 {
		c.Abort("404")
	}

	c.Ctx.Request.ParseForm()
	ts := c.Ctx.Request.Form["cbox-tag"] //包含tags参数的多个值

	//存草稿
	if action == "save" {
		draft = &models.Draft{
			Work: &models.Work{},
			User: &models.User{},
			Tags: []*models.Tag{},
		}

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
		draft.Created = created
		draft.User.Phone = "13594777290"

		write.SaveDraft(draft)
	} else // 发表
	if action == "publish" {
		post = &models.Post{
			User: &models.User{},
			Work: &models.Work{},
			Tags: []*models.Tag{},
		}
		if ts != nil && len(ts) != 0 {
			for i := 0; i < len(ts); i++ {
				id, _ := strconv.Atoi(ts[i])
				//特别要注意切片的赋值方式，不能用数组tags[i]=...这种初始化。
				post.Tags = append(post.Tags, &models.Tag{Id: id})
			}
		} else {
			c.Abort("404")
		}
		post.Title = blogTitle
		post.Content = blogContent
		post.Created=created
		post.Work.Id = workId
		post.User.Phone = "13594777290"

		write.Publish(post)
	} else {
		c.Abort("404") //action不是上面两种
	}

}
