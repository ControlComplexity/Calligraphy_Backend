package services

import (
	"errors"
	"github.com/ControlComplexity/Calligraphy_Backend/model"
	repositories "github.com/ControlComplexity/Calligraphy_Backend/repository"
	"github.com/mlogclub/simple/sqls"
	"github.com/mlogclub/simple/web/params"
	"time"
)

var TechService = newTechService()

func newTechService() *techService {
	return &techService{}
}

type techService struct{}

func (s *techService) Find(params *params.QueryParams) ([]model.TechDO, *sqls.Paging) {
	return repositories.TechRepository.Find(sqls.DB(), &params.Cnd)
}

func (s *techService) FindRanked(params *params.QueryParams) ([]model.TechDO, *sqls.Paging) {
	return repositories.TechRepository.FindRanked(sqls.DB(), &params.Cnd)
}

// FindRelated 获取相关文档
func (s *techService) FindRelated(params *params.QueryParams) ([]model.TechDO, *sqls.Paging) {
	return repositories.TechRepository.FindRelated(sqls.DB(), &params.Cnd)
}

func (s *techService) Get(id int64) *model.TechDO {
	return repositories.TechRepository.Get(sqls.DB(), id)
}

// AddHit 点击量+1
func (s *techService) AddHit(techId int64) error {
	tech := s.Get(techId)
	if tech == nil || tech.Title == "" {
		return errors.New("文章不存在")
	}
	hits := tech.Hits
	if err := s.Updates(techId, map[string]interface{}{
		"hits":       hits + 1,
		"updated_at": time.Now(),
	}); err != nil {
		return err
	}
	return nil
}

// AddLike 点击量+1
func (s *techService) AddLike(techId int64) error {
	tech := s.Get(techId)
	if tech == nil || tech.Uuid == "" {
		return errors.New("文章不存在")
	}
	like := tech.Like
	if err := s.Updates(techId, map[string]interface{}{
		"like":       like + 1,
		"updated_at": time.Now(),
	}); err != nil {
		return err
	}
	return nil
}

func (s *techService) Updates(techId int64, columns map[string]interface{}) error {
	if err := repositories.TechRepository.Updates(sqls.DB(), techId, columns); err != nil {
		return err
	}

	return nil
}

func (s *techService) UpdateColumn(id int64, name string, value interface{}) error {
	if err := repositories.TechRepository.UpdateColumn(sqls.DB(), id, name, value); err != nil {
		return err
	}
	return nil
}
