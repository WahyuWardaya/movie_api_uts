package models

type Actors struct {
	ID     uint     `gorm:"column:id_actors;primary_key" json:"id"`
	Name   string   `gorm:"column:name_actors" json:"name_actors"`
	Movies []Movies `gorm:"many2many:movie_actors;foreignKey:ID;joinForeignKey:id_actors;References:ID;joinReferences:id_movies" json:"-"`
}

func (Actors) TableName() string {
	return "actors"
}
