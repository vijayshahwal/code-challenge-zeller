package rules

import "github.com/zeller_assignment/models"

type DefaultRule struct {
	Sku string
}

func (r *DefaultRule) Apply(items []models.Item) float64 {
	count := 0
	price := 0.0
	for _, item := range items {
		if item.Sku == r.Sku {
			count++
			price = item.Price
		}
	}
	return float64(count) * price
}
