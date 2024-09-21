package model

type Poc struct {
	Id                int    `gorm:"type:int;primary_key;AUTO_INCREMENT"`
	VulnerabilityName string `gorm:"type:varchar(64)"`
	PocName           string `gorm:"type:varchar(64);not null;unique"`
	AppName           string `gorm:"type:varchar(64)"`
	VulnerabilityType string `gorm:"type:varchar(64)"`
	AddTime           string `gorm:"type:varchar(64)"`
}
