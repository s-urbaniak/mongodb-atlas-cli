// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/mongodb/mongodb-atlas-cli/internal/store (interfaces: PipelinesLister,PipelinesDescriber,PipelinesCreator,PipelinesUpdater,PipelinesDeleter)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	mongodbatlasv2 "go.mongodb.org/atlas/mongodbatlasv2"
)

// MockPipelinesLister is a mock of PipelinesLister interface.
type MockPipelinesLister struct {
	ctrl     *gomock.Controller
	recorder *MockPipelinesListerMockRecorder
}

// MockPipelinesListerMockRecorder is the mock recorder for MockPipelinesLister.
type MockPipelinesListerMockRecorder struct {
	mock *MockPipelinesLister
}

// NewMockPipelinesLister creates a new mock instance.
func NewMockPipelinesLister(ctrl *gomock.Controller) *MockPipelinesLister {
	mock := &MockPipelinesLister{ctrl: ctrl}
	mock.recorder = &MockPipelinesListerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPipelinesLister) EXPECT() *MockPipelinesListerMockRecorder {
	return m.recorder
}

// Pipelines mocks base method.
func (m *MockPipelinesLister) Pipelines(arg0 string) ([]mongodbatlasv2.IngestionPipeline, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Pipelines", arg0)
	ret0, _ := ret[0].([]mongodbatlasv2.IngestionPipeline)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Pipelines indicates an expected call of Pipelines.
func (mr *MockPipelinesListerMockRecorder) Pipelines(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pipelines", reflect.TypeOf((*MockPipelinesLister)(nil).Pipelines), arg0)
}

// MockPipelinesDescriber is a mock of PipelinesDescriber interface.
type MockPipelinesDescriber struct {
	ctrl     *gomock.Controller
	recorder *MockPipelinesDescriberMockRecorder
}

// MockPipelinesDescriberMockRecorder is the mock recorder for MockPipelinesDescriber.
type MockPipelinesDescriberMockRecorder struct {
	mock *MockPipelinesDescriber
}

// NewMockPipelinesDescriber creates a new mock instance.
func NewMockPipelinesDescriber(ctrl *gomock.Controller) *MockPipelinesDescriber {
	mock := &MockPipelinesDescriber{ctrl: ctrl}
	mock.recorder = &MockPipelinesDescriberMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPipelinesDescriber) EXPECT() *MockPipelinesDescriberMockRecorder {
	return m.recorder
}

// Pipeline mocks base method.
func (m *MockPipelinesDescriber) Pipeline(arg0, arg1 string) (*mongodbatlasv2.IngestionPipeline, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Pipeline", arg0, arg1)
	ret0, _ := ret[0].(*mongodbatlasv2.IngestionPipeline)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Pipeline indicates an expected call of Pipeline.
func (mr *MockPipelinesDescriberMockRecorder) Pipeline(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Pipeline", reflect.TypeOf((*MockPipelinesDescriber)(nil).Pipeline), arg0, arg1)
}

// MockPipelinesCreator is a mock of PipelinesCreator interface.
type MockPipelinesCreator struct {
	ctrl     *gomock.Controller
	recorder *MockPipelinesCreatorMockRecorder
}

// MockPipelinesCreatorMockRecorder is the mock recorder for MockPipelinesCreator.
type MockPipelinesCreatorMockRecorder struct {
	mock *MockPipelinesCreator
}

// NewMockPipelinesCreator creates a new mock instance.
func NewMockPipelinesCreator(ctrl *gomock.Controller) *MockPipelinesCreator {
	mock := &MockPipelinesCreator{ctrl: ctrl}
	mock.recorder = &MockPipelinesCreatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPipelinesCreator) EXPECT() *MockPipelinesCreatorMockRecorder {
	return m.recorder
}

// CreatePipeline mocks base method.
func (m *MockPipelinesCreator) CreatePipeline(arg0 string, arg1 mongodbatlasv2.IngestionPipeline) (*mongodbatlasv2.IngestionPipeline, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreatePipeline", arg0, arg1)
	ret0, _ := ret[0].(*mongodbatlasv2.IngestionPipeline)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreatePipeline indicates an expected call of CreatePipeline.
func (mr *MockPipelinesCreatorMockRecorder) CreatePipeline(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreatePipeline", reflect.TypeOf((*MockPipelinesCreator)(nil).CreatePipeline), arg0, arg1)
}

// MockPipelinesUpdater is a mock of PipelinesUpdater interface.
type MockPipelinesUpdater struct {
	ctrl     *gomock.Controller
	recorder *MockPipelinesUpdaterMockRecorder
}

// MockPipelinesUpdaterMockRecorder is the mock recorder for MockPipelinesUpdater.
type MockPipelinesUpdaterMockRecorder struct {
	mock *MockPipelinesUpdater
}

// NewMockPipelinesUpdater creates a new mock instance.
func NewMockPipelinesUpdater(ctrl *gomock.Controller) *MockPipelinesUpdater {
	mock := &MockPipelinesUpdater{ctrl: ctrl}
	mock.recorder = &MockPipelinesUpdaterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPipelinesUpdater) EXPECT() *MockPipelinesUpdaterMockRecorder {
	return m.recorder
}

// UpdatePipeline mocks base method.
func (m *MockPipelinesUpdater) UpdatePipeline(arg0, arg1 string, arg2 mongodbatlasv2.IngestionPipeline) (*mongodbatlasv2.IngestionPipeline, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatePipeline", arg0, arg1, arg2)
	ret0, _ := ret[0].(*mongodbatlasv2.IngestionPipeline)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdatePipeline indicates an expected call of UpdatePipeline.
func (mr *MockPipelinesUpdaterMockRecorder) UpdatePipeline(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatePipeline", reflect.TypeOf((*MockPipelinesUpdater)(nil).UpdatePipeline), arg0, arg1, arg2)
}

// MockPipelinesDeleter is a mock of PipelinesDeleter interface.
type MockPipelinesDeleter struct {
	ctrl     *gomock.Controller
	recorder *MockPipelinesDeleterMockRecorder
}

// MockPipelinesDeleterMockRecorder is the mock recorder for MockPipelinesDeleter.
type MockPipelinesDeleterMockRecorder struct {
	mock *MockPipelinesDeleter
}

// NewMockPipelinesDeleter creates a new mock instance.
func NewMockPipelinesDeleter(ctrl *gomock.Controller) *MockPipelinesDeleter {
	mock := &MockPipelinesDeleter{ctrl: ctrl}
	mock.recorder = &MockPipelinesDeleterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPipelinesDeleter) EXPECT() *MockPipelinesDeleterMockRecorder {
	return m.recorder
}

// DeletePipeline mocks base method.
func (m *MockPipelinesDeleter) DeletePipeline(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePipeline", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePipeline indicates an expected call of DeletePipeline.
func (mr *MockPipelinesDeleterMockRecorder) DeletePipeline(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePipeline", reflect.TypeOf((*MockPipelinesDeleter)(nil).DeletePipeline), arg0, arg1)
}