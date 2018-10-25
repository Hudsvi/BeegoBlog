package blog

import (
	"Blog/models"
	"Blog/models/work"
	"github.com/astaxie/beego"
	"strconv"
)

type ShowAllBlogController struct {
	beego.Controller
}

type PageInfo struct {
	CurrentPage int            //当前页码
	TotalPage   int            //总页码
	TotalNums   int            //总博客数量
	Limit       int            //每页博客数量
	LastPage    bool           //最后一页
	FirstPage   bool           //第一页
	Posts       []*models.Post //博客列表
	UID         string         //用户ID，页面跳转时需要追加在URL后面
	WorkID      int            //文集ID，页面跳转时需要追加在URL后面
}

func (c *ShowAllBlogController) Get() {
	//获取路径/blog/works/:userID的参数userID
	uid := c.Ctx.Input.Param(":userID")

	//获取work_id参数，下面使用了最底层的Request对象获取,还可以使用高级点的c.Ctx.Input.Query()方法，当然也可以
	//使用最简单的GetString(),GetInt()等方法。
	//注意：下面获取url参数的方式并不推荐，建议使用Get系列方法，极大简化代码以及增大可读性。
	c.Ctx.Request.ParseForm()
	if c.Ctx.Request.Form["work_id"] == nil || c.Ctx.Request.Form["work_id"][0] == "" {
		//重定向时追加work_id=-1，这里用来标识查询所有文集下的博客
		c.Redirect("/blog/works/"+uid+"?work_id=-1&page=1", 302)
	}
	wid := c.Ctx.Request.Form["work_id"][0]
	workId, _ := strconv.Atoi(wid)

	//获取页码参数page,没有页码参数或者参数不正确时，设置默认值1
	page, _ := c.GetInt("page", 0)
	if page < 1 {
		c.Redirect("/blog/works/"+uid+"?work_id="+wid+"&page=1", 302)
	}
	var navList []*models.Tag       //菜单Item
	var worksInfo []*work.WorksInfo //文集信息
	var posts []*models.Post
	var pageInfo *PageInfo
	var totalNums int  //总博客数
	var totalPages int //总页码
	var limit int      //每页数量

	c.Layout = "layout.html"
	c.TplName = "works/content.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Head"] = "works/head.html"

	navList = models.GetNavList() //菜单

	worksInfo = work.GetWorksInfo(uid) //查询各个文集的数量

	//初始化pageInfo的一些信息
	limit = 10
	for p := 1; p < len(worksInfo); p++ {
		if workId == -1 {
			totalNums = worksInfo[0].Num
			break
		} else if workId == worksInfo[p].Work.Id {
			totalNums = worksInfo[p].Num
			break
		}
	}
	if totalNums%limit == 0 {
		totalPages = totalNums / limit
	} else {
		totalPages = totalNums/limit + 1
	}
	posts = work.GetBlogList(uid, workId, page, limit) //查询某文集下第page页
	pageInfo = &PageInfo{}
	pageInfo.Limit = limit
	pageInfo.CurrentPage = page
	pageInfo.TotalPage = totalPages
	pageInfo.FirstPage = page == 1
	pageInfo.TotalNums = totalNums
	pageInfo.LastPage = page == pageInfo.TotalPage || pageInfo.TotalNums == 0
	pageInfo.Posts = posts
	pageInfo.UID = uid
	pageInfo.WorkID = workId
	//命名Toolbar的标题
	if workId == -1 {
		c.Data["BlogTitle"] = "所有文集" //默认
	} else {
		flag := true //避免这种情况：用户在url后面手动追加work_id,但并不存在该文集。默认禁止该行为，除非该work_id存在。
		for i := 1; i < len(worksInfo); i++ {
			if worksInfo[i].Work.Id == workId {
				flag = false
				c.Data["BlogTitle"] = worksInfo[i].Work.Name //文集名称
			}
		}
		if flag {
			//c.Abort("404")
			c.Redirect("/blog/works/"+uid+"?work_id=-1&page=1", 302) //处理措施改为重定向
		}
	}
	//模板数据
	c.Data["uid"] = uid
	c.Data["navList"] = navList
	c.Data["works"] = worksInfo
	c.Data["blogList"] = posts
	c.Data["pageInfo"] = pageInfo

}
