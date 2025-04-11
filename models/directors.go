package models

type Directors struct {
	ID     uint     `gorm:"column:id_directors;primary_key" json:"id"`
	Name   string   `gorm:"column:name_director" json:"name_director"`
	Movies []Movies `gorm:"many2many:movie_directors;foreignKey:ID;joinForeignKey:id_directors;References:ID;joinReferences:id_movies" json:"-"`
}

func (Directors) TableName() string {
	return "directors"
}
