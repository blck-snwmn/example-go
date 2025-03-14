package ginkgo_test

import (
	"github.com/blck-snwmn/example-go/test/ginkgo"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Calculator", func() {
	var calculator *ginkgo.Calculator

	// Executed before each test
	BeforeEach(func() {
		calculator = &ginkgo.Calculator{}
	})

	// Addition tests
	Describe("Add", func() {
		Context("adding positive numbers", func() {
			It("should add correctly", func() {
				result := calculator.Add(2, 3)
				Expect(result).To(Equal(5))
			})
		})

		Context("adding with negative numbers", func() {
			It("should add a negative and a positive number", func() {
				result := calculator.Add(-2, 5)
				Expect(result).To(Equal(3))
			})

			It("should add two negative numbers", func() {
				result := calculator.Add(-2, -3)
				Expect(result).To(Equal(-5))
			})
		})
	})

	// Subtraction tests
	Describe("Subtract", func() {
		Context("subtracting positive numbers", func() {
			It("should subtract correctly", func() {
				result := calculator.Subtract(5, 3)
				Expect(result).To(Equal(2))
			})
		})

		Context("subtracting with negative numbers", func() {
			It("should add when subtracting a negative from a positive", func() {
				result := calculator.Subtract(5, -3)
				Expect(result).To(Equal(8))
			})
		})
	})

	// Multiplication tests
	Describe("Multiply", func() {
		Context("multiplying positive numbers", func() {
			It("should multiply correctly", func() {
				result := calculator.Multiply(2, 3)
				Expect(result).To(Equal(6))
			})
		})

		Context("multiplying with negative numbers", func() {
			It("should result in a negative when multiplying a positive and a negative", func() {
				result := calculator.Multiply(2, -3)
				Expect(result).To(Equal(-6))
			})

			It("should result in a positive when multiplying two negatives", func() {
				result := calculator.Multiply(-2, -3)
				Expect(result).To(Equal(6))
			})
		})

		Context("multiplying with zero", func() {
			It("should result in zero when either number is zero", func() {
				result := calculator.Multiply(0, 5)
				Expect(result).To(Equal(0))

				result = calculator.Multiply(5, 0)
				Expect(result).To(Equal(0))
			})
		})
	})

	// Division tests
	Describe("Divide", func() {
		Context("dividing positive numbers", func() {
			It("should divide correctly", func() {
				result, err := calculator.Divide(6, 3)
				Expect(err).NotTo(HaveOccurred())
				Expect(result).To(Equal(2))
			})

			It("should truncate decimal results in integer division", func() {
				result, err := calculator.Divide(5, 2)
				Expect(err).NotTo(HaveOccurred())
				Expect(result).To(Equal(2)) // 2.5 -> 2 (integer division)
			})
		})

		Context("dividing by zero", func() {
			It("should return an error when dividing by zero", func() {
				_, err := calculator.Divide(5, 0)
				Expect(err).To(Equal(ginkgo.ErrDivideByZero))
			})
		})

		Context("dividing with negative numbers", func() {
			It("should result in a negative when dividing by a negative", func() {
				result, err := calculator.Divide(6, -3)
				Expect(err).NotTo(HaveOccurred())
				Expect(result).To(Equal(-2))
			})

			It("should calculate correctly when the dividend is negative", func() {
				result, err := calculator.Divide(-6, 3)
				Expect(err).NotTo(HaveOccurred())
				Expect(result).To(Equal(-2))
			})
		})
	})
})
