package transaction

func ComputeTotal(transaction *Transaction) int64 {
	return 0
}

func ComputeLineTotal(line_item *TransactionLineItem) int64 {
	product := line_item.GetProduct()
	productPrice := product.GetPrice()
	quantity := line_item.GetQuantity()

	total := productPrice * int64(quantity)

	return total
}
