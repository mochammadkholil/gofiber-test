package models

type Produk struct {
	Id     string `json:"id"`
	Nama   string `json:"nama"`
	SKU    string `json:"sku"`
	Jumlah int64  `json:"jumlah"`
}
