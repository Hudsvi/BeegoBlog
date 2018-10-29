package write

import (
	"Blog/models"
	"github.com/astaxie/beego/orm"
	log "Blog/utils/log"
)

func Publish(post *models.Post) {

	o := orm.NewOrm()
	_, err := o.Insert(post)
	if err != nil {
		log.Println("Publish()|Insert()|err:", err)
	}
	qm := o.QueryM2M(post, "Tags")
	qm.Add(post.Tags)

}
