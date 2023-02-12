package repositories

import (
	"fmt"
	"github.com/ControlComplexity/Calligraphy_Backend/model"
	"github.com/mlogclub/simple/sqls"
	"gorm.io/gorm"
)

var EssayRepository = newEssayRepository()

func newEssayRepository() *essayRepository {
	return &essayRepository{}
}

type essayRepository struct {
}

func (r *essayRepository) Find(db *gorm.DB, cnd *sqls.Cnd) (list []model.EssayDO, paging *sqls.Paging) {
	cnd.Find(db, &list)
	fmt.Print("Find list: ", list)
	count := cnd.Count(db, &model.EssayDO{})
	paging = &sqls.Paging{
		Page:  cnd.Paging.Page,
		Limit: cnd.Paging.Limit,
		Total: count,
	}
	return
}

func (r *essayRepository) FindRanked(db *gorm.DB, cnd *sqls.Cnd) (list []model.EssayDO, paging *sqls.Paging) {
	cnd.Find(db, &list)
	fmt.Print("FindRanked list: ", list)
	count := cnd.Count(db, &model.EssayDO{})
	paging = &sqls.Paging{
		Page:  cnd.Paging.Page,
		Limit: cnd.Paging.Limit,
		Total: count,
	}
	return
}

func (r *essayRepository) FindRelated(db *gorm.DB, cnd *sqls.Cnd) (list []model.EssayDO, paging *sqls.Paging) {
	cnd.Find(db, &list)
	count := cnd.Count(db, &model.EssayDO{})
	fmt.Print("FindRelated list: ", list)
	paging = &sqls.Paging{
		Page:  cnd.Paging.Page,
		Limit: cnd.Paging.Limit,
		Total: count,
	}
	return
}

func (r *essayRepository) Get(db *gorm.DB, id int64) *model.EssayDO {
	ret := &model.EssayDO{}
	if err := db.First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func (r *essayRepository) Updates(db *gorm.DB, id int64, columns map[string]interface{}) (err error) {
	err = db.Model(&model.EssayDO{}).Where("id = ?", id).Updates(columns).Error
	return
}

func (r *essayRepository) UpdateColumn(db *gorm.DB, id int64, name string, value interface{}) (err error) {
	err = db.Model(&model.EssayDO{}).Where("id = ?", id).UpdateColumn(name, value).Error
	return
}
