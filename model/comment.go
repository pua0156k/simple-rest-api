package model

import (
	"time"
)

type Comment struct {
	Uuid       string    `json:"uuid"`
	ParentId   string    `json:"parentid"`
	Comment    string    `json:"comment"`
	Author     string    `json:"author"`
	UpdateTime time.Time `json:"update"`
	Favorite   int       `json:"favorite"`
}
