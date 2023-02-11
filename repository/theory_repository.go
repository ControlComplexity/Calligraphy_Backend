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

func (r *theoryRepository) Find(db *gorm.DB, cnd *sqls.Cnd) (list []model.TheoryDO, paging *sqls.Paging) {
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
