package repositories

import (
	"fmt"
	"github.com/ControlComplexity/Calligraphy_Backend/model"
	"github.com/mlogclub/simple/sqls"
	"gorm.io/gorm"
)

var ActivityRepository = newActivityRepository()

func newActivityRepository() *activityRepository {
	return &activityRepository{}
}

type activityRepository struct {
}

func (r *activityRepository) Find(db *gorm.DB, cnd *sqls.Cnd) (list []model.ActivityDO, paging *sqls.Paging) {
	cnd.Find(db, &list)
	fmt.Println("222")
	count := cnd.Count(db, &model.ActivityDO{})
	fmt.Println("333", cnd.Paging)
	paging = &sqls.Paging{
		Page:  cnd.Paging.Page,
		Limit: cnd.Paging.Limit,
		Total: count,
	}
	fmt.Print("list: ", list)
	return
}

func (r *activityRepository) Get(db *gorm.DB, id int64) *model.ActivityDO {
	ret := &model.ActivityDO{}
	if err := db.First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}
