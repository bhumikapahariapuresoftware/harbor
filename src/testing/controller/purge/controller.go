// Code generated by mockery v2.1.0. DO NOT EDIT.

package purge

import (
	context "context"

	purge "github.com/goharbor/harbor/src/controller/purge"
	mock "github.com/stretchr/testify/mock"
)

// Controller is an autogenerated mock type for the Controller type
type Controller struct {
	mock.Mock
}

// Start provides a mock function with given fields: ctx, policy, trigger
func (_m *Controller) Start(ctx context.Context, policy purge.JobPolicy, trigger string) (int64, error) {
	ret := _m.Called(ctx, policy, trigger)

	var r0 int64
	if rf, ok := ret.Get(0).(func(context.Context, purge.JobPolicy, string) int64); ok {
		r0 = rf(ctx, policy, trigger)
	} else {
		r0 = ret.Get(0).(int64)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, purge.JobPolicy, string) error); ok {
		r1 = rf(ctx, policy, trigger)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
