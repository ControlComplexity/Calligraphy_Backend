package repositories

import (
	"fmt"
	"github.com/ControlComplexity/Calligraphy_Backend/model"
	"github.com/mlogclub/simple/sqls"
	"gorm.io/gorm"
)

var CarouselImageRepository = newCarouselImageRepository()

func newCarouselImageRepository() *carouselImageRepository {
	return &carouselImageRepository{}
}

type carouselImageRepository struct {
}

func (r *carouselImageRepository) Find(db *gorm.DB, cnd *sqls.Cnd) (list []model.CarouselDO) {
	cnd.Find(db, &list)
	fmt.Print("list: ", list)
	return
}
