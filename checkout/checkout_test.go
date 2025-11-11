package checkout

import (
	"fmt"
	"math"
	"testing"

	"github.com/zeller_assignment/models"
	"github.com/zeller_assignment/rules"
)

func almostEqual(a, b float64) bool {
	return math.Abs(a-b) < 0.01
}

func TestScenarios(t *testing.T) {
	pricingRules := map[string]rules.PricingRule{
		"atv": &rules.ThreeForTwoRule{Sku: "atv"},
		"ipd": &rules.BulkDiscountRule{Sku: "ipd", MinQuantity: 4, DiscountedPrince: 499.99},
	}

	// Scenario 1: atv, atv, atv, vga --> $249.00
	co := NewCheckout(pricingRules)
	co.Scan(models.Catalogue["atv"])
	co.Scan(models.Catalogue["atv"])
	co.Scan(models.Catalogue["atv"])
	co.Scan(models.Catalogue["vga"])
	total := co.Total()
	fmt.Println("total: ", total)
	if !almostEqual(total, 249.00) {
		t.Errorf("Scenario 1: got %.2f, want 249.00", total)
	}

	// Scenario 2: atv, ipd, ipd, atv, ipd, ipd, ipd --> $2718.95
	co = NewCheckout(pricingRules)
	co.Scan(models.Catalogue["atv"])
	co.Scan(models.Catalogue["ipd"])
	co.Scan(models.Catalogue["ipd"])
	co.Scan(models.Catalogue["atv"])
	co.Scan(models.Catalogue["ipd"])
	co.Scan(models.Catalogue["ipd"])
	co.Scan(models.Catalogue["ipd"])
	total = co.Total()
	if !almostEqual(total, 2718.95) {
		t.Errorf("Scenario 2: got %.2f, want 2718.95", total)
	}
}

func TestDefaultRule(t *testing.T) {
	co := NewCheckout(map[string]rules.PricingRule{})
	co.Scan(models.Catalogue["vga"])
	co.Scan(models.Catalogue["mbp"])
	total := co.Total()
	if !almostEqual(total, 30+1399.99) {
		t.Errorf("Default rule: got %.2f, want 1429.99", total)
	}
}
