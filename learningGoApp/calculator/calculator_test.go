package calculator_test

import (
	"fmt"

	. "practice/goPractice/learningGoApp/learningGoApp/calculator"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Calculator", func() {

	var (
		num1, num2 int
	)

	BeforeEach(func() {
		fmt.Println("Running test case.")
		num1 = 6
		num2 = 4
	})

	AfterEach(func() {
		fmt.Println("Test case executed.")
	})

	// Describe("Testing ", func() {
	Context("calculating sum of (6, 4).", func() {
		It("Sum of 6 and 4 should be 10", func() {
			Expect(Add(num1, num2)).To(Equal(10))
		})
	})

	Context("calculating subtraction of (6, 4).", func() {
		It("Subtraction of 6 and 4 should be -2", func() {
			Expect(Subtract(num1, num2)).To(Equal(2))
		})
	})

	Context("calculating multiplication of (6, 4).", func() {
		It("Multiplication of 6 and 4 should be 24", func() {
			Expect(Multiply(num1, num2)).To(Equal(24))
		})
	})

	Context("calculating square root of (9).", func() {
		It("Square root of 9 should be 3.", func() {
			Expect(Sqrt(9)).To(Equal((float64(3))))
		})
	})

})
