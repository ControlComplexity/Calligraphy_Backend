package services

import (
	"github.com/ControlComplexity/Calligraphy_Backend/model"
	repositories "github.com/ControlComplexity/Calligraphy_Backend/repository"
	"github.com/mlogclub/simple/sqls"
	"github.com/mlogclub/simple/web/params"
)

var SearchService = newSearchService()

func newSearchService() *searchService {
	return &searchService{}
}
type searchService struct{}

func (s *searchService) Find(params *params.QueryParams) ([]model.ActivityDO, *sqls.Paging) {
	return repositories.SearchRepository.Find(sqls.DB(), &params.Cnd)
}

func (s *searchService) Top() []model.Theory {
	return repositories.SearchRepository.Top()
}