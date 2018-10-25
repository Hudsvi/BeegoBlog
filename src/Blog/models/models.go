package models

import (
	log "Blog/utils/log"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/astaxie/beego/session/mysql"
)

type User struct {
	Phone           string `orm:"pk;size(11)"`
	Password        string `orm:"size(20);null"`
	Name            string `orm:"size(20)"`
	Area            string `orm:"size(20);null"`
	Signature       string `orm:"size(255);null"`
	PublishComments int    `orm:"size(6);null;default(0)"`
	ObtainComments  int    `orm:"size(6);null;default(0)"`
	StarNums        int    `orm:"size(6);null;default(0)"`

	Profile *Profile `orm:"rel(one);"`
	Tags    []*Tag   `orm:"rel(m2m)"`
	Posts   []*Post  `orm:"reverse(many)"`
	Works   []*Work  `orm:"reverse(many)"`
}
type Profile struct {
	Id       int    `orm:"auto;size(10)"`
	Age      int    `orm:"null;size(3)"`
	Portrait string `orm:"size(255);null"`
	Gender   string `orm:"size(1);null"`

	User *User `orm:"reverse(one);"`
}
type Post struct {
	Id              int               `orm:"auto;size(10)"`
	Title           string            `orm:"size(100)"`
	Content         string            `orm:"type(text)"`
	Created         string            `orm:"auto_now_add;type(date)"`
	ReadingNums     int               `orm:"size(6);null;default(0)"`
	CommentNums     int               `orm:"size(6);null;default(0)"`
	User            *User             `orm:"rel(fk)"`
	Tags            []*Tag            `orm:"rel(m2m)"`
	Work            *Work             `orm:"rel(fk)"`
	VisitorComments []*VisitorComment `orm:"reverse(many)"`
}
type Draft struct {
	Id      int    `orm:"auto;size(10)"`
	Title   string `orm:"size(100)"`
	Content string `orm:"type(text)"`

	User *User  `orm:"rel(fk)"`
	Tags []*Tag `orm:"rel(m2m)"`
	Work *Work  `orm:"rel(fk)"`
}

type Tag struct {
	Id       int
	Name     string `orm:"size(20)"`
	ParentId int
	Users    []*User  `orm:"reverse(many)"`
	Posts    []*Post  `orm:"reverse(many)"`
	Drafts   []*Draft `orm:"reverse(many)"`
}

type Work struct {
	Id    int     `orm:"auto;size(10);pk"`
	Name  string  `orm:"size(20)"`
	User  *User   `orm:"rel(fk)"`
	Posts []*Post `orm:"reverse(many)"`
}

type Visitor struct {
	Id              string            `orm:"pk;size(15)"`
	Portrait        string            `orm:"size(255)"`
	VisitorComments []*VisitorComment `orm:"reverse(many)"`
}

type VisitorComment struct {
	Id      int      `orm:"pk;auto;size(10)"`
	Content string   `orm:"type(text)"`
	Floor   int      `orm:"size(6)"`
	Time    string   `orm:"auto_now_add;type(datetime)"`
	Visitor *Visitor `orm:"rel(fk)"`
	Post    *Post    `orm:"rel(fk)"`
}

func GetNavList() (nav []*Tag) {
	o := orm.NewOrm()
	nav = []*Tag{}
	_, err := o.QueryTable("tag").Filter("parent_id", 0).All(&nav)
	if err != nil {
		//错误处理
		log.Println("GetNavList()>>", err)
	}
	return
}

func init() {
	db_user := beego.AppConfig.String("dbuer")
	db_passwd := beego.AppConfig.String("dbpass")
	db_host := beego.AppConfig.String("dbhost")
	db_port := beego.AppConfig.String("dbport")
	db := beego.AppConfig.String("db")
	orm.RegisterDataBase("default", "mysql", db_user+":"+db_passwd+"@tcp("+db_host+":"+db_port+")/"+db+"?charset=utf8")
	orm.RegisterModel(new(User), new(Post), new(Profile), new(Tag), new(Work), new(Visitor), new(VisitorComment), new(Draft))
	if beego.AppConfig.String("runmode") == "dev" {
		orm.RunSyncdb("default", false, true)
	} else {
		orm.RunSyncdb("default", false, false)
	}
}
