package write

import (
	"Blog/models"
	"github.com/astaxie/beego/orm"
	log "Blog/utils/log"
)

func SaveDraft(draft *models.Draft) {
	o := orm.NewOrm()
	_, err := o.Insert(draft)
	if err != nil {
		log.Println("SaveDraft()|Insert()|err:", err)
	}
	qm := o.QueryM2M(draft, "Tags") //多对多关系中间表插入数据，
	qm.Add(draft.Tags)
}
