package models

type Roles struct {
	ID   uint   `gorm:"column:id_roles;primary_key" json:"id"`
	Name string `gorm:"column:roles_name" json:"name"`
}

func (Roles) TableName() string {
	return "roles"
}
