package services

import (
	"errors"
	"github.com/ControlComplexity/Calligraphy_Backend/model"
	repositories "github.com/ControlComplexity/Calligraphy_Backend/repository"
	"github.com/mlogclub/simple/sqls"
	"github.com/mlogclub/simple/web/params"
	"time"
)

var NewsService = newNewsService()

func newNewsService() *newsService {
	return &newsService{}
}

type newsService struct{}

func (s *newsService) Find(params *params.QueryParams) ([]model.NewsDO, *sqls.Paging) {
	return repositories.NewsRepository.Find(sqls.DB(), &params.Cnd)
}

func (s *newsService) FindRanked(params *params.QueryParams) ([]model.NewsDO, *sqls.Paging) {
	return repositories.NewsRepository.FindRanked(sqls.DB(), &params.Cnd)
}

// FindRelated 获取相关文档
func (s *newsService) FindRelated(params *params.QueryParams) ([]model.NewsDO, *sqls.Paging) {
	return repositories.NewsRepository.FindRelated(sqls.DB(), &params.Cnd)
}

func (s *newsService) Get(id int64) *model.NewsDO {
	return repositories.NewsRepository.Get(sqls.DB(), id)
}

// AddHit 点击量+1
func (s *newsService) AddHit(newsId int64) error {
	news := s.Get(newsId)
	if news == nil || news.Title == "" {
		return errors.New("文章不存在")
	}
	hits := news.Hits
	if err := s.Updates(newsId, map[string]interface{}{
		"hits":       hits + 1,
		"updated_at": time.Now(),
	}); err != nil {
		return err
	}
	return nil
}

// AddLike 点击量+1
func (s *newsService) AddLike(newsId int64) error {
	news := s.Get(newsId)
	if news == nil || news.Uuid == "" {
		return errors.New("文章不存在")
	}
	like := news.Like
	if err := s.Updates(newsId, map[string]interface{}{
		"like":       like + 1,
		"updated_at": time.Now(),
	}); err != nil {
		return err
	}
	return nil
}

func (s *newsService) Updates(newsId int64, columns map[string]interface{}) error {
	if err := repositories.NewsRepository.Updates(sqls.DB(), newsId, columns); err != nil {
		return err
	}

	return nil
}

func (s *newsService) UpdateColumn(id int64, name string, value interface{}) error {
	if err := repositories.NewsRepository.UpdateColumn(sqls.DB(), id, name, value); err != nil {
		return err
	}
	return nil
}
