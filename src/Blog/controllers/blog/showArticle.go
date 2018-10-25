package blog

import (
	"Blog/models"
	"Blog/models/article"
	"github.com/astaxie/beego"
	"math/rand"
	"strconv"
)

type ShowArticleController struct {
	beego.Controller
}

func (c *ShowArticleController) Get() {
	var post *models.Post
	var navList []*models.Tag
	var cmtList []*models.VisitorComment
	c.Layout = "layout.html"
	c.TplName = "article/content.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Head"] = "article/head.html"

	articleId := c.Ctx.Input.Param(":articleID")
	id, err := strconv.Atoi(articleId) //转换为int
	if err == nil {
		post = article.GetBlogDetails(id) //获取博客详细信息
	}
	navList = models.GetNavList() //菜单

	//博客阅读数增1,此方法并不能解决跨站伪造请求，比如客户端手动追加uuid
	uuid := c.GetString("uuid", "")
	if uuid != "" {
		if c.GetSession(uuid) != nil && c.GetSession(uuid) == 1 {

		} else {
			c.SetSession(uuid, 0)
			article.ReadingNumAdd(id)
			c.SetSession(uuid, 1)
		}
	}

	//获取评论列表
	cmtList = article.GetCommentList(id)

	c.Data["BlogTitle"] = post.User.Name + "的博客" //工具的栏标题
	c.Data["Action"] = "/blog/write"             //博客正文界面，工具栏右侧的行为
	c.Data["ActionName"] = "写博客"
	c.Data["navList"] = navList  //左侧菜单
	c.Data["blogContent"] = post //博客详情
	c.Data["cmtList"] = cmtList  //博客评论
}

func (c *ShowArticleController) Post() {
	var cmtContent string
	var articleID int
	var floor int
	var datetime string
	var vid string                  //访客的id,由UUID生成
	var portrait string             //头像
	uuid := c.GetString("uuid", "") //获取uuid
	if uuid != "" && len(uuid) > 6 {
		vid = "visitor_" + uuid[0:6] //生成vid
	} else {
		c.Abort("404")
	}
	aid := c.Ctx.Input.Param(":articleID")

	articleID, _ = strconv.Atoi(aid)       //获取文章ID
	cmtContent = c.GetString("cmtContent") //获取评论内容
	datetime = c.GetString("datetime")     //获取评论时间
	floor = article.GetFloor(articleID)    //获取楼层
	portrait = "/static/img/profiles/user" + strconv.Itoa(rand.Intn(24)) + ".jpg"
	comment := &models.VisitorComment{
		Content: cmtContent,
		Floor:   floor,
		Time:    datetime,
		Visitor: &models.Visitor{
			Id:       vid,
			Portrait: portrait,
		},
		Post: &models.Post{
			Id: articleID,
		},
	}
	article.PublishComment(comment)          //发表评论
	c.Redirect(c.Ctx.Request.Referer(), 302) //重定向到评论位置
}
