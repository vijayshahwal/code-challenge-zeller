package main

import (
	"fmt"

	"github.com/zeller_assignment/checkout"
	"github.com/zeller_assignment/models"
	"github.com/zeller_assignment/rules"
)

func main() {
	pricingRules := map[string]rules.PricingRule{
		"atv": &rules.ThreeForTwoRule{Sku: "atv"},
		"ipd": &rules.BulkDiscountRule{Sku: "ipd", MinQuantity: 4, DiscountedPrince: 499.99},
	}

	// Scenario 1: atv, atv, atv, vga --> $249.00
	co := checkout.NewCheckout(pricingRules)
	co.Scan(models.Catalogue["atv"])
	co.Scan(models.Catalogue["atv"])
	co.Scan(models.Catalogue["atv"])
	co.Scan(models.Catalogue["vga"])
	fmt.Println()
	total := co.Total()
	fmt.Println("total: ", total)
}
