package repositories

import (
	"fmt"
	"github.com/ControlComplexity/Calligraphy_Backend/model"
	"github.com/mlogclub/simple/sqls"
	"gorm.io/gorm"
)

var TheoryRepository = newTheoryRepository()

func newTheoryRepository() *theoryRepository {
	return &theoryRepository{}
}

type theoryRepository struct {
}

// Find 返回书法理论的id和标题
func (r *theoryRepository) Find(db *gorm.DB, cnd *sqls.Cnd) (list []model.TheoryDO, paging *sqls.Paging) {
	cnd.SelectCols = []string{"id", "title", "image"}
	cnd.Find(db, &list)
	fmt.Print("Find list: ", list)
	count := cnd.Count(db, &model.TheoryDO{})
	paging = &sqls.Paging{
		Page:  cnd.Paging.Page,
		Limit: cnd.Paging.Limit,
		Total: count,
	}
	return
}

func (r *theoryRepository) FindRanked(db *gorm.DB, cnd *sqls.Cnd) (list []model.TheoryDO, paging *sqls.Paging) {
	cnd.Find(db, &list)
	fmt.Print("FindRanked list: ", list)
	count := cnd.Count(db, &model.TheoryDO{})
	paging = &sqls.Paging{
		Page:  cnd.Paging.Page,
		Limit: cnd.Paging.Limit,
		Total: count,
	}
	return
}

func (r *theoryRepository) FindRelated(db *gorm.DB, cnd *sqls.Cnd) (list []model.TheoryDO, paging *sqls.Paging) {
	cnd.Find(db, &list)
	count := cnd.Count(db, &model.TheoryDO{})
	fmt.Print("FindRelated list: ", list)
	paging = &sqls.Paging{
		Page:  cnd.Paging.Page,
		Limit: cnd.Paging.Limit,
		Total: count,
	}
	return
}

func (r *theoryRepository) Get(db *gorm.DB, id int64) *model.TheoryDO {
	ret := &model.TheoryDO{}
	if err := db.First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func (r *theoryRepository) Updates(db *gorm.DB, id int64, columns map[string]interface{}) (err error) {
	err = db.Model(&model.TheoryDO{}).Where("id = ?", id).Updates(columns).Error
	return
}

func (r *theoryRepository) UpdateColumn(db *gorm.DB, id int64, name string, value interface{}) (err error) {
	err = db.Model(&model.TheoryDO{}).Where("id = ?", id).UpdateColumn(name, value).Error
	return
}
