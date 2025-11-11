package rules

import "github.com/zeller_assignment/models"

type BulkDiscountRule struct {
	Sku              string
	MinQuantity      int
	DiscountedPrince float64
}

func (r *BulkDiscountRule) Apply(items []models.Item) float64 {
	count := 0
	price := 0.0
	for _, item := range items {
		if item.Sku == r.Sku {
			count++
			price = item.Price
		}
	}
	if count > r.MinQuantity {
		return float64(count) * r.DiscountedPrince
	}
	return float64(count) * price
}
