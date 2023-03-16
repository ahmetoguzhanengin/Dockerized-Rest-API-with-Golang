package models

type Card struct {
	ID          uint64 `json:"id" gorm:"primary_key; not null"`
	Title       string `json:"title" gorm:"not null"`
	Description string `json:"description" gorm:"not null"`
	ImageUrl    string `json:"imageurl" gorm:"not null"`
	Createddate string `json:"createddate"`
}
