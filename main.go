package main

import (
	"fmt"

	"github.com/shocknawe/test-shopingcart/pb"
	"google.golang.org/genproto/googleapis/type/money"
)

func main() {
	transaction := pb.Transaction{
		Transactions: []*pb.TransactionLineItem{
			&pb.TransactionLineItem{
				Product: &pb.Product{
					Sku:  "120P90",
					Name: "GoogleHome",
					Price: &money.Money{
						CurrencyCode: "USD",
						Units:        49,
						Nanos:        990000000,
					},
				},
				Quantity: 1,
				Discount: &pb.Discount{
					DiscountType: pb.Discount_DISCOUNTTYPE_NONE,
				},
			},
		},
	}
	fmt.Println(transaction)
}
