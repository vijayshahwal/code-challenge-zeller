# Zeller Checkout System

A flexible checkout system implementation in Go that supports various pricing rules and promotional offers.

## 📋 Table of Contents

- [System Overview](#system-overview)
- [Architecture Flow](#architecture-flow)
- [Data Flow](#data-flow)
- [Execution Flow](#execution-flow)
- [Pricing Rule Flow](#pricing-rule-flow)
- [Testing Flow](#testing-flow)

## System Overview

This checkout system implements a flexible rule-based pricing architecture that processes items through configurable pricing strategies.

## Architecture Flow

```
┌─────────────────────────────────────────────────────────────┐
│                     Application Entry                       │
│                        (main.go)                            │
└────────────────────────┬────────────────────────────────────┘
                         │
                         │ 1. Initialize pricing rules
                         │ 2. Create checkout instance
                         │
                         ▼
┌─────────────────────────────────────────────────────────────┐
│                    Checkout System                          │
│                  (checkout/checkout.go)                     │
│                                                             │
│  ┌─────────────┐    ┌─────────────┐    ┌─────────────┐   │
│  │  Scan()     │───▶│  Store Item │───▶│  Total()    │   │
│  │  Add items  │    │  in memory  │    │  Calculate  │   │
│  └─────────────┘    └─────────────┘    └──────┬──────┘   │
└────────────────────────────────────────────────┼──────────┘
                                                  │
                         ┌────────────────────────┘
                         │
                         ▼
┌─────────────────────────────────────────────────────────────┐
│                   Pricing Rule Engine                       │
│                    (rules/*.go)                             │
│                                                             │
│  ┌──────────────┐  ┌──────────────┐  ┌─────────────────┐  │
│  │ DefaultRule  │  │3-for-2 Rule  │  │ BulkDiscount    │  │
│  │   Apply()    │  │   Apply()    │  │    Rule         │  │
│  └──────────────┘  └──────────────┘  └─────────────────┘  │
└─────────────────────────────────────────────────────────────┘
                         │
                         ▼
                    Return Total
```

## Data Flow

### 1. Item Catalog Flow
```
┌──────────────────┐
│  models/item.go  │
│                  │
│  ┌────────────┐  │
│  │ Catalogue  │  │──┐
│  │  - ipd     │  │  │
│  │  - mbp     │  │  │ Provides static
│  │  - atv     │  │  │ product data
│  │  - vga     │  │  │
│  └────────────┘  │  │
└──────────────────┘  │
                      │
                      ▼
              Used by main.go
              and checkout system
```

### 2. Checkout Flow
```
Start
  │
  ▼
Initialize Pricing Rules
  │
  ├─── "atv" → ThreeForTwoRule
  ├─── "ipd" → BulkDiscountRule
  └─── Others → DefaultRule (auto)
  │
  ▼
Create Checkout Instance
  │
  ▼
Scan Items (one by one)
  │
  ├─── Item 1 → Added to items[]
  ├─── Item 2 → Added to items[]
  └─── Item N → Added to items[]
  │
  ▼
Call Total()
  │
  ▼
End
```

## Execution Flow

### Main Application Flow

```
main.go execution:

1. Define Pricing Rules
   └─ Map SKU to Rule implementation
      ├─ "atv" → ThreeForTwoRule
      └─ "ipd" → BulkDiscountRule

2. Create Checkout
   └─ NewCheckout(pricingRules)

3. Scan Items
   └─ For each item:
      ├─ Get from Catalogue
      └─ co.Scan(item)

4. Calculate Total
   └─ co.Total()
      ├─ Groups items by SKU
      ├─ Applies rules per group
      └─ Returns sum
```

### Checkout Total Calculation Flow

```
Total() Method Flow:

Input: All scanned items
  │
  ▼
Group items by SKU
  │
  ├─ SKU: "atv" → [Item, Item, Item]
  ├─ SKU: "ipd" → [Item, Item, Item, Item, Item]
  └─ SKU: "vga" → [Item]
  │
  ▼
For each SKU group:
  │
  ├─ Check if pricing rule exists
  │  │
  │  ├─ YES → Use custom rule
  │  └─ NO → Use DefaultRule
  │
  ▼
Apply rule to group
  │
  └─ rule.Apply(items[])
     │
     └─ Returns subtotal
  │
  ▼
Sum all subtotals
  │
  ▼
Return final total
```

## Pricing Rule Flow

### Rule Selection and Application

```
                    ┌─────────────────┐
                    │  Items by SKU   │
                    └────────┬────────┘
                             │
                             ▼
                    ┌─────────────────┐
                    │ Rule exists for │
                    │   this SKU?     │
                    └────────┬────────┘
                             │
                 ┌───────────┴───────────┐
                 │                       │
                YES                     NO
                 │                       │
                 ▼                       ▼
        ┌────────────────┐      ┌────────────────┐
        │  Apply Custom  │      │ Apply Default  │
        │     Rule       │      │     Rule       │
        └────────┬───────┘      └────────┬───────┘
                 │                       │
                 └───────────┬───────────┘
                             │
                             ▼
                    ┌────────────────┐
                    │  Calculate &   │
                    │ Return Subtotal│
                    └────────────────┘
```

### Three-for-Two Rule Flow

```
Input: [Item, Item, Item, Item, Item]
  │
  ▼
Count items: 5
  │
  ▼
Calculate eligible groups: 5 ÷ 3 = 1 (with remainder 2)
  │
  ├─ Free items: 1 × 1 = 1
  ├─ Paid items: (1 × 2) + 2 = 4
  │
  ▼
Total = 4 × price
  │
  ▼
Return total
```

### Bulk Discount Rule Flow

```
Input: [Item, Item, Item, Item, Item]
  │
  ▼
Count items: 5
  │
  ▼
Check quantity: 5 > MinQuantity (4)?
  │
  ├─ YES → Use discounted price
  │    └─ Total = 5 × $499.99
  │
  └─ NO → Use regular price
       └─ Total = count × regular_price
  │
  ▼
Return total
```

### Default Rule Flow

```
Input: [Item, Item]
  │
  ▼
Count items: 2
  │
  ▼
Calculate: count × price
  │
  ▼
Return total
```

## Testing Flow

### Test Execution Flow

```
Test Suite Execution:

1. Setup Phase
   └─ Define pricing rules
      ├─ ThreeForTwoRule for "atv"
      └─ BulkDiscountRule for "ipd"

2. Test Scenario 1
   └─ Create new checkout
   └─ Scan: atv, atv, atv, vga
   └─ Assert: total ≈ $249.00

3. Test Scenario 2
   └─ Create new checkout
   └─ Scan: atv, ipd, ipd, atv, ipd, ipd, ipd
   └─ Assert: total ≈ $2718.95

4. Test Scenario 3 (Default Rule)
   └─ Create checkout (no rules)
   └─ Scan: vga, mbp
   └─ Assert: total ≈ $1429.99
```

### Test Validation Flow

```
For each test:
  │
  ▼
Calculate expected total manually
  │
  ▼
Run checkout.Total()
  │
  ▼
Compare actual vs expected
  │
  ├─ Difference < $0.01 → PASS
  └─ Otherwise → FAIL
```

## Complete Transaction Flow Example

```
Scenario: Buy 3 Apple TVs and 1 VGA adapter

Step 1: Initialize
   pricingRules = {"atv": ThreeForTwoRule}
   checkout = NewCheckout(pricingRules)

Step 2: Scan Items
   checkout.Scan(Catalogue["atv"])  → items = [atv]
   checkout.Scan(Catalogue["atv"])  → items = [atv, atv]
   checkout.Scan(Catalogue["atv"])  → items = [atv, atv, atv]
   checkout.Scan(Catalogue["vga"])  → items = [atv, atv, atv, vga]

Step 3: Calculate Total
   checkout.Total()
   │
   ├─ Group by SKU:
   │  ├─ "atv": [atv, atv, atv]
   │  └─ "vga": [vga]
   │
   ├─ Apply Rules:
   │  ├─ "atv" → ThreeForTwoRule.Apply()
   │  │  └─ 3 items → pay for 2
   │  │  └─ Subtotal: $219.00
   │  │
   │  └─ "vga" → DefaultRule.Apply()
   │     └─ 1 item → regular price
   │     └─ Subtotal: $30.00
   │
   └─ Sum: $219.00 + $30.00 = $249.00

Step 4: Return Result
   → $249.00
```

## Module Dependencies Flow

```
main.go
  │
  ├─── imports checkout
  ├─── imports models
  └─── imports rules
           │
           │
checkout/checkout.go
  │
  ├─── imports models
  └─── imports rules
           │
           │
rules/*.go
  │
  └─── imports models
           │
           │
models/item.go
  │
  └─── (no dependencies)
```

---

**Quick Reference:**
- Entry point: [`main.go`](main.go:1)
- Checkout logic: [`checkout/checkout.go`](checkout/checkout.go:1)
- Pricing rules: [`rules/`](rules/)
- Product catalog: [`models/item.go`](models/item.go:1)
- Tests: [`checkout/checkout_test.go`](checkout/checkout_test.go:1)
