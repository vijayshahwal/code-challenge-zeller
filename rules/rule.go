package rules

import "github.com/zeller_assignment/models"

type PricingRule interface {
	Apply(items []models.Item) float64
}
