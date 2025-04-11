package models

type Genres struct {
	ID     uint     `gorm:"column:id_genres;primary_key" json:"id"`
	Genres string   `gorm:"column:genres" json:"genres"`
	Movies []Movies `gorm:"many2many:movie_genres;foreignKey:ID;joinForeignKey:id_genres;References:ID;joinReferences:id_movies" json:"-"`
}

func (Genres) TableName() string {
	return "genres"
}
