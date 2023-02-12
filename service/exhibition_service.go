package services

import (
	"github.com/ControlComplexity/Calligraphy_Backend/model"
	repositories "github.com/ControlComplexity/Calligraphy_Backend/repository"
	"github.com/mlogclub/simple/sqls"
)

var ExhibitionService = newExhibitionService()

func newExhibitionService() *exhibitionService {
	return &exhibitionService{}
}

type exhibitionService struct{}

func (s *exhibitionService) Find(cnd *sqls.Cnd) ([]model.ExhibitionDO, *sqls.Paging) {
	return repositories.ExhibitionRepository.Find(sqls.DB(), cnd)
}

func (s *exhibitionService) Get(id int64) interface{} {
	return repositories.ExhibitionRepository.Get(sqls.DB(), id)
}
