package models

type Roles struct {
	Id *string `gorm:"column:id;primaryKey;autoIncrement;not null"` 
	Role *int32 `gorm:"column:role;not null"` 
}
func (entity *Roles) TableName() string {
	return "roles"
}