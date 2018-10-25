package article

import (
	"github.com/astaxie/beego/orm"
	log "Blog/utils/log"
)

func ReadingNumAdd(articleID int){
o:=orm.NewOrm()
_,err=o.Raw("update post set reading_nums=reading_nums+1 where id=?",articleID).Exec()
if err!=nil{
	log.Println("ReadingNumAdd()|err:",err)
}
}
