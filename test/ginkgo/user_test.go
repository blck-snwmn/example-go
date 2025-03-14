package ginkgo_test

import (
	"context"
	"errors"
	"time"

	"github.com/blck-snwmn/example-go/test/ginkgo"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// MockUserRepository is a mock repository for testing
type MockUserRepository struct {
	users map[int]*ginkgo.User
	// Counter to record the number of calls
	FindByIDCalled int
	SaveCalled     int
	// Flags to simulate errors
	ShouldFailFindByID bool
	ShouldFailSave     bool
}

// NewMockUserRepository creates a new mock repository
func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{
		users: make(map[int]*ginkgo.User),
	}
}

// FindByID simulates user lookup by ID
func (m *MockUserRepository) FindByID(ctx context.Context, id int) (*ginkgo.User, error) {
	m.FindByIDCalled++

	if m.ShouldFailFindByID {
		return nil, errors.New("simulated database error")
	}

	user, exists := m.users[id]
	if !exists {
		return nil, errors.New("user not found")
	}

	return user, nil
}

// Save simulates saving a user
func (m *MockUserRepository) Save(ctx context.Context, user *ginkgo.User) error {
	m.SaveCalled++

	if m.ShouldFailSave {
		return errors.New("simulated database error")
	}

	m.users[user.ID] = user
	return nil
}

// AddUser is a helper method to add a user for testing
func (m *MockUserRepository) AddUser(user *ginkgo.User) {
	m.users[user.ID] = user
}

var _ = Describe("UserService", func() {
	var (
		mockRepo    *MockUserRepository
		userService *ginkgo.UserService
		ctx         context.Context
		cancelFunc  context.CancelFunc
	)

	// Executed before each test
	BeforeEach(func() {
		mockRepo = NewMockUserRepository()
		userService = ginkgo.NewUserService(mockRepo)
		ctx, cancelFunc = context.WithTimeout(context.Background(), 1*time.Second)
	})

	// Executed after each test
	AfterEach(func() {
		cancelFunc() // Cancel the context to clean up resources
	})

	Describe("GetUser", func() {
		Context("when the user exists", func() {
			BeforeEach(func() {
				// Add a test user
				mockRepo.AddUser(&ginkgo.User{
					ID:       1,
					Name:     "John Doe",
					Email:    "john@example.com",
					Age:      30,
					IsActive: true,
				})
			})

			It("should retrieve the user correctly", func() {
				user, err := userService.GetUser(ctx, 1)

				Expect(err).NotTo(HaveOccurred())
				Expect(user).NotTo(BeNil())
				Expect(user.ID).To(Equal(1))
				Expect(user.Name).To(Equal("John Doe"))
				Expect(user.Email).To(Equal("john@example.com"))
				Expect(user.Age).To(Equal(30))
				Expect(user.IsActive).To(BeTrue())

				// Verify that the repository method was called
				Expect(mockRepo.FindByIDCalled).To(Equal(1))
			})
		})

		Context("when the user does not exist", func() {
			It("should return an error", func() {
				user, err := userService.GetUser(ctx, 999)

				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("user not found"))
				Expect(user).To(BeNil())
			})
		})

		Context("when a database error occurs", func() {
			BeforeEach(func() {
				mockRepo.ShouldFailFindByID = true
			})

			It("should return an error", func() {
				user, err := userService.GetUser(ctx, 1)

				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("simulated database error"))
				Expect(user).To(BeNil())
			})
		})

		Context("when the context is canceled", func() {
			It("should return an error", func() {
				// Cancel the context
				cancelFunc()

				user, err := userService.GetUser(ctx, 1)

				Expect(err).To(HaveOccurred())
				Expect(err).To(MatchError(context.Canceled))
				Expect(user).To(BeNil())
			})
		})
	})

	Describe("ActivateUser", func() {
		Context("when the user is inactive", func() {
			BeforeEach(func() {
				// Add an inactive user
				mockRepo.AddUser(&ginkgo.User{
					ID:       1,
					Name:     "John Doe",
					Email:    "john@example.com",
					Age:      30,
					IsActive: false,
				})
			})

			It("should activate the user", func() {
				err := userService.ActivateUser(ctx, 1)

				Expect(err).NotTo(HaveOccurred())

				// Verify that the user is now active
				user, _ := mockRepo.FindByID(ctx, 1)
				Expect(user.IsActive).To(BeTrue())

				// Verify method call counts
				Expect(mockRepo.FindByIDCalled).To(Equal(2)) // Once in ActivateUser and once in verification
				Expect(mockRepo.SaveCalled).To(Equal(1))
			})
		})

		Context("when the user is already active", func() {
			BeforeEach(func() {
				// Add an active user
				mockRepo.AddUser(&ginkgo.User{
					ID:       1,
					Name:     "John Doe",
					Email:    "john@example.com",
					Age:      30,
					IsActive: true,
				})
			})

			It("should return an error", func() {
				err := userService.ActivateUser(ctx, 1)

				Expect(err).To(HaveOccurred())
				Expect(err.Error()).To(Equal("user is already active"))

				// Verify that Save method was not called
				Expect(mockRepo.SaveCalled).To(Equal(0))
			})
		})
	})

	Describe("ValidateAge", func() {
		Context("when the age is valid", func() {
			It("should validate age 0 as valid", func() {
				Expect(userService.ValidateAge(0)).To(BeTrue())
			})

			It("should validate normal ages as valid", func() {
				Expect(userService.ValidateAge(30)).To(BeTrue())
			})

			It("should validate the maximum age as valid", func() {
				Expect(userService.ValidateAge(120)).To(BeTrue())
			})
		})

		Context("when the age is invalid", func() {
			It("should invalidate negative ages", func() {
				Expect(userService.ValidateAge(-1)).To(BeFalse())
			})

			It("should invalidate ages above the maximum", func() {
				Expect(userService.ValidateAge(121)).To(BeFalse())
			})
		})
	})

	Describe("ProcessUserDataAsync", func() {
		BeforeEach(func() {
			// Add a test user
			mockRepo.AddUser(&ginkgo.User{
				ID:       1,
				Name:     "John Doe",
				Email:    "john@example.com",
				Age:      30,
				IsActive: true,
			})
		})

		It("should retrieve user data asynchronously", func() {
			resultCh := userService.ProcessUserDataAsync(ctx, 1)

			// Use Eventually to verify the asynchronous result
			Eventually(resultCh).Should(Receive(And(
				Not(BeNil()),
				WithTransform(func(u *ginkgo.User) int { return u.ID }, Equal(1)),
				WithTransform(func(u *ginkgo.User) string { return u.Name }, Equal("John Doe")),
			)))
		})

		It("should not return a result when the context is canceled", func() {
			// Cancel the context immediately
			cancelFunc()

			resultCh := userService.ProcessUserDataAsync(ctx, 1)

			// Verify that the channel is closed
			Eventually(resultCh).Should(BeClosed())

			// Verify that no result is sent
			Consistently(resultCh).ShouldNot(Receive())
		})
	})
})
