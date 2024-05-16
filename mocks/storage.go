// Code generated by mockery v3.0.0-alpha.0. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/voltento/go-blog-project/internal/domain"

	mock "github.com/stretchr/testify/mock"
)

// Storage is an autogenerated mock type for the Storage type
type Storage struct {
	mock.Mock
}

// CreatePost provides a mock function with given fields: ctx, post
func (_m *Storage) CreatePost(ctx context.Context, post *domain.Post) (domain.PostId, error) {
	ret := _m.Called(ctx, post)

	var r0 domain.PostId
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Post) domain.PostId); ok {
		r0 = rf(ctx, post)
	} else {
		r0 = ret.Get(0).(domain.PostId)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *domain.Post) error); ok {
		r1 = rf(ctx, post)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewStorage interface {
	mock.TestingT
	Cleanup(func())
}

// NewStorage creates a new instance of Storage. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewStorage(t mockConstructorTestingTNewStorage) *Storage {
	mock := &Storage{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}