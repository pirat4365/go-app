package models

type Users struct {
	Name string `gorm:"varchar(255);not null"`
	Age  int    `gorm:"int(3);not null"`
}
