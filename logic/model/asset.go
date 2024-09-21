package model

type Asset struct {
	AssetId     int    `gorm:"type:int;primary_key;AUTO_INCREMENT"`
	AssetName   string `gorm:"type:varchar(256);not null"`
	ProductName string `gorm:"type:varchar(256);not null"`
	IpList      string `gorm:"type:text"`
	IpNumber    int    `gorm:"type:int"`
}
