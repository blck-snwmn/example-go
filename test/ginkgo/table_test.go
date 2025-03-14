package ginkgo_test

import (
	"github.com/blck-snwmn/example-go/test/ginkgo"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Table-Driven Tests", func() {
	// Example of table-driven tests
	DescribeTable("Calculator.Add",
		func(a, b, expected int) {
			calc := &ginkgo.Calculator{}
			result := calc.Add(a, b)
			Expect(result).To(Equal(expected))
		},
		Entry("adding positive numbers", 2, 3, 5),
		Entry("adding with a negative number", -2, 5, 3),
		Entry("adding two negative numbers", -2, -3, -5),
		Entry("adding with zero", 0, 5, 5),
	)

	DescribeTable("Calculator.Subtract",
		func(a, b, expected int) {
			calc := &ginkgo.Calculator{}
			result := calc.Subtract(a, b)
			Expect(result).To(Equal(expected))
		},
		Entry("subtracting positive numbers", 5, 3, 2),
		Entry("subtracting a negative number", 5, -3, 8),
		Entry("subtraction resulting in a negative", 3, 5, -2),
		Entry("subtracting zero", 5, 0, 5),
	)

	DescribeTable("Calculator.Multiply",
		func(a, b, expected int) {
			calc := &ginkgo.Calculator{}
			result := calc.Multiply(a, b)
			Expect(result).To(Equal(expected))
		},
		Entry("multiplying positive numbers", 2, 3, 6),
		Entry("multiplying with a negative number", 2, -3, -6),
		Entry("multiplying two negative numbers", -2, -3, 6),
		Entry("multiplying with zero", 0, 5, 0),
	)

	// Table-driven test with multiple inputs and outputs
	type divideTestCase struct {
		a, b     int
		expected int
		hasError bool
		errorMsg string
	}

	DescribeTable("Calculator.Divide",
		func(tc divideTestCase) {
			calc := &ginkgo.Calculator{}
			result, err := calc.Divide(tc.a, tc.b)

			if tc.hasError {
				Expect(err).To(HaveOccurred())
				if tc.errorMsg != "" {
					Expect(err.Error()).To(Equal(tc.errorMsg))
				}
			} else {
				Expect(err).NotTo(HaveOccurred())
				Expect(result).To(Equal(tc.expected))
			}
		},
		Entry("dividing positive numbers", divideTestCase{
			a:        6,
			b:        3,
			expected: 2,
			hasError: false,
		}),
		Entry("division with truncation", divideTestCase{
			a:        5,
			b:        2,
			expected: 2, // Integer division truncates decimal part
			hasError: false,
		}),
		Entry("dividing with a negative number", divideTestCase{
			a:        -6,
			b:        3,
			expected: -2,
			hasError: false,
		}),
		Entry("dividing by zero", divideTestCase{
			a:        5,
			b:        0,
			hasError: true,
			errorMsg: "cannot divide by zero",
		}),
	)

	// Table-driven test with labels
	DescribeTable("UserService.ValidateAge",
		func(age int, expected bool) {
			service := &ginkgo.UserService{}
			result := service.ValidateAge(age)
			Expect(result).To(Equal(expected))
		},
		Entry("minimum valid age", 0, true),
		Entry("normal valid age", 30, true),
		Entry("maximum valid age", 120, true),
		Entry("negative invalid age", -1, false),
		Entry("above maximum invalid age", 121, false),
	)
})
