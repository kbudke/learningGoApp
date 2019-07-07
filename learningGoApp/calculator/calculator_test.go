package calculator_test

import (
	"fmt"

	. "./calculator"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Calculator", func() {

	var (
		num1, num2 int32
	)

	BeforeEach(func() {
		fmt.Println("Running test case.")
		num1 = 6
		num2 = 4
	})

	AfterEach(func() {
		fmt.Println("Test case executed.")
	})

	Describe("Testing Calculator.", func() {
		Context("calculating sum of (6, 4).", func() {
			It("Sum of 6 and 4 should be 10", func() {
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
