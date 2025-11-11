package rules

import "github.com/zeller_assignment/models"

type ThreeForTwoRule struct {
	Sku string
}

func (r *ThreeForTwoRule) Apply(items []models.Item) float64 {
	count := 0
	price := 0.0
	for _, item := range items {
		if item.Sku == r.Sku {
			count++
			price = item.Price
		}
	}
	eligible := count / 3
	regular := count % 3
	return float64(eligible*2+regular) * price
}
