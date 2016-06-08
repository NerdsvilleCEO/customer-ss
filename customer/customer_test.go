package customer

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSpec(t *testing.T) {
	Convey("When a customer is created", t, func() {
		customer := Customer{
			Name: "Josh",
		}

		Convey("I should be able to get the name of the customer", func() {
			So(customer.Name, ShouldEqual, "Josh")
		})

	})
}
