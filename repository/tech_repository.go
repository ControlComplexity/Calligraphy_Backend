package repositories

import (
	"fmt"
	"github.com/ControlComplexity/Calligraphy_Backend/model"
	"github.com/mlogclub/simple/sqls"
	"gorm.io/gorm"
)

var TechRepository = newTechRepository()

func newTechRepository() *techRepository {
	return &techRepository{}
}

type techRepository struct {
}

func (r *techRepository) Find(db *gorm.DB, cnd *sqls.Cnd) (list []model.TechDO, paging *sqls.Paging) {
	cnd.Find(db, &list)
	fmt.Print("Find list: ", list)
	count := cnd.Count(db, &model.TechDO{})
	paging = &sqls.Paging{
		Page:  cnd.Paging.Page,
		Limit: cnd.Paging.Limit,
		Total: count,
	}
	return
}

func (r *techRepository) FindRanked(db *gorm.DB, cnd *sqls.Cnd) (list []model.TechDO, paging *sqls.Paging) {
	cnd.Find(db, &list)
	fmt.Print("FindRanked list: ", list)
	count := cnd.Count(db, &model.TechDO{})
	paging = &sqls.Paging{
		Page:  cnd.Paging.Page,
		Limit: cnd.Paging.Limit,
		Total: count,
	}
	return
}

func (r *techRepository) FindRelated(db *gorm.DB, cnd *sqls.Cnd) (list []model.TechDO, paging *sqls.Paging) {
	cnd.Find(db, &list)
	count := cnd.Count(db, &model.TechDO{})
	fmt.Print("FindRelated list: ", list)
	paging = &sqls.Paging{
		Page:  cnd.Paging.Page,
		Limit: cnd.Paging.Limit,
		Total: count,
	}
	return
}

func (r *techRepository) Get(db *gorm.DB, id int64) *model.TechDO {
	ret := &model.TechDO{}
	if err := db.First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func (r *techRepository) Updates(db *gorm.DB, id int64, columns map[string]interface{}) (err error) {
	err = db.Model(&model.TechDO{}).Where("id = ?", id).Updates(columns).Error
	return
}

func (r *techRepository) UpdateColumn(db *gorm.DB, id int64, name string, value interface{}) (err error) {
	err = db.Model(&model.TechDO{}).Where("id = ?", id).UpdateColumn(name, value).Error
	return
}
