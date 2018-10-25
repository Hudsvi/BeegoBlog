package routers

import (
	"Blog/controllers"
	"Blog/controllers/blog"
	"github.com/astaxie/beego"
	"html/template"
	"net/http"
)

func init() {
	beego.ErrorHandler("404",err404)                                                    //自定义错误404处理界面
    beego.Router("/blog", &controllers.MainController{})                                //首页
    beego.Router("/blog/write", &blog.WriteBlogController{})                            //编辑博客
    beego.Router("/blog/publish", &blog.PublishBlogController{})                        //发表博客
    beego.Router("/blog/article/:articleID", &blog.ShowArticleController{})             //博客正文
    beego.Router("/blog/works/:userID", &blog.ShowAllBlogController{})                  //查看所有博客分类
}

func err404(rw http.ResponseWriter, r *http.Request){
	t,_:= template.New("404.html").ParseFiles(beego.BConfig.WebConfig.ViewsPath+"/error/404.html")
	t.Execute(rw, nil)
}