package transaction

import (
	"fmt"
	"reflect"
	"testing"

	money "google.golang.org/genproto/googleapis/type/money"
)

func TesComputeTotal(t *testing.T) {

	fmt.Println(money.Money{
		CurrencyCode: "USD",
		Units:        1,
		Nanos:        0,
	})
}

// TestComputeLineTotalNoDiscount does 1 product with no discount
func TestComputeLineTotalNoDiscount(t *testing.T) {
	expected := int64(49990)
	actual := ComputeLineTotal(&TransactionLineItem{
		Product: &Product{
			Price: 49990,
		},
		Quantity: 1,
		Discount: &Discount{
			DiscountType: Discount_DISCOUNTTYPE_NONE,
		},
	})
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Should return %v instead of %v", expected, actual)
	}
}

// TestComputeLineTotalDiscountPercent 1 product with percent discount
func TestComputeLineTotalDiscountPercent(t *testing.T) {
	expected := int64(44991)
	actual := ComputeLineTotal(&TransactionLineItem{
		Product: &Product{
			Price: 49990,
		},
		Quantity: 1,
		Discount: &Discount{
			DiscountType:  Discount_DISCOUNTTYPE_PERCENT,
			DiscountValue: .10,
		},
	})
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Should return %v instead of %v", expected, actual)
	}
}

// TestComputeLineTotalDiscountPercentWithCondition 1 product with percent discount with condition
func TestComputeLineTotalDiscountPercentWithCondition(t *testing.T) {
	// 2 for the price of 1
	expected := int64(219000)
	actual := ComputeLineTotal(&TransactionLineItem{
		Product: &Product{
			Sku:   "A304SD",
			Name:  "AlexaSpeaker",
			Price: 109500,
		},
		Quantity: 3,
		Discount: &Discount{
			DiscountType:      Discount_DISCOUNTTYPE_PERCENT,
			DiscountValue:     1.0 / 3.0,
			DiscountCondition: "q>=3",
		},
	})
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("Should return %v instead of %v", expected, actual)
	}
}
