package marketing_utils

import (
	. "github.com/NerdsvilleCEO/customer-ss/customer"
	. "github.com/NerdsvilleCEO/customer-ss/product"
	. "github.com/NerdsvilleCEO/customer-ss/utils"
)

type SustainabilityScore struct {
	Product  *Product
	Customer *Customer
	Score    float64
}

func (s *SustainabilityScore) FindBaseScore() float64 {
	var base float64
	if len(s.Product.Name)%2 == 0 {
		base = float64(FindVowels(s.Customer.Name)) * 1.5
	} else {
		base = float64(FindConsonants(s.Customer.Name))
	}
	return base
}

func (s *SustainabilityScore) Find() float64 {
	base := s.FindBaseScore()
	if HasCommonFactors(int64(len(s.Product.Name)), int64(len(s.Customer.Name))) {
		base *= 1.5
	}
	return base
}
