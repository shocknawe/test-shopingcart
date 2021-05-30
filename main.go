package main

import (
	"fmt"

	"github.com/shocknawe/test-shopingcart/proto"
	"google.golang.org/genproto/googleapis/type/money"
)

func main() {
	transaction := proto.Transaction{
		Transactions: []*proto.TransactionLineItem{
			&proto.TransactionLineItem{
				Product: &proto.Product{
					Sku:  "120P90",
					Name: "GoogleHome",
					Price: &money.Money{
						CurrencyCode: "USD",
						Units:        49,
						Nanos:        990000000,
					},
				},
				Quantity: 1,
				Discount: &proto.Discount{
					DiscountType: proto.Discount_DISCOUNTTYPE_NONE,
				},
			},
		},
	}
	fmt.Println(transaction)
}
