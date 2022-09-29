package models

type Passenger struct {
	Id       uint   `gorm:"primaryKey"`
	Name     string `json:"name" gorm:"type:varchar(100);not null"`
	Age      int    `json:"age" gorm:"not null"`
	Email    string `json:"email" gorm:"primaryKey"`
	Password string `json:"password" gorm:"type:varchar(100);not null"`
}

type Train struct {
	Train_number uint   `gorm:"primaryKey" gorm:"not null"`
	From         string `json:"from" gorm:"type:varchar(100);not null"`
	To           string `json:"to" gorm:"type:varchar(100);not null"`
	Coach_number string `json:"coach_number" gorm:"not null"`
	Seat_number  int    `json:"Seat_number" gorm:"not null"`
	Seat_type    string `json:"Seat_type" gorm:"not null"`
}
