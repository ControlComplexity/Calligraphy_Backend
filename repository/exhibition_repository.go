package repositories

import (
	"fmt"
	"github.com/ControlComplexity/Calligraphy_Backend/model"
	"github.com/mlogclub/simple/sqls"
	"gorm.io/gorm"
)

var ExhibitionRepository = newExhibitionRepository()

func newExhibitionRepository() *exhibitionRepository {
	return &exhibitionRepository{}
}

type exhibitionRepository struct {
}

func (r *exhibitionRepository) Find(db *gorm.DB, cnd *sqls.Cnd) (list []model.ExhibitionDO, paging *sqls.Paging) {
	cnd.Find(db, &list)
	count := cnd.Count(db, &model.ExhibitionDO{})
	paging = &sqls.Paging{
		Page:  cnd.Paging.Page,
		Limit: cnd.Paging.Limit,
		Total: count,
	}
	fmt.Print("list: ", list)
	return
}

func (r *exhibitionRepository) Get(db *gorm.DB, id int64) *model.ExhibitionDO {
	ret := &model.ExhibitionDO{}
	if err := db.First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}
