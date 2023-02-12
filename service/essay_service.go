package services

import (
	"errors"
	"github.com/ControlComplexity/Calligraphy_Backend/model"
	repositories "github.com/ControlComplexity/Calligraphy_Backend/repository"
	"github.com/mlogclub/simple/sqls"
	"github.com/mlogclub/simple/web/params"
	"time"
)

var EssayService = newEssayService()

func newEssayService() *essayService {
	return &essayService{}
}

type essayService struct{}

func (s *essayService) Find(params *params.QueryParams) ([]model.EssayDO, *sqls.Paging) {
	return repositories.EssayRepository.Find(sqls.DB(), &params.Cnd)
}

func (s *essayService) FindRanked(params *params.QueryParams) ([]model.EssayDO, *sqls.Paging) {
	return repositories.EssayRepository.FindRanked(sqls.DB(), &params.Cnd)
}

// FindRelated 获取相关文档
func (s *essayService) FindRelated(params *params.QueryParams) ([]model.EssayDO, *sqls.Paging) {
	return repositories.EssayRepository.FindRelated(sqls.DB(), &params.Cnd)
}

func (s *essayService) Get(id int64) *model.EssayDO {
	return repositories.EssayRepository.Get(sqls.DB(), id)
}

// AddHit 点击量+1
func (s *essayService) AddHit(essayId int64) error {
	essay := s.Get(essayId)
	if essay == nil || essay.Title == "" {
		return errors.New("文章不存在")
	}
	hits := essay.Hits
	if err := s.Updates(essayId, map[string]interface{}{
		"hits":       hits + 1,
		"updated_at": time.Now(),
	}); err != nil {
		return err
	}
	return nil
}

// AddLike 点击量+1
func (s *essayService) AddLike(essayId int64) error {
	essay := s.Get(essayId)
	if essay == nil || essay.Uuid == "" {
		return errors.New("文章不存在")
	}
	like := essay.Like
	if err := s.Updates(essayId, map[string]interface{}{
		"like":       like + 1,
		"updated_at": time.Now(),
	}); err != nil {
		return err
	}
	return nil
}

func (s *essayService) Updates(essayId int64, columns map[string]interface{}) error {
	if err := repositories.EssayRepository.Updates(sqls.DB(), essayId, columns); err != nil {
		return err
	}

	return nil
}

func (s *essayService) UpdateColumn(id int64, name string, value interface{}) error {
	if err := repositories.EssayRepository.UpdateColumn(sqls.DB(), id, name, value); err != nil {
		return err
	}
	return nil
}
