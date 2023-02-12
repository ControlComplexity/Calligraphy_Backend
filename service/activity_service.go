package services

import (
	"fmt"
	"github.com/ControlComplexity/Calligraphy_Backend/model"
	repositories "github.com/ControlComplexity/Calligraphy_Backend/repository"
	"github.com/mlogclub/simple/sqls"
	"github.com/mlogclub/simple/web/params"
)

var ActivityService = newActivityService()

func newActivityService() *activityService {
	return &activityService{}
}

type activityService struct{}

func (s *activityService) Find(params *params.QueryParams) ([]model.ActivityDO, *sqls.Paging) {
	fmt.Println("111")
	return repositories.ActivityRepository.Find(sqls.DB(), &params.Cnd)
}

func (s *activityService) Get(id int64) interface{} {
	return repositories.ActivityRepository.Get(sqls.DB(), id)
}
