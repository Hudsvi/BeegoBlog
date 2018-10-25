package controllers

import (
	"Blog/models"
	"Blog/models/home"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	uid:=c.GetString("userID","")
	if uid==""{
		c.Redirect("/blog/?userID=13594777290",302)
	}
	var navList []*models.Tag
	var user *models.User
	var blog []*models.Post
	var nums *home.BlogNums
	c.Layout = "index.html"         //布局文件
	c.TplName = "home/content.html" //模板文件
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Head"] = "home/head.html" //引入界面的js和css
	navList = models.GetNavList()               //初始化左侧菜单列表
	user = home.GetUserInfo(uid)      //获取用户信息
	if user.Name==""{
		c.Abort("404")
	}
	blog = home.GetBlogInfo(uid)      //获取博客列表信息
	nums = home.GetBlogNums(uid)      //获取博客总数，评论数等信息
	c.Data["navList"] = navList
	c.Data["userInfo"] = user
	c.Data["blogInfo"] = blog
	c.Data["blogNums"] = nums
}
func (c *MainController) Post() {
}
