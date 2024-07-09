// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	model "scheduler-service/pkg/model"

	mock "github.com/stretchr/testify/mock"
)

// JobScheduler is an autogenerated mock type for the JobScheduler type
type JobScheduler struct {
	mock.Mock
}

// AddJob provides a mock function with given fields: job
func (_m *JobScheduler) AddJob(job *model.Job) {
	_m.Called(job)
}

// ProcessJobs provides a mock function with given fields:
func (_m *JobScheduler) ProcessJobs() {
	_m.Called()
}

// NewJobScheduler creates a new instance of JobScheduler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewJobScheduler(t interface {
	mock.TestingT
	Cleanup(func())
}) *JobScheduler {
	mock := &JobScheduler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
