// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/influxdata/influxdb/v2 (interfaces: OrganizationService)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	influxdb "github.com/influxdata/influxdb/v2"
	"github.com/influxdata/influxdb/v2/kit/platform"
)

// MockOrganizationService is a mock of OrganizationService interface
type MockOrganizationService struct {
	ctrl     *gomock.Controller
	recorder *MockOrganizationServiceMockRecorder
}

// MockOrganizationServiceMockRecorder is the mock recorder for MockOrganizationService
type MockOrganizationServiceMockRecorder struct {
	mock *MockOrganizationService
}

// NewMockOrganizationService creates a new mock instance
func NewMockOrganizationService(ctrl *gomock.Controller) *MockOrganizationService {
	mock := &MockOrganizationService{ctrl: ctrl}
	mock.recorder = &MockOrganizationServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOrganizationService) EXPECT() *MockOrganizationServiceMockRecorder {
	return m.recorder
}

// CreateOrganization mocks base method
func (m *MockOrganizationService) CreateOrganization(arg0 context.Context, arg1 *influxdb.Organization) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrganization", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateOrganization indicates an expected call of CreateOrganization
func (mr *MockOrganizationServiceMockRecorder) CreateOrganization(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrganization", reflect.TypeOf((*MockOrganizationService)(nil).CreateOrganization), arg0, arg1)
}

// DeleteOrganization mocks base method
func (m *MockOrganizationService) DeleteOrganization(arg0 context.Context, arg1 platform.ID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteOrganization", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteOrganization indicates an expected call of DeleteOrganization
func (mr *MockOrganizationServiceMockRecorder) DeleteOrganization(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteOrganization", reflect.TypeOf((*MockOrganizationService)(nil).DeleteOrganization), arg0, arg1)
}

// FindOrganization mocks base method
func (m *MockOrganizationService) FindOrganization(arg0 context.Context, arg1 influxdb.OrganizationFilter) (*influxdb.Organization, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrganization", arg0, arg1)
	ret0, _ := ret[0].(*influxdb.Organization)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrganization indicates an expected call of FindOrganization
func (mr *MockOrganizationServiceMockRecorder) FindOrganization(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrganization", reflect.TypeOf((*MockOrganizationService)(nil).FindOrganization), arg0, arg1)
}

// FindOrganizationByID mocks base method
func (m *MockOrganizationService) FindOrganizationByID(arg0 context.Context, arg1 platform.ID) (*influxdb.Organization, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOrganizationByID", arg0, arg1)
	ret0, _ := ret[0].(*influxdb.Organization)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOrganizationByID indicates an expected call of FindOrganizationByID
func (mr *MockOrganizationServiceMockRecorder) FindOrganizationByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrganizationByID", reflect.TypeOf((*MockOrganizationService)(nil).FindOrganizationByID), arg0, arg1)
}

// FindOrganizations mocks base method
func (m *MockOrganizationService) FindOrganizations(arg0 context.Context, arg1 influxdb.OrganizationFilter, arg2 ...influxdb.FindOptions) ([]*influxdb.Organization, int, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "FindOrganizations", varargs...)
	ret0, _ := ret[0].([]*influxdb.Organization)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// FindOrganizations indicates an expected call of FindOrganizations
func (mr *MockOrganizationServiceMockRecorder) FindOrganizations(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOrganizations", reflect.TypeOf((*MockOrganizationService)(nil).FindOrganizations), varargs...)
}

// UpdateOrganization mocks base method
func (m *MockOrganizationService) UpdateOrganization(arg0 context.Context, arg1 platform.ID, arg2 influxdb.OrganizationUpdate) (*influxdb.Organization, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateOrganization", arg0, arg1, arg2)
	ret0, _ := ret[0].(*influxdb.Organization)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateOrganization indicates an expected call of UpdateOrganization
func (mr *MockOrganizationServiceMockRecorder) UpdateOrganization(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateOrganization", reflect.TypeOf((*MockOrganizationService)(nil).UpdateOrganization), arg0, arg1, arg2)
}
