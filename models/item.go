package models

type Item struct {
	Sku   string
	Name  string
	Price float64
}

var Catalogue = map[string]Item{
	"ipd": {Sku: "ipd", Name: "Super iPad", Price: 549.99},
	"mbp": {Sku: "mbp", Name: "MacBook Pro", Price: 1399.99},
	"atv": {Sku: "atv", Name: "Apple TV", Price: 109.50},
	"vga": {Sku: "vga", Name: "VGA adapter", Price: 30.00},
}
