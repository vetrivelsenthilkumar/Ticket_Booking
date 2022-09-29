package models

import "gorm.io/gorm"

type Passenger struct {
	gorm.Model

	Id       uint   `gorm:"primaryKey;not null"`
	Name     string `json:"name" gorm:"type:varchar(100)"`
	Age      int    `json:"age" gorm:"not null"`
	Email    string `json:"email" gorm:"type:varchar(100);PrimaryKey;not null"`
	Password string `json:"password" gorm:"type:varchar(100)"`
}

type Train struct {
	gorm.Model

	Train_number uint   `json:"train_number" gorm:"primaryKey"`
	From         string `json:"from" gorm:"type:varchar(100)"`
	To           string `json:"to" gorm:"type:varchar(100)"`
	Coach_number string `json:"coach_number" gorm:"not null"`
	Seat_number  int    `json:"seat_number" gorm:"not null"`
	Seat_type    string `json:"seat_type" gorm:"not null"`
}
