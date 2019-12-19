// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	context "context"
	io "io"

	mock "github.com/stretchr/testify/mock"

	swarm "github.com/docker/docker/api/types/swarm"

	types "github.com/docker/docker/api/types"
)

// ServiceAPIClient is an autogenerated mock type for the ServiceAPIClient type
type ServiceAPIClient struct {
	mock.Mock
}

// ServiceCreate provides a mock function with given fields: ctx, service, options
func (_m *ServiceAPIClient) ServiceCreate(ctx context.Context, service swarm.ServiceSpec, options types.ServiceCreateOptions) (types.ServiceCreateResponse, error) {
	ret := _m.Called(ctx, service, options)

	var r0 types.ServiceCreateResponse
	if rf, ok := ret.Get(0).(func(context.Context, swarm.ServiceSpec, types.ServiceCreateOptions) types.ServiceCreateResponse); ok {
		r0 = rf(ctx, service, options)
	} else {
		r0 = ret.Get(0).(types.ServiceCreateResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, swarm.ServiceSpec, types.ServiceCreateOptions) error); ok {
		r1 = rf(ctx, service, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ServiceInspectWithRaw provides a mock function with given fields: ctx, serviceID, options
func (_m *ServiceAPIClient) ServiceInspectWithRaw(ctx context.Context, serviceID string, options types.ServiceInspectOptions) (swarm.Service, []byte, error) {
	ret := _m.Called(ctx, serviceID, options)

	var r0 swarm.Service
	if rf, ok := ret.Get(0).(func(context.Context, string, types.ServiceInspectOptions) swarm.Service); ok {
		r0 = rf(ctx, serviceID, options)
	} else {
		r0 = ret.Get(0).(swarm.Service)
	}

	var r1 []byte
	if rf, ok := ret.Get(1).(func(context.Context, string, types.ServiceInspectOptions) []byte); ok {
		r1 = rf(ctx, serviceID, options)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]byte)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string, types.ServiceInspectOptions) error); ok {
		r2 = rf(ctx, serviceID, options)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// ServiceList provides a mock function with given fields: ctx, options
func (_m *ServiceAPIClient) ServiceList(ctx context.Context, options types.ServiceListOptions) ([]swarm.Service, error) {
	ret := _m.Called(ctx, options)

	var r0 []swarm.Service
	if rf, ok := ret.Get(0).(func(context.Context, types.ServiceListOptions) []swarm.Service); ok {
		r0 = rf(ctx, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]swarm.Service)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, types.ServiceListOptions) error); ok {
		r1 = rf(ctx, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ServiceLogs provides a mock function with given fields: ctx, serviceID, options
func (_m *ServiceAPIClient) ServiceLogs(ctx context.Context, serviceID string, options types.ContainerLogsOptions) (io.ReadCloser, error) {
	ret := _m.Called(ctx, serviceID, options)

	var r0 io.ReadCloser
	if rf, ok := ret.Get(0).(func(context.Context, string, types.ContainerLogsOptions) io.ReadCloser); ok {
		r0 = rf(ctx, serviceID, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.ReadCloser)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, types.ContainerLogsOptions) error); ok {
		r1 = rf(ctx, serviceID, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ServiceRemove provides a mock function with given fields: ctx, serviceID
func (_m *ServiceAPIClient) ServiceRemove(ctx context.Context, serviceID string) error {
	ret := _m.Called(ctx, serviceID)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, serviceID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ServiceUpdate provides a mock function with given fields: ctx, serviceID, version, service, options
func (_m *ServiceAPIClient) ServiceUpdate(ctx context.Context, serviceID string, version swarm.Version, service swarm.ServiceSpec, options types.ServiceUpdateOptions) (types.ServiceUpdateResponse, error) {
	ret := _m.Called(ctx, serviceID, version, service, options)

	var r0 types.ServiceUpdateResponse
	if rf, ok := ret.Get(0).(func(context.Context, string, swarm.Version, swarm.ServiceSpec, types.ServiceUpdateOptions) types.ServiceUpdateResponse); ok {
		r0 = rf(ctx, serviceID, version, service, options)
	} else {
		r0 = ret.Get(0).(types.ServiceUpdateResponse)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, swarm.Version, swarm.ServiceSpec, types.ServiceUpdateOptions) error); ok {
		r1 = rf(ctx, serviceID, version, service, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TaskInspectWithRaw provides a mock function with given fields: ctx, taskID
func (_m *ServiceAPIClient) TaskInspectWithRaw(ctx context.Context, taskID string) (swarm.Task, []byte, error) {
	ret := _m.Called(ctx, taskID)

	var r0 swarm.Task
	if rf, ok := ret.Get(0).(func(context.Context, string) swarm.Task); ok {
		r0 = rf(ctx, taskID)
	} else {
		r0 = ret.Get(0).(swarm.Task)
	}

	var r1 []byte
	if rf, ok := ret.Get(1).(func(context.Context, string) []byte); ok {
		r1 = rf(ctx, taskID)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]byte)
		}
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(context.Context, string) error); ok {
		r2 = rf(ctx, taskID)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// TaskList provides a mock function with given fields: ctx, options
func (_m *ServiceAPIClient) TaskList(ctx context.Context, options types.TaskListOptions) ([]swarm.Task, error) {
	ret := _m.Called(ctx, options)

	var r0 []swarm.Task
	if rf, ok := ret.Get(0).(func(context.Context, types.TaskListOptions) []swarm.Task); ok {
		r0 = rf(ctx, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]swarm.Task)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, types.TaskListOptions) error); ok {
		r1 = rf(ctx, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// TaskLogs provides a mock function with given fields: ctx, taskID, options
func (_m *ServiceAPIClient) TaskLogs(ctx context.Context, taskID string, options types.ContainerLogsOptions) (io.ReadCloser, error) {
	ret := _m.Called(ctx, taskID, options)

	var r0 io.ReadCloser
	if rf, ok := ret.Get(0).(func(context.Context, string, types.ContainerLogsOptions) io.ReadCloser); ok {
		r0 = rf(ctx, taskID, options)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(io.ReadCloser)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, types.ContainerLogsOptions) error); ok {
		r1 = rf(ctx, taskID, options)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
