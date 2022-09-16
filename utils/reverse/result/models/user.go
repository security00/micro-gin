package models

type User struct {
	Id *string `gorm:"column:id;primaryKey;autoIncrement;not null"` 
	Name *string `gorm:"column:name;not null"` 
}
func (entity *User) TableName() string {
	return "user"
}