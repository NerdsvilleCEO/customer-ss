package product

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSpec(t *testing.T) {
	Convey("When I create a product", t, func() {
		product := Product{
			Name: "Computer",
		}

		Convey("I should be able to get the product's name", func() {
			So(product.Name, ShouldEqual, "Computer")
		})
	})
}
