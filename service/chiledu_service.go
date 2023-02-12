package services

import (
	"errors"
	"github.com/ControlComplexity/Calligraphy_Backend/model"
	repositories "github.com/ControlComplexity/Calligraphy_Backend/repository"
	"github.com/mlogclub/simple/sqls"
	"github.com/mlogclub/simple/web/params"
	"time"
)

var ChileduService = newChileduService()

func newChileduService() *chileduService {
	return &chileduService{}
}

type chileduService struct{}

func (s *chileduService) Find(params *params.QueryParams) ([]model.ChileduDO, *sqls.Paging) {
	return repositories.ChileduRepository.Find(sqls.DB(), &params.Cnd)
}

func (s *chileduService) FindRanked(params *params.QueryParams) ([]model.ChileduDO, *sqls.Paging) {
	return repositories.ChileduRepository.FindRanked(sqls.DB(), &params.Cnd)
}

// FindRelated 获取相关文档
func (s *chileduService) FindRelated(params *params.QueryParams) ([]model.ChileduDO, *sqls.Paging) {
	return repositories.ChileduRepository.FindRelated(sqls.DB(), &params.Cnd)
}

func (s *chileduService) Get(id int64) *model.ChileduDO {
	return repositories.ChileduRepository.Get(sqls.DB(), id)
}

// AddHit 点击量+1
func (s *chileduService) AddHit(chileduId int64) error {
	chiledu := s.Get(chileduId)
	if chiledu == nil || chiledu.Title == "" {
		return errors.New("文章不存在")
	}
	hits := chiledu.Hits
	if err := s.Updates(chileduId, map[string]interface{}{
		"hits":       hits + 1,
		"updated_at": time.Now(),
	}); err != nil {
		return err
	}
	return nil
}

// AddLike 点击量+1
func (s *chileduService) AddLike(chileduId int64) error {
	chiledu := s.Get(chileduId)
	if chiledu == nil || chiledu.Uuid == "" {
		return errors.New("文章不存在")
	}
	like := chiledu.Like
	if err := s.Updates(chileduId, map[string]interface{}{
		"like":       like + 1,
		"updated_at": time.Now(),
	}); err != nil {
		return err
	}
	return nil
}

func (s *chileduService) Updates(chileduId int64, columns map[string]interface{}) error {
	if err := repositories.ChileduRepository.Updates(sqls.DB(), chileduId, columns); err != nil {
		return err
	}

	return nil
}

func (s *chileduService) UpdateColumn(id int64, name string, value interface{}) error {
	if err := repositories.ChileduRepository.UpdateColumn(sqls.DB(), id, name, value); err != nil {
		return err
	}
	return nil
}
