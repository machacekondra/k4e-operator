// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/jakub-dzon/k4e-device-worker/internal/workload/mapping (interfaces: MappingRepository)

// Package mapping is a generated GoMock package.
package mapping

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockMappingRepository is a mock of MappingRepository interface.
type MockMappingRepository struct {
	ctrl     *gomock.Controller
	recorder *MockMappingRepositoryMockRecorder
}

// MockMappingRepositoryMockRecorder is the mock recorder for MockMappingRepository.
type MockMappingRepositoryMockRecorder struct {
	mock *MockMappingRepository
}

// NewMockMappingRepository creates a new mock instance.
func NewMockMappingRepository(ctrl *gomock.Controller) *MockMappingRepository {
	mock := &MockMappingRepository{ctrl: ctrl}
	mock.recorder = &MockMappingRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMappingRepository) EXPECT() *MockMappingRepositoryMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockMappingRepository) Add(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockMappingRepositoryMockRecorder) Add(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockMappingRepository)(nil).Add), arg0, arg1)
}

// GetId mocks base method.
func (m *MockMappingRepository) GetId(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetId", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetId indicates an expected call of GetId.
func (mr *MockMappingRepositoryMockRecorder) GetId(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetId", reflect.TypeOf((*MockMappingRepository)(nil).GetId), arg0)
}

// GetName mocks base method.
func (m *MockMappingRepository) GetName(arg0 string) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetName", arg0)
	ret0, _ := ret[0].(string)
	return ret0
}

// GetName indicates an expected call of GetName.
func (mr *MockMappingRepositoryMockRecorder) GetName(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetName", reflect.TypeOf((*MockMappingRepository)(nil).GetName), arg0)
}

// Persist mocks base method.
func (m *MockMappingRepository) Persist() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Persist")
	ret0, _ := ret[0].(error)
	return ret0
}

// Persist indicates an expected call of Persist.
func (mr *MockMappingRepositoryMockRecorder) Persist() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Persist", reflect.TypeOf((*MockMappingRepository)(nil).Persist))
}

// Remove mocks base method.
func (m *MockMappingRepository) Remove(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remove", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove.
func (mr *MockMappingRepositoryMockRecorder) Remove(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockMappingRepository)(nil).Remove), arg0)
}

// RemoveMappingFile mocks base method.
func (m *MockMappingRepository) RemoveMappingFile() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveMappingFile")
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveMappingFile indicates an expected call of RemoveMappingFile.
func (mr *MockMappingRepositoryMockRecorder) RemoveMappingFile() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveMappingFile", reflect.TypeOf((*MockMappingRepository)(nil).RemoveMappingFile))
}

// Size mocks base method.
func (m *MockMappingRepository) Size() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Size")
	ret0, _ := ret[0].(int)
	return ret0
}

// Size indicates an expected call of Size.
func (mr *MockMappingRepositoryMockRecorder) Size() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Size", reflect.TypeOf((*MockMappingRepository)(nil).Size))
}