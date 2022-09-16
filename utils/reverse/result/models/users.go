package models

type Users struct {
	Id *string `gorm:"column:id;primaryKey;autoIncrement;not null"` 
	Name *string `gorm:"column:name;not null"` //用户名称
	Mobile *string `gorm:"column:mobile;not null"` //用户手机号
	Password *string `gorm:"column:password;not null"` //用户密码
	CreatedAt *jsontime.JsonTime `gorm:"column:created_at"` 
	UpdatedAt *jsontime.JsonTime `gorm:"column:updated_at"` 
	DeletedAt *jsontime.JsonTime `gorm:"column:deleted_at"` 
}
func (entity *Users) TableName() string {
	return "users"
}