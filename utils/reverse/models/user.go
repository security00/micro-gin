package models

type User struct {
	Id string `json:"id" gorm:"column:id;int unsigned;PRI;AUTO_INCREMENT;not null"` 
	Name string `json:"name" gorm:"column:name;varchar(255);not null"` 
}
func (entity *User) TableName() string {
	return "user"
}