package services

import (
	"github.com/ControlComplexity/Calligraphy_Backend/model"
	repositories "github.com/ControlComplexity/Calligraphy_Backend/repository"
	"github.com/mlogclub/simple/sqls"
	"github.com/mlogclub/simple/web/params"
)

var TheoryService = newTheoryService()

func newTheoryService() *theoryService {
	return &theoryService{}
}

type theoryService struct{}

func (s *theoryService) Find(params *params.QueryParams) ([]model.TheoryDO, *sqls.Paging) {
	return repositories.TheoryRepository.Find(sqls.DB(), &params.Cnd)
}
