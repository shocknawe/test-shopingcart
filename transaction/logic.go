package transaction

import (
	"log"
	"math"
	"regexp"
	"strconv"
)

func ComputeTotal(transaction *Transaction) int64 {
	return 0
}

func ComputeLineTotal(line_item *TransactionLineItem) int64 {
	var total int64
	product := line_item.GetProduct()
	quantity := line_item.GetQuantity()
	discount := line_item.GetDiscount()
	condition := discount.GetDiscountCondition()
	lineTotal := product.GetPrice() * int64(quantity)

	didIPassCondition := false

	if condition == "" {
		// no condition for discount
		didIPassCondition = true
	} else {
		qCondition := regexp.MustCompile(`^q>=(\d+)/?`)
		matches := qCondition.FindStringSubmatch(condition)
		if len(matches) == 2 {
			qConditionNum, err := strconv.Atoi(matches[1])
			if err != nil {
				log.Printf("Couldn't convert to in %v\n", matches[1])
			}
			if quantity >= int32(qConditionNum) {
				didIPassCondition = true
			}
		}
	}

	switch discount.GetDiscountType() {
	case Discount_DISCOUNTTYPE_NONE:
		total = lineTotal
	case Discount_DISCOUNTTYPE_PERCENT:
		if !didIPassCondition {
			total = lineTotal
			break
		}
		discountedValue := float32(lineTotal) * (1.0 - discount.GetDiscountValue())
		discountedValue = float32(math.Round(float64(discountedValue)))
		total = int64(discountedValue)
	}

	return total
}
