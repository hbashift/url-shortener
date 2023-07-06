package model

type Url struct {
	ID  uint64 `gorm:"type:bigserial; primarykey"`
	Url string `gorm:"unique; not null"`
}
