package model

type Manager struct {
	Id       int    `gorm:"type:int;primary_key;AUTO_INCREMENT"`
	Level    int    `grom:"type:int;not null"`
	Name     string `gorm:"type:varchar(32);not null"`
	Account  string `gorm:"type:varchar(32);not null"`
	Password string `gorm:"type:varchar(128);not null"`
}
