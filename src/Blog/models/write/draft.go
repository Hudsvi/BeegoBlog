package write

import (
	"Blog/models"
	"github.com/astaxie/beego/orm"
)

func SaveDraft(draft *models.Draft) {
	o := orm.NewOrm()
	o.Insert(draft)
	qm := o.QueryM2M(draft, "Tags") //多对多关系中间表插入数据，
	qm.Add(draft.Tags)
}
