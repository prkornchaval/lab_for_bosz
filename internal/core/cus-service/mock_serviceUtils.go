// Code generated by mockery v2.33.3. DO NOT EDIT.

package cusservice

import mock "github.com/stretchr/testify/mock"

// mockServiceUtils is an autogenerated mock type for the serviceUtils type
type mockServiceUtils struct {
	mock.Mock
}

// createCustomers provides a mock function with given fields: s
func (_m *mockServiceUtils) createCustomers(s *service) error {
	ret := _m.Called(s)

	var r0 error
	if rf, ok := ret.Get(0).(func(*service) error); ok {
		r0 = rf(s)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// newMockServiceUtils creates a new instance of mockServiceUtils. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func newMockServiceUtils(t interface {
	mock.TestingT
	Cleanup(func())
}) *mockServiceUtils {
	mock := &mockServiceUtils{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
