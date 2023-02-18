package repositories

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ControlComplexity/Calligraphy_Backend/constant"
	"github.com/ControlComplexity/Calligraphy_Backend/model"
	"github.com/go-redis/redis/v8"
	"github.com/mlogclub/simple/sqls"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var SearchRepository = newSearchRepository()

func newSearchRepository() *searchRepository {
	return &searchRepository{}
}

type searchRepository struct {}

func (r *searchRepository) Find(db *gorm.DB, cnd *sqls.Cnd) (list []model.ActivityDO, paging *sqls.Paging) {
	cnd.Find(db, &list)
	fmt.Print("Find list: ", list)
	count := cnd.Count(db, &model.ActivityDO{})
	paging = &sqls.Paging{
		Page:  cnd.Paging.Page,
		Limit: cnd.Paging.Limit,
		Total: count,
	}
	return
}

func (r *searchRepository) Top() (list []model.Theory) {
	ret, err := GetRankList(context.Background())
	if err!= nil{
		return nil
	}
	if len(ret)==0{
		return nil
	}
	res := make([]model.Theory, 0)
	for _, r := range ret{
		str := r.Member.(string)
		var theory model.Theory
		e := json.Unmarshal([]byte(str), &theory)
		if e!= nil{
			log.Error("failed to unmarshall theory: "+ e.Error())
			continue
		}
		res = append(res, theory)
	}
	return res
}

func init() {
	log, _ = zap.NewProduction()
}

var log *zap.Logger

// GetRankList 取排行获榜列表
func GetRankList(ctx context.Context) ([]redis.Z, error) {
	rds := getRedis()
	ret, err := rds.ZRevRangeWithScores(ctx, constant.CALLIGRAPHY_THEORY_ZSET_KEY, 0, 9).Result()
	if err != nil {
		log.Error("failed to get ranked list: " + err.Error())
		return nil, fmt.Errorf("GetRankList  error : %v", err)
	}
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}
	return ret, err
}
