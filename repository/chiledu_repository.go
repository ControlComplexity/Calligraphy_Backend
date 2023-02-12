package repositories

import (
	"fmt"
	"github.com/ControlComplexity/Calligraphy_Backend/model"
	"github.com/mlogclub/simple/sqls"
	"gorm.io/gorm"
)

var ChileduRepository = newChileduRepository()

func newChileduRepository() *chileduRepository {
	return &chileduRepository{}
}

type chileduRepository struct {
}

func (r *chileduRepository) Find(db *gorm.DB, cnd *sqls.Cnd) (list []model.ChileduDO, paging *sqls.Paging) {
	cnd.Find(db, &list)
	fmt.Print("Find list: ", list)
	count := cnd.Count(db, &model.ChileduDO{})
	paging = &sqls.Paging{
		Page:  cnd.Paging.Page,
		Limit: cnd.Paging.Limit,
		Total: count,
	}
	return
}

func (r *chileduRepository) FindRanked(db *gorm.DB, cnd *sqls.Cnd) (list []model.ChileduDO, paging *sqls.Paging) {
	cnd.Find(db, &list)
	fmt.Print("FindRanked list: ", list)
	count := cnd.Count(db, &model.ChileduDO{})
	paging = &sqls.Paging{
		Page:  cnd.Paging.Page,
		Limit: cnd.Paging.Limit,
		Total: count,
	}
	return
}

func (r *chileduRepository) FindRelated(db *gorm.DB, cnd *sqls.Cnd) (list []model.ChileduDO, paging *sqls.Paging) {
	cnd.Find(db, &list)
	count := cnd.Count(db, &model.ChileduDO{})
	fmt.Print("FindRelated list: ", list)
	paging = &sqls.Paging{
		Page:  cnd.Paging.Page,
		Limit: cnd.Paging.Limit,
		Total: count,
	}
	return
}

func (r *chileduRepository) Get(db *gorm.DB, id int64) *model.ChileduDO {
	ret := &model.ChileduDO{}
	if err := db.First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func (r *chileduRepository) Updates(db *gorm.DB, id int64, columns map[string]interface{}) (err error) {
	err = db.Model(&model.ChileduDO{}).Where("id = ?", id).Updates(columns).Error
	return
}

func (r *chileduRepository) UpdateColumn(db *gorm.DB, id int64, name string, value interface{}) (err error) {
	err = db.Model(&model.ChileduDO{}).Where("id = ?", id).UpdateColumn(name, value).Error
	return
}
