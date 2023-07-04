package types

type GoShortInput struct {
	Redirect string `json:"short"`
}

type Goshort struct {
	ID       uint64 `json:"id" gorm:"primaryKey"`
	Redirect string `json:"redirect" gorm:"not null"`
	Goshort  string `json:"short" gorm:"unique;not null"`
	Clicked  uint64 `json:"clicked"`
}
