package models

type Roles struct {
	Id string `json:"id" gorm:"column:id;int unsigned;PRI;AUTO_INCREMENT;not null"` 
	Role int32 `json:"role" gorm:"column:role;int;not null"` 
}
func (entity *Roles) TableName() string {
	return "roles"
}