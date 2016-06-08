package customer

import (
	. "github.com/NerdsvilleCEO/customer-ss/product"
)

type Customer struct {
	Name    string
	Product *Product
}
