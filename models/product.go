package models

type Product struct {
	Id    int64   `gorm:"primaryKey" json:"id" form:"id"`
	Nama  string  `json:"nama" form:"nama"`
	Stok  int     `json:"stok" form:"stok"`
	Harga float64 `json:"harga" form:"harga"`
}
