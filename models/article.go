package models

import "time"

type Article struct {
	Id int
	Title string
	Description string
	Author string
	Status uint8
	CreatedAt time.Time
	CreatedAtStr string  `orm:"-"`
}
