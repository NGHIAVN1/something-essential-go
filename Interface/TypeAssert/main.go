package main

import "fmt"

func getExpenseReport(e expense) (string, float64) {
	// ?
	i, ok := e.(email)
	fmt.Println(ok)
	return i.toAddress, i.cost()

	s, ok := e.(sms)
	return s.toPhoneNumber, s.cost()

	n, ok := e.(invalid)
	return "", n.cost()
}

// don't touch below this line

type expense interface {
	cost() float64
}

type email struct {
	isSubscribed bool
	body         string
	toAddress    string
}

type sms struct {
	isSubscribed  bool
	body          string
	toPhoneNumber string
}

type invalid struct{}

func (e email) cost() float64 {
	if !e.isSubscribed {
		return float64(len(e.body)) * .05
	}
	return float64(len(e.body)) * .01
}

func (s sms) cost() float64 {
	if !s.isSubscribed {
		return float64(len(s.body)) * .1
	}
	return float64(len(s.body)) * .03
}

func (i invalid) cost() float64 {
	return 0.0
}
