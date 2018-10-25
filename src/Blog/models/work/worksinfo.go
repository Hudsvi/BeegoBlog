package work

import (
	"Blog/models"
	"github.com/astaxie/beego/orm"
	log "Blog/utils/log"
)

var (
	err error
)

type WorksInfo struct {
	Work *models.Work
	Num  int //统计每个文集的博客数量
}

func GetWorksInfo(uid string) (worksInfo []*WorksInfo) {
	o := orm.NewOrm()
	var works []*models.Work
	//查询works列表
	_, err = o.QueryTable("work").Filter("User__Phone", uid).All(&works)
	if err != nil {
		log.Println("GetWorksInfo()|works|err", err)
	}
	//这里多用一个work的长度，是用来保存默认文集“我的文集”
	for i := -1; i <len(works); i++ {
		var Num int64
		if i == -1{
			//查询总博客数量
			Num, err = o.QueryTable("post").Filter("User__Phone", uid).Count()
			if err != nil {
				log.Println("GetWorksInfo()|Num|err", err)
			}
			worksInfo = append(worksInfo, &WorksInfo{&models.Work{User: &models.User{Phone: uid}},int(Num)})
		} else {
			//查询各文集数量
			Num, err = o.QueryTable("post").Filter("Work__Id", works[i].Id).Count()
			works[i].User.Phone = uid
			worksInfo = append(worksInfo, &WorksInfo{works[i], int(Num)})
			if err != nil {
				log.Println("GetWorksInfo()|Num|err", err)
			}
		}
	}
	return
}
