// Code generated by mockery v2.33.2. DO NOT EDIT.

package mocks

import (
	context "context"

	dto "github.com/rahul7668gupta/go-url-shortner/pkg/dto"
	mock "github.com/stretchr/testify/mock"
)

// IShortnerRepo is an autogenerated mock type for the IShortnerRepo type
type IShortnerRepo struct {
	mock.Mock
}

// CreateIndexOnOriginalUrl provides a mock function with given fields: ctx, url, shortCode
func (_m *IShortnerRepo) CreateIndexOnOriginalUrl(ctx context.Context, url string, shortCode string) error {
	ret := _m.Called(ctx, url, shortCode)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, url, shortCode)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CreateShortCodeRecord provides a mock function with given fields: ctx, shortCode, requestUrl
func (_m *IShortnerRepo) CreateShortCodeRecord(ctx context.Context, shortCode string, requestUrl string) error {
	ret := _m.Called(ctx, shortCode, requestUrl)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string) error); ok {
		r0 = rf(ctx, shortCode, requestUrl)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetMetrics provides a mock function with given fields: ctx
func (_m *IShortnerRepo) GetMetrics(ctx context.Context) ([]dto.Metrics, error) {
	ret := _m.Called(ctx)

	var r0 []dto.Metrics
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) ([]dto.Metrics, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) []dto.Metrics); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]dto.Metrics)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOriginalUrlForShortCode provides a mock function with given fields: ctx, shortCode
func (_m *IShortnerRepo) GetOriginalUrlForShortCode(ctx context.Context, shortCode string) (string, error) {
	ret := _m.Called(ctx, shortCode)

	var r0 string
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (string, error)); ok {
		return rf(ctx, shortCode)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(ctx, shortCode)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, shortCode)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IncrementCounterForShortCode provides a mock function with given fields: ctx
func (_m *IShortnerRepo) IncrementCounterForShortCode(ctx context.Context) (int64, error) {
	ret := _m.Called(ctx)

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (int64, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) int64); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IncrementDomainCounter provides a mock function with given fields: ctx, domain
func (_m *IShortnerRepo) IncrementDomainCounter(ctx context.Context, domain string) error {
	ret := _m.Called(ctx, domain)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, domain)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LookupURL provides a mock function with given fields: ctx, url
func (_m *IShortnerRepo) LookupURL(ctx context.Context, url string) (string, bool) {
	ret := _m.Called(ctx, url)

	var r0 string
	var r1 bool
	if rf, ok := ret.Get(0).(func(context.Context, string) (string, bool)); ok {
		return rf(ctx, url)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(ctx, url)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) bool); ok {
		r1 = rf(ctx, url)
	} else {
		r1 = ret.Get(1).(bool)
	}

	return r0, r1
}

// NewIShortnerRepo creates a new instance of IShortnerRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIShortnerRepo(t interface {
	mock.TestingT
	Cleanup(func())
}) *IShortnerRepo {
	mock := &IShortnerRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
