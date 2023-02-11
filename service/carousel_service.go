package services

import (
	"github.com/ControlComplexity/Calligraphy_Backend/model"
	repositories "github.com/ControlComplexity/Calligraphy_Backend/repository"
	"github.com/mlogclub/simple/sqls"
)

var CarouselService = newCarouselService()

func newCarouselService() *carouselService {
	return &carouselService{}
}

type carouselService struct{}

func (s *carouselService) Find(cnd *sqls.Cnd) []model.CarouselDO {
	return repositories.CarouselImageRepository.Find(sqls.DB(), cnd)
}
