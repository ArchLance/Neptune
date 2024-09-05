package model

type User struct {
	Id       int    `gorm:"type:int;primary_key;AUTO_INCREMENT"`
	UserName string `gorm:"type:varchar(64);not null;unique"`
	Avatar   string `gorm:"type:varchar(64);unique"`
	Account  string `gorm:"type:varchar(64);not null"`
	Email    string `gorm:"type:varchar(64);not null;unique"`
	Password string `gorm:"type:varchar(64);not null"`
	Role     string `gorm:"type:varchar(20);not null"`
}
