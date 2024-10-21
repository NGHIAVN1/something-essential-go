package main

func calculateFinalBill(costPerMessage float64, numMessages int) float64 {
	// ?
	discount := calculateBaseBill(costPerMessage, numMessages) * calculateDiscountRate(numMessages)
	return calculateBaseBill(costPerMessage, numMessages) - discount
}

func calculateDiscountRate(messagesSent int) float64 {
	// ?
	if a := messagesSent; a > 1000 && a <= 5000 {
		return 0.10
	} else if a > 5000 {
		return 0.20
	} else {
		return 0.0
	}
}

// don't touch below this line

func calculateBaseBill(costPerMessage float64, messagesSent int) float64 {
	return costPerMessage * float64(messagesSent)
}
