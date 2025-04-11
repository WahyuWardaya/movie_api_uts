package models

import "time"

type Movies struct {
	ID          uint        `gorm:"column:id_movies;primary_key" json:"id"`
	Name        string      `gorm:"column:name_movies" json:"name_movies"`
	Directors 	[]Directors `gorm:"many2many:movie_directors;foreignKey:ID;joinForeignKey:id_movies;References:ID;joinReferences:id_directors" json:"directors"`
	Actors 		[]Actors 	`gorm:"many2many:movie_actors;foreignKey:ID;joinForeignKey:id_movies;References:ID;joinReferences:id_actors" json:"actors"`
	Genres 		[]Genres 	`gorm:"many2many:movie_genres;foreignKey:ID;joinForeignKey:id_movies;References:ID;joinReferences:id_genres" json:"genres"`
	Description string      `gorm:"column:description_movies" json:"description_movies"`
	Rating      float32     `gorm:"column:rating" json:"rating"`
	RealeseDate time.Time   `gorm:"column:realese_date" json:"realese_date"`
	Duration    int         `gorm:"column:duration" json:"duration"`
	Lenguage    string      `gorm:"column:lenguage" json:"lenguage"`
	CreatedAt   time.Time   `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time   `gorm:"column:updated_at" json:"updated_at"`
}