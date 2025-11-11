package checkout

import (
	"github.com/zeller_assignment/models"
	"github.com/zeller_assignment/rules"
)

type Checkout struct {
	pricingRule map[string]rules.PricingRule
	items       []models.Item
}

func (co *Checkout) Scan(item models.Item) {
	co.items = append(co.items, item)
}

// ...existing code...
func (co *Checkout) Total() float64 {
	// Group items by SKU
	skuGroups := map[string][]models.Item{}
	for _, item := range co.items {
		skuGroups[item.Sku] = append(skuGroups[item.Sku], item)
	}

	total := 0.0
	for sku, group := range skuGroups {
		rule, ok := co.pricingRule[sku]
		if !ok {
			rule = &rules.DefaultRule{Sku: sku}
		}
		total += rule.Apply(group)
	}
	return total
}

// ...existing code...

func NewCheckout(pricingRules map[string]rules.PricingRule) *Checkout {
	return &Checkout{pricingRule: pricingRules}
}
