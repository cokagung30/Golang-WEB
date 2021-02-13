package entity

type Product struct {
	ID    int
	Name  string
	Price int
	Stock int
}

func (product Product) StockStatus() string {
	var status string
	if product.Stock < 3 {
		status = "Stock hampir habis"
	} else if product.Stock < 10 {
		status = "Stok terbatas"
	} else {
		status = "Stok Tersedia"
	}

	return status
}
