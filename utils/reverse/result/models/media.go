package models

type Media struct {
	Id        *string            `gorm:"column:id;primaryKey;autoIncrement;not null"`
	DiskType  *string            `gorm:"column:disk_type;not null"` //存储类型
	SrcType   *int8              `gorm:"column:src_type;not null"`  //链接类型 1相对路径 2外链
	Src       *string            `gorm:"column:src;not null"`       //资源链接
	CreatedAt *jsontime.JsonTime `gorm:"column:created_at"`
	UpdatedAt *jsontim.JsonTime  `gorm:"column:updated_at"`
}

func (entity *Media) TableName() string {
	return "media"
}