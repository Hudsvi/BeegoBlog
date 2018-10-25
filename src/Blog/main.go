package main

import (
	_ "Blog/routers"
	"Blog/utils/funcmap"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"
	_ "Blog/models"

)
func main() {
	beego.Run()
}
func init(){
	//beego默认会初始化Session，如果要自定义部分参数，可使用下面的方式
	sessionConfig := &session.ManagerConfig{
		CookieName:"blogsession",
		EnableSetCookie: true,
		Gclifetime:3600,
		Maxlifetime: 3600,
		DisableHTTPOnly:false,
		Secure: false,
		CookieLifeTime: 3600,
	}
	beego.GlobalSessions, _ = session.NewManager("memory",sessionConfig)
	go beego.GlobalSessions.GC()

	beego.AddFuncMap("middle",funcmap.Middle)
	beego.AddFuncMap("ge_n",funcmap.GeN)
	beego.AddFuncMap("sub_n",funcmap.SubN)
}