// Code generated by mockery v3.0.0-alpha.0. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/voltento/go-blog-project/internal/domain"

	mock "github.com/stretchr/testify/mock"
)

// BlogService is an autogenerated mock type for the BlogService type
type BlogService struct {
	mock.Mock
}

// CreatePost provides a mock function with given fields: ctx, p
func (_m *BlogService) CreatePost(ctx context.Context, p *domain.Post) (domain.PostId, error) {
	ret := _m.Called(ctx, p)

	var r0 domain.PostId
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Post) domain.PostId); ok {
		r0 = rf(ctx, p)
	} else {
		r0 = ret.Get(0).(domain.PostId)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *domain.Post) error); ok {
		r1 = rf(ctx, p)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeletePost provides a mock function with given fields: ctx, id
func (_m *BlogService) DeletePost(ctx context.Context, id domain.PostId) error {
	ret := _m.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.PostId) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Post provides a mock function with given fields: ctx, id
func (_m *BlogService) Post(ctx context.Context, id domain.PostId) (*domain.Post, error) {
	ret := _m.Called(ctx, id)

	var r0 *domain.Post
	if rf, ok := ret.Get(0).(func(context.Context, domain.PostId) *domain.Post); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Post)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, domain.PostId) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdatePost provides a mock function with given fields: ctx, post, id
func (_m *BlogService) UpdatePost(ctx context.Context, post *domain.Post, id domain.PostId) error {
	ret := _m.Called(ctx, post, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Post, domain.PostId) error); ok {
		r0 = rf(ctx, post, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewBlogService interface {
	mock.TestingT
	Cleanup(func())
}

// NewBlogService creates a new instance of BlogService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewBlogService(t mockConstructorTestingTNewBlogService) *BlogService {
	mock := &BlogService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
