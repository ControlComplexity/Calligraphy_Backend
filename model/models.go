package model

import "time"

var Models []interface{}

type Model struct {
	Id int64 `gorm:"primaryKey;autoIncrement" json:"ID" form:"ID"`
}

const (
	activityTableName   = "activity"
	essayTableName      = "essay"
	carouselTableName   = "carousel"
	exhibitionTableName = "exhibition"
	theoryTableName     = "theory"
	unieduTableName     = "university_education"
)

type CarouselDO struct {
	Model
	Uuid      string    `gorm:"column:uuid;type:varchar(128);not null;default:'';comment:轮播图编号"`
	Title     string    `gorm:"column:title;type:string;"`
	URL       string    `gorm:"column:url;type:string;"`
	Time      time.Time `gorm:"column:time;type:time;not null;comment:理论内容"`
	Type      string    `gorm:"column:type;type:varchar(128);not null;default:'';comment:图片类型"`
	Link      string    `gorm:"column:link;type:varchar(128);not null;default:'';comment:超链接"`
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

// EssayDO 文章
type EssayDO struct {
	Model
	Uuid            string    `gorm:"column:uuid;type:varchar(128);not null;default:'';comment:文章编号"`
	Title           string    `gorm:"column:title;type:varchar(128);not null;default:'';comment:文章标题"`
	Abstract        string    `gorm:"column:abstract;type:varchar(300);not null;default:'';comment:文章摘要"`
	Content         string    `gorm:"column:content;type:varchar(3000);not null;default:'';comment:文章内容"`
	Time            time.Time `gorm:"column:time;type:time;not null;comment:文章内容"`
	Type            string    `gorm:"column:type;type:varchar(128);not null;default:'';comment:文章类型"`
	Image           string    `gorm:"column:image;type:varchar(128);not null;default:'';comment:图片"`
	HomeImage       string    `gorm:"column:home_image;type:varchar(128);not null;default:'';comment:首页图片"`
	WriterDynasty   string    `gorm:"column:writerdynasty;type:varchar(128);not null;default:'';comment:作者的朝代"`
	ContentDynasty  string    `gorm:"column:contentdynasty;type:varchar(128);not null;default:'';comment:内容所属的朝代"`
	Calligrapher    string    `gorm:"column:calligrapher;type:varchar(128);not null;default:'';comment:书法家"`
	Hits            int64     `gorm:"column:hits;type:int;not null;default:0;comment:访问量"`
	Like            int64     `gorm:"column:like;type:int;not null;default:0;comment:喜欢次数"`
	DisplayHomePage int32     `gorm:"column:display_home_page;type:int;not null;default:0;comment:是否在首页展示"`
	CreatedAt       time.Time `gorm:"column:created_at;type:timestamp;autoCreateTime;not null;default:CURRENT_TIMESTAMP;comment:创建时间"`
	UpdatedAt       time.Time `gorm:"column:updated_at;type:timestamp;autoUpdateTime;not null;default:CURRENT_TIMESTAMP;comment:更新时间"`
}

func (EssayDO) TableName() string {
	return essayTableName
}

// UnieduDO 大学书法教育
type UnieduDO struct {
	Model
	Uuid            string    `gorm:"column:uuid;type:varchar(128);not null;default:'';comment:文章编号"`
	Title           string    `gorm:"column:title;type:varchar(128);not null;default:'';comment:文章标题"`
	Abstract        string    `gorm:"column:abstract;type:varchar(300);not null;default:'';comment:文章摘要"`
	Content         string    `gorm:"column:content;type:varchar(3000);not null;default:'';comment:文章内容"`
	Time            time.Time `gorm:"column:time;type:time;not null;comment:文章内容"`
	Type            string    `gorm:"column:type;type:varchar(128);not null;default:'';comment:文章类型"`
	Image           string    `gorm:"column:image;type:varchar(128);not null;default:'';comment:图片"`
	HomeImage       string    `gorm:"column:home_image;type:varchar(128);not null;default:'';comment:首页图片"`
	WriterDynasty   string    `gorm:"column:writerdynasty;type:varchar(128);not null;default:'';comment:作者的朝代"`
	ContentDynasty  string    `gorm:"column:contentdynasty;type:varchar(128);not null;default:'';comment:内容所属的朝代"`
	Calligrapher    string    `gorm:"column:calligrapher;type:varchar(128);not null;default:'';comment:书法家"`
	Hits            int64     `gorm:"column:hits;type:int;not null;default:0;comment:访问量"`
	Like            int64     `gorm:"column:like;type:int;not null;default:0;comment:喜欢次数"`
	DisplayHomePage int32     `gorm:"column:display_home_page;type:int;not null;default:0;comment:是否在首页展示"`
	CreatedAt       time.Time `gorm:"column:created_at;type:timestamp;autoCreateTime;not null;default:CURRENT_TIMESTAMP;comment:创建时间"`
	UpdatedAt       time.Time `gorm:"column:updated_at;type:timestamp;autoUpdateTime;not null;default:CURRENT_TIMESTAMP;comment:更新时间"`
}

func (UnieduDO) TableName() string {
	return unieduTableName
}

type ActivityDO struct {
	Model
	Name        string    `gorm:"column:name;type:varchar(50);index;not null;default:'';comment:活动名称"`
	Label       string    `gorm:"column:label;type:varchar(50);index;not null;default:'';comment:活动标签"`
	Price       float32   `gorm:"column:price;type:float;index;not null;default:0;comment:价格"`
	Location    string    `gorm:"column:location;type:varchar(50);index;not null;default:'';comment:地点"`
	Time        time.Time `gorm:"column:time;type:timestamp;autoUpdateTime;not null;default:CURRENT_TIMESTAMP;comment:更新时间"`
	SmallImg    string    `gorm:"column:small_img;type:varchar(50);index;not null;default:'';comment:小图"`
	BigImg      string    `gorm:"column:big_img;type:varchar(50);index;not null;default:'';comment:大图"`
	Description string    `gorm:"column:description;type:varchar(50);index;not null;default:'';comment:描述"`
	CreatedAt   time.Time `gorm:"column:created_at;type:timestamp;autoCreateTime;not null;default:CURRENT_TIMESTAMP;comment:创建时间"`
	UpdatedAt   time.Time `gorm:"column:updated_at;type:timestamp;autoUpdateTime;not null;default:CURRENT_TIMESTAMP;comment:更新时间"`
}

func (ActivityDO) TableName() string {
	return activityTableName
}

// ExhibitionDO 展览
type ExhibitionDO struct {
	Model
	Uuid      string    `gorm:"column:uuid;type:varchar(128);not null;default:'';comment:文章编号"`
	Title     string    `gorm:"column:title;type:varchar(128);not null;default:'';comment:文章标题"`
	Content   string    `gorm:"column:content;type:varchar(128);not null;default:'';comment:文章内容"`
	Time      time.Time `gorm:"column:time;type:time;not null;comment:文章内容"`
	Type      string    `gorm:"column:type;type:varchar(128);not null;default:'';comment:文章类型"`
	City      string    `gorm:"column:city;type:varchar(128);not null;default:'';comment:所在城市"`
	Image     string    `gorm:"column:image;type:varchar(128);not null;default:'';comment:图片"`
	Hits      int64     `gorm:"column:hits;type:int;not null;default:0;comment:访问量"`
	Like      int64     `gorm:"column:like;type:int;not null;default:0;comment:喜欢次数"`
	CreatedAt time.Time `gorm:"column:created_at;type:timestamp;autoCreateTime;not null;default:CURRENT_TIMESTAMP;comment:创建时间"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:timestamp;autoUpdateTime;not null;default:CURRENT_TIMESTAMP;comment:更新时间"`
}

func (ExhibitionDO) TableName() string {
	return exhibitionTableName
}
