// Code generated by mockery v2.14.1. DO NOT EDIT.

package mocks

import (
	users "echo-recipe/businesses/users"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// GetByEmail provides a mock function with given fields: email, password
func (_m *Repository) GetByEmail(email string, password string) users.Domain {
	ret := _m.Called(email, password)

	var r0 users.Domain
	if rf, ok := ret.Get(0).(func(string, string) users.Domain); ok {
		r0 = rf(email, password)
	} else {
		r0 = ret.Get(0).(users.Domain)
	}

	return r0
}

// Register provides a mock function with given fields: userDomain
func (_m *Repository) Register(userDomain *users.Domain) users.Domain {
	ret := _m.Called(userDomain)

	var r0 users.Domain
	if rf, ok := ret.Get(0).(func(*users.Domain) users.Domain); ok {
		r0 = rf(userDomain)
	} else {
		r0 = ret.Get(0).(users.Domain)
	}

	return r0
}

// Update provides a mock function with given fields: domain
func (_m *Repository) Update(domain *users.Domain) users.Domain {
	ret := _m.Called(domain)

	var r0 users.Domain
	if rf, ok := ret.Get(0).(func(*users.Domain) users.Domain); ok {
		r0 = rf(domain)
	} else {
		r0 = ret.Get(0).(users.Domain)
	}

	return r0
}

type mockConstructorTestingTNewRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewRepository creates a new instance of Repository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewRepository(t mockConstructorTestingTNewRepository) *Repository {
	mock := &Repository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
