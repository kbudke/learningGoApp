package area_test

import (
	"fmt"

	. "./area"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Area", func() {

	var (
		circleRadius, squareLength float64
	)

	BeforeEach(func() {
		fmt.Println("Running test case.")
		circleRadius = 1
		squareLength = 5
	})

	AfterEach(func() {
		fmt.Println("Test case executed.")
	})

	Describe("Testing Areas.", func() {
		Context("calculating Area of Circle with radius of 3.", func() {
			It("(Area of a circle with radius 1) will be 3.14 ", func() {
				Expect(Area.Circle.Length).toEqaul(3.14)
			})
		})

	})
})
