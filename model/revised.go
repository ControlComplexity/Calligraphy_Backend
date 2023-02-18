package model

// Theory 理论
type Theory struct {
	Model
	Title           string    `gorm:"column:title;type:varchar(128);not null;default:'';comment:理论标题"`
	Abstract        string    `gorm:"column:abstract;type:varchar(300);not null;default:'';comment:理论摘要"`
	Time            string    `gorm:"column:time;type:time;not null;comment:理论时间"`
	Image           string    `gorm:"column:image;type:varchar(128);not null;default:'';comment:图片"`
}

// Essay 文章
type Essay struct {
	Model
	Title           string    `gorm:"column:title;type:varchar(128);not null;default:'';comment:理论标题"`
	Abstract        string    `gorm:"column:abstract;type:varchar(300);not null;default:'';comment:理论摘要"`
	Time            string    `gorm:"column:time;type:time;not null;comment:理论时间"`
	Image           string    `gorm:"column:image;type:varchar(128);not null;default:'';comment:图片"`
}

// Tech 科技与书法
type Tech struct {
	Model
	Title           string    `gorm:"column:title;type:varchar(128);not null;default:'';comment:标题"`
	Abstract        string    `gorm:"column:abstract;type:varchar(300);not null;default:'';comment:摘要"`
	Time            string    `gorm:"column:time;type:time;not null;comment:理论时间"`
	Image           string    `gorm:"column:image;type:varchar(128);not null;default:'';comment:图片"`
}