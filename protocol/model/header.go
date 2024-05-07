package model

type Header struct {
	Name   string `gorm:"type:text"`
	Value  string `gorm:"type:text"`
	Parent uint
}
