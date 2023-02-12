package repositories

import (
	"fmt"
	"github.com/ControlComplexity/Calligraphy_Backend/model"
	"github.com/mlogclub/simple/sqls"
	"gorm.io/gorm"
)

var UnieduRepository = newUnieduRepository()

func newUnieduRepository() *unieduRepository {
	return &unieduRepository{}
}

type unieduRepository struct {
}

func (r *unieduRepository) Find(db *gorm.DB, cnd *sqls.Cnd) (list []model.UnieduDO, paging *sqls.Paging) {
	cnd.Find(db, &list)
	fmt.Print("Find list: ", list)
	count := cnd.Count(db, &model.UnieduDO{})
	paging = &sqls.Paging{
		Page:  cnd.Paging.Page,
		Limit: cnd.Paging.Limit,
		Total: count,
	}
	return
}

func (r *unieduRepository) FindRanked(db *gorm.DB, cnd *sqls.Cnd) (list []model.UnieduDO, paging *sqls.Paging) {
	cnd.Find(db, &list)
	fmt.Print("FindRanked list: ", list)
	count := cnd.Count(db, &model.UnieduDO{})
	paging = &sqls.Paging{
		Page:  cnd.Paging.Page,
		Limit: cnd.Paging.Limit,
		Total: count,
	}
	return
}

func (r *unieduRepository) FindRelated(db *gorm.DB, cnd *sqls.Cnd) (list []model.UnieduDO, paging *sqls.Paging) {
	cnd.Find(db, &list)
	count := cnd.Count(db, &model.UnieduDO{})
	fmt.Print("FindRelated list: ", list)
	paging = &sqls.Paging{
		Page:  cnd.Paging.Page,
		Limit: cnd.Paging.Limit,
		Total: count,
	}
	return
}

func (r *unieduRepository) Get(db *gorm.DB, id int64) *model.UnieduDO {
	ret := &model.UnieduDO{}
	if err := db.First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func (r *unieduRepository) Updates(db *gorm.DB, id int64, columns map[string]interface{}) (err error) {
	err = db.Model(&model.UnieduDO{}).Where("id = ?", id).Updates(columns).Error
	return
}

func (r *unieduRepository) UpdateColumn(db *gorm.DB, id int64, name string, value interface{}) (err error) {
	err = db.Model(&model.UnieduDO{}).Where("id = ?", id).UpdateColumn(name, value).Error
	return
}
