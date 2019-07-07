package area_test

import "fmt"

var _ = Describe("Area", func() {

	var (
		circleRadius, squareLength float64
	)

	BeforeEach(func() {
		fmt.Println("Running test case.")
		circleRadius = 1
		squareLength = 4
	})

	AfterEach(func() {
		fmt.Println("Test case executed.")
	})

	Describe("Testing Areas.", func() {
		Context("calculating Area of Circle with radius of 3.", func() {
			It("Area of a circle with radius 1 should be 3.14 ", func() {
				Expect(calculator.Sum(num1, num2)).To(Equal((int32(10))))
			})
		})

		Context("calculating subtraction of (6, 4).", func() {
			It("Subtraction of 6 and 4 should be -2", func() {
				Expect(calculator.Sub(num1, num2)).To(Equal((int32(-2))))
			})
		})

		Context("calculating multiplication of (6, 4).", func() {
			It("Multiplication of 6 and 4 should be 24", func() {
				Expect(calculator.Multi(num1, num2)).To(Equal((int32(24))))
			})
		})

		Context("calculating square root of (9).", func() {
			It("Square root of 9 should be 3.", func() {
				Expect(calculator.Sqrt(9)).To(Equal((float64(3))))
			})
		})
	})
})
