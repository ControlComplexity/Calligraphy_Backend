package services

import (
	"errors"
	"github.com/ControlComplexity/Calligraphy_Backend/model"
	repositories "github.com/ControlComplexity/Calligraphy_Backend/repository"
	"github.com/mlogclub/simple/sqls"
	"github.com/mlogclub/simple/web/params"
	"time"
)

var UnieduService = newUnieduService()

func newUnieduService() *unieduService {
	return &unieduService{}
}

type unieduService struct{}

func (s *unieduService) Find(params *params.QueryParams) ([]model.UnieduDO, *sqls.Paging) {
	return repositories.UnieduRepository.Find(sqls.DB(), &params.Cnd)
}

func (s *unieduService) FindRanked(params *params.QueryParams) ([]model.UnieduDO, *sqls.Paging) {
	return repositories.UnieduRepository.FindRanked(sqls.DB(), &params.Cnd)
}

// FindRelated 获取相关文档
func (s *unieduService) FindRelated(params *params.QueryParams) ([]model.UnieduDO, *sqls.Paging) {
	return repositories.UnieduRepository.FindRelated(sqls.DB(), &params.Cnd)
}

func (s *unieduService) Get(id int64) *model.UnieduDO {
	return repositories.UnieduRepository.Get(sqls.DB(), id)
}

// AddHit 点击量+1
func (s *unieduService) AddHit(unieduId int64) error {
	uniedu := s.Get(unieduId)
	if uniedu == nil || uniedu.Title == "" {
		return errors.New("文章不存在")
	}
	hits := uniedu.Hits
	if err := s.Updates(unieduId, map[string]interface{}{
		"hits":       hits + 1,
		"updated_at": time.Now(),
	}); err != nil {
		return err
	}
	return nil
}

// AddLike 点击量+1
func (s *unieduService) AddLike(unieduId int64) error {
	uniedu := s.Get(unieduId)
	if uniedu == nil || uniedu.Uuid == "" {
		return errors.New("文章不存在")
	}
	like := uniedu.Like
	if err := s.Updates(unieduId, map[string]interface{}{
		"like":       like + 1,
		"updated_at": time.Now(),
	}); err != nil {
		return err
	}
	return nil
}

func (s *unieduService) Updates(unieduId int64, columns map[string]interface{}) error {
	if err := repositories.UnieduRepository.Updates(sqls.DB(), unieduId, columns); err != nil {
		return err
	}

	return nil
}

func (s *unieduService) UpdateColumn(id int64, name string, value interface{}) error {
	if err := repositories.UnieduRepository.UpdateColumn(sqls.DB(), id, name, value); err != nil {
		return err
	}
	return nil
}
