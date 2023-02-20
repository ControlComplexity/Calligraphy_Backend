package repositories

import (
	"fmt"
	"github.com/ControlComplexity/Calligraphy_Backend/model"
	"github.com/mlogclub/simple/sqls"
	"gorm.io/gorm"
)

var NewsRepository = newNewsRepository()

func newNewsRepository() *newsRepository {
	return &newsRepository{}
}

type newsRepository struct {
}

func (r *newsRepository) Find(db *gorm.DB, cnd *sqls.Cnd) (list []model.NewsDO, paging *sqls.Paging) {
	cnd.Find(db, &list)
	fmt.Print("Find list: ", list)
	count := cnd.Count(db, &model.NewsDO{})
	paging = &sqls.Paging{
		Page:  cnd.Paging.Page,
		Limit: cnd.Paging.Limit,
		Total: count,
	}
	return
}

func (r *newsRepository) FindRanked(db *gorm.DB, cnd *sqls.Cnd) (list []model.NewsDO, paging *sqls.Paging) {
	cnd.Find(db, &list)
	fmt.Print("FindRanked list: ", list)
	count := cnd.Count(db, &model.NewsDO{})
	paging = &sqls.Paging{
		Page:  cnd.Paging.Page,
		Limit: cnd.Paging.Limit,
		Total: count,
	}
	return
}

func (r *newsRepository) FindRelated(db *gorm.DB, cnd *sqls.Cnd) (list []model.NewsDO, paging *sqls.Paging) {
	cnd.Find(db, &list)
	count := cnd.Count(db, &model.NewsDO{})
	fmt.Print("FindRelated list: ", list)
	paging = &sqls.Paging{
		Page:  cnd.Paging.Page,
		Limit: cnd.Paging.Limit,
		Total: count,
	}
	return
}

func (r *newsRepository) Get(db *gorm.DB, id int64) *model.NewsDO {
	ret := &model.NewsDO{}
	if err := db.First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func (r *newsRepository) Updates(db *gorm.DB, id int64, columns map[string]interface{}) (err error) {
	err = db.Model(&model.NewsDO{}).Where("id = ?", id).Updates(columns).Error
	return
}

func (r *newsRepository) UpdateColumn(db *gorm.DB, id int64, name string, value interface{}) (err error) {
	err = db.Model(&model.NewsDO{}).Where("id = ?", id).UpdateColumn(name, value).Error
	return
}
