package marketing_utils

import (
	. "github.com/NerdsvilleCEO/customer-ss/customer"
	. "github.com/NerdsvilleCEO/customer-ss/product"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSpec(t *testing.T) {
	Convey("Given a product which shares common factors besides 1 with a customer", t, func() {
		product := &Product{
			Name: "Pool",
		}
		customer := &Customer{
			Name: "Josh",
		}

		ss := &SustainabilityScore{
			Customer: customer,
			Product:  product,
		}

		Convey("When find is called, our sustainability score should be increased by a total of 50%", func() {
			So(ss.Find(), ShouldEqual, 2.25)

			ss.Product.Name = "ab"
			ss.Customer.Name = "a"
			So(ss.Find(), ShouldEqual, 1.5)

			ss.Product.Name = "ab"
			ss.Customer.Name = "abc"
			So(ss.Find(), ShouldEqual, 1.5)
		})
	})

	Convey("Given a product with an odd length", t, func() {
		product := &Product{
			Name: "Phone",
		}

		Convey("The base sustainability score should be 1 times the number of consonants in the name of the customer", func() {
			customer := &Customer{
				Name: "Josh",
			}
			ss := SustainabilityScore{
				Customer: customer,
				Product:  product,
			}
			So(ss.FindBaseScore(), ShouldEqual, 3)
		})
	})

	Convey("Given a product with an even length", t, func() {
		product := &Product{
			Name: "Computer",
		}

		Convey("The base sustainability score should be 1.5 times the number of vowels in the name of the customer ", func() {
			customer := &Customer{
				Name: "Josh",
			}
			ss := SustainabilityScore{
				Customer: customer,
				Product:  product,
			}

			So(ss.FindBaseScore(), ShouldEqual, 1.5)

			ss.Customer.Name = "Andrew"
			So(ss.FindBaseScore(), ShouldEqual, 3)

		})
	})
}
