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

// TestComputeLineTotal does a single entry test with no discount
func TestComputeLineTotal(t *testing.T) {
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
