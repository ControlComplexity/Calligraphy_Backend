package model

import "time"

var Models []interface{}

type Model struct {
	Id int64 `gorm:"primaryKey;autoIncrement" json:"ID" form:"ID"`
}

const (
	essayTableName      = "essay"
	carouselTableName   = "carousel"
	exhibitionTableName = "exhibition"
	theoryTableName     = "theory"
)

type CarouselDO struct {
	Model
	Uuid      string    `gorm:"column:uuid;type:varchar(128);not null;default:'';comment:轮播图编号"`
	Title     string    `gorm:"column:title;type:string;"`
	URL       string    `gorm:"column:url;type:string;"`
	Time      time.Time `gorm:"column:time;type:time;not null;comment:理论内容"`
	Type      string    `gorm:"column:type;type:varchar(128);not null;default:'';comment:图片类型"`
	City      string    `gorm:"column:city;type:varchar(128);not null;default:'';comment:所在城市"`
	Hits      int64     `gorm:"column:hits;type:int;not null;default:0;comment:访问量"`
	Like      int64     `gorm:"column:like;type:int;not null;default:0;comment:喜欢次数"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;autoCreateTime;not null;default:CURRENT_TIMESTAMP;comment:创建时间"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp;autoUpdateTime;not null;default:CURRENT_TIMESTAMP;comment:更新时间"`
}

func (CarouselDO) TableName() string {
	return carouselTableName
}

// TheoryDO 理论
type TheoryDO struct {
	Model
	Uuid            string    `gorm:"column:uuid;type:varchar(128);not null;default:'';comment:理论编号"`
	Title           string    `gorm:"column:title;type:varchar(128);not null;default:'';comment:理论标题"`
	Abstract        string    `gorm:"column:abstract;type:varchar(300);not null;default:'';comment:理论摘要"`
	Content         string    `gorm:"column:content;type:varchar(3000);not null;default:'';comment:理论内容"`
	Time            time.Time `gorm:"column:time;type:time;not null;comment:理论时间"`
	Type            string    `gorm:"column:type;type:varchar(128);not null;default:'';comment:理论类型"`
	Image           string    `gorm:"column:image;type:varchar(128);not null;default:'';comment:图片"`
	HomeImage       string    `gorm:"column:home_image;type:varchar(128);not null;default:'';comment:首页图片"`
	Hits            int64     `gorm:"column:hits;type:int;not null;default:0;comment:访问量"`
	Like            int64     `gorm:"column:like;type:int;not null;default:0;comment:喜欢次数"`
	DisplayHomePage int32     `gorm:"column:display_home_page;type:int;not null;default:0;comment:是否在首页展示"`
	CreatedAt       time.Time `gorm:"column:created_at;type:timestamp;autoCreateTime;not null;default:CURRENT_TIMESTAMP;comment:创建时间"`
	UpdatedAt       time.Time `gorm:"column:updated_at;type:timestamp;autoUpdateTime;not null;default:CURRENT_TIMESTAMP;comment:更新时间"`
}

func (TheoryDO) TableName() string {
	return theoryTableName
}
