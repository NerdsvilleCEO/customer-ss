package utils

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestSpec(t *testing.T) {
	Convey("Given a string with some vowels", t, func() {
		aString := "testabc"
		anotherString := "abcdefghi"
		Convey("I should be able to find the number of vowels :D", func() {
			So(FindVowels(aString), ShouldEqual, 2)
			So(FindVowels(anotherString), ShouldEqual, 3)
		})
	})

	Convey("Given two numbers", t, func() {
		num1 := 4
		num2 := 6

		Convey("I should be able to determine if they have common factors", func() {
			So(HasCommonFactors(int64(num1), int64(num2)), ShouldEqual, true)
			So(HasCommonFactors(1, 2), ShouldEqual, false)
		})
	})
}
