// Code generated by MockGen. DO NOT EDIT.
// Source: ./multicluster_reconcilers.go

// Package mock_controller is a generated GoMock package.
package mock_controller

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1alpha1 "github.com/solo-io/service-mesh-hub/pkg/api/discovery.smh.solo.io/v1alpha1"
	controller "github.com/solo-io/service-mesh-hub/pkg/api/discovery.smh.solo.io/v1alpha1/controller"
	reconcile "github.com/solo-io/skv2/pkg/reconcile"
	predicate "sigs.k8s.io/controller-runtime/pkg/predicate"
)

// MockMulticlusterMeshServiceReconciler is a mock of MulticlusterMeshServiceReconciler interface.
type MockMulticlusterMeshServiceReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterMeshServiceReconcilerMockRecorder
}

// MockMulticlusterMeshServiceReconcilerMockRecorder is the mock recorder for MockMulticlusterMeshServiceReconciler.
type MockMulticlusterMeshServiceReconcilerMockRecorder struct {
	mock *MockMulticlusterMeshServiceReconciler
}

// NewMockMulticlusterMeshServiceReconciler creates a new mock instance.
func NewMockMulticlusterMeshServiceReconciler(ctrl *gomock.Controller) *MockMulticlusterMeshServiceReconciler {
	mock := &MockMulticlusterMeshServiceReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterMeshServiceReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterMeshServiceReconciler) EXPECT() *MockMulticlusterMeshServiceReconcilerMockRecorder {
	return m.recorder
}

// ReconcileMeshService mocks base method.
func (m *MockMulticlusterMeshServiceReconciler) ReconcileMeshService(clusterName string, obj *v1alpha1.MeshService) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileMeshService", clusterName, obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileMeshService indicates an expected call of ReconcileMeshService.
func (mr *MockMulticlusterMeshServiceReconcilerMockRecorder) ReconcileMeshService(clusterName, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileMeshService", reflect.TypeOf((*MockMulticlusterMeshServiceReconciler)(nil).ReconcileMeshService), clusterName, obj)
}

// MockMulticlusterMeshServiceDeletionReconciler is a mock of MulticlusterMeshServiceDeletionReconciler interface.
type MockMulticlusterMeshServiceDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterMeshServiceDeletionReconcilerMockRecorder
}

// MockMulticlusterMeshServiceDeletionReconcilerMockRecorder is the mock recorder for MockMulticlusterMeshServiceDeletionReconciler.
type MockMulticlusterMeshServiceDeletionReconcilerMockRecorder struct {
	mock *MockMulticlusterMeshServiceDeletionReconciler
}

// NewMockMulticlusterMeshServiceDeletionReconciler creates a new mock instance.
func NewMockMulticlusterMeshServiceDeletionReconciler(ctrl *gomock.Controller) *MockMulticlusterMeshServiceDeletionReconciler {
	mock := &MockMulticlusterMeshServiceDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterMeshServiceDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterMeshServiceDeletionReconciler) EXPECT() *MockMulticlusterMeshServiceDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileMeshServiceDeletion mocks base method.
func (m *MockMulticlusterMeshServiceDeletionReconciler) ReconcileMeshServiceDeletion(clusterName string, req reconcile.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReconcileMeshServiceDeletion", clusterName, req)
}

// ReconcileMeshServiceDeletion indicates an expected call of ReconcileMeshServiceDeletion.
func (mr *MockMulticlusterMeshServiceDeletionReconcilerMockRecorder) ReconcileMeshServiceDeletion(clusterName, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileMeshServiceDeletion", reflect.TypeOf((*MockMulticlusterMeshServiceDeletionReconciler)(nil).ReconcileMeshServiceDeletion), clusterName, req)
}

// MockMulticlusterMeshServiceReconcileLoop is a mock of MulticlusterMeshServiceReconcileLoop interface.
type MockMulticlusterMeshServiceReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterMeshServiceReconcileLoopMockRecorder
}

// MockMulticlusterMeshServiceReconcileLoopMockRecorder is the mock recorder for MockMulticlusterMeshServiceReconcileLoop.
type MockMulticlusterMeshServiceReconcileLoopMockRecorder struct {
	mock *MockMulticlusterMeshServiceReconcileLoop
}

// NewMockMulticlusterMeshServiceReconcileLoop creates a new mock instance.
func NewMockMulticlusterMeshServiceReconcileLoop(ctrl *gomock.Controller) *MockMulticlusterMeshServiceReconcileLoop {
	mock := &MockMulticlusterMeshServiceReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockMulticlusterMeshServiceReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterMeshServiceReconcileLoop) EXPECT() *MockMulticlusterMeshServiceReconcileLoopMockRecorder {
	return m.recorder
}

// AddMulticlusterMeshServiceReconciler mocks base method.
func (m *MockMulticlusterMeshServiceReconcileLoop) AddMulticlusterMeshServiceReconciler(ctx context.Context, rec controller.MulticlusterMeshServiceReconciler, predicates ...predicate.Predicate) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddMulticlusterMeshServiceReconciler", varargs...)
}

// AddMulticlusterMeshServiceReconciler indicates an expected call of AddMulticlusterMeshServiceReconciler.
func (mr *MockMulticlusterMeshServiceReconcileLoopMockRecorder) AddMulticlusterMeshServiceReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMulticlusterMeshServiceReconciler", reflect.TypeOf((*MockMulticlusterMeshServiceReconcileLoop)(nil).AddMulticlusterMeshServiceReconciler), varargs...)
}

// MockMulticlusterMeshWorkloadReconciler is a mock of MulticlusterMeshWorkloadReconciler interface.
type MockMulticlusterMeshWorkloadReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterMeshWorkloadReconcilerMockRecorder
}

// MockMulticlusterMeshWorkloadReconcilerMockRecorder is the mock recorder for MockMulticlusterMeshWorkloadReconciler.
type MockMulticlusterMeshWorkloadReconcilerMockRecorder struct {
	mock *MockMulticlusterMeshWorkloadReconciler
}

// NewMockMulticlusterMeshWorkloadReconciler creates a new mock instance.
func NewMockMulticlusterMeshWorkloadReconciler(ctrl *gomock.Controller) *MockMulticlusterMeshWorkloadReconciler {
	mock := &MockMulticlusterMeshWorkloadReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterMeshWorkloadReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterMeshWorkloadReconciler) EXPECT() *MockMulticlusterMeshWorkloadReconcilerMockRecorder {
	return m.recorder
}

// ReconcileMeshWorkload mocks base method.
func (m *MockMulticlusterMeshWorkloadReconciler) ReconcileMeshWorkload(clusterName string, obj *v1alpha1.MeshWorkload) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileMeshWorkload", clusterName, obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileMeshWorkload indicates an expected call of ReconcileMeshWorkload.
func (mr *MockMulticlusterMeshWorkloadReconcilerMockRecorder) ReconcileMeshWorkload(clusterName, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileMeshWorkload", reflect.TypeOf((*MockMulticlusterMeshWorkloadReconciler)(nil).ReconcileMeshWorkload), clusterName, obj)
}

// MockMulticlusterMeshWorkloadDeletionReconciler is a mock of MulticlusterMeshWorkloadDeletionReconciler interface.
type MockMulticlusterMeshWorkloadDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterMeshWorkloadDeletionReconcilerMockRecorder
}

// MockMulticlusterMeshWorkloadDeletionReconcilerMockRecorder is the mock recorder for MockMulticlusterMeshWorkloadDeletionReconciler.
type MockMulticlusterMeshWorkloadDeletionReconcilerMockRecorder struct {
	mock *MockMulticlusterMeshWorkloadDeletionReconciler
}

// NewMockMulticlusterMeshWorkloadDeletionReconciler creates a new mock instance.
func NewMockMulticlusterMeshWorkloadDeletionReconciler(ctrl *gomock.Controller) *MockMulticlusterMeshWorkloadDeletionReconciler {
	mock := &MockMulticlusterMeshWorkloadDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterMeshWorkloadDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterMeshWorkloadDeletionReconciler) EXPECT() *MockMulticlusterMeshWorkloadDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileMeshWorkloadDeletion mocks base method.
func (m *MockMulticlusterMeshWorkloadDeletionReconciler) ReconcileMeshWorkloadDeletion(clusterName string, req reconcile.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReconcileMeshWorkloadDeletion", clusterName, req)
}

// ReconcileMeshWorkloadDeletion indicates an expected call of ReconcileMeshWorkloadDeletion.
func (mr *MockMulticlusterMeshWorkloadDeletionReconcilerMockRecorder) ReconcileMeshWorkloadDeletion(clusterName, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileMeshWorkloadDeletion", reflect.TypeOf((*MockMulticlusterMeshWorkloadDeletionReconciler)(nil).ReconcileMeshWorkloadDeletion), clusterName, req)
}

// MockMulticlusterMeshWorkloadReconcileLoop is a mock of MulticlusterMeshWorkloadReconcileLoop interface.
type MockMulticlusterMeshWorkloadReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterMeshWorkloadReconcileLoopMockRecorder
}

// MockMulticlusterMeshWorkloadReconcileLoopMockRecorder is the mock recorder for MockMulticlusterMeshWorkloadReconcileLoop.
type MockMulticlusterMeshWorkloadReconcileLoopMockRecorder struct {
	mock *MockMulticlusterMeshWorkloadReconcileLoop
}

// NewMockMulticlusterMeshWorkloadReconcileLoop creates a new mock instance.
func NewMockMulticlusterMeshWorkloadReconcileLoop(ctrl *gomock.Controller) *MockMulticlusterMeshWorkloadReconcileLoop {
	mock := &MockMulticlusterMeshWorkloadReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockMulticlusterMeshWorkloadReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterMeshWorkloadReconcileLoop) EXPECT() *MockMulticlusterMeshWorkloadReconcileLoopMockRecorder {
	return m.recorder
}

// AddMulticlusterMeshWorkloadReconciler mocks base method.
func (m *MockMulticlusterMeshWorkloadReconcileLoop) AddMulticlusterMeshWorkloadReconciler(ctx context.Context, rec controller.MulticlusterMeshWorkloadReconciler, predicates ...predicate.Predicate) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddMulticlusterMeshWorkloadReconciler", varargs...)
}

// AddMulticlusterMeshWorkloadReconciler indicates an expected call of AddMulticlusterMeshWorkloadReconciler.
func (mr *MockMulticlusterMeshWorkloadReconcileLoopMockRecorder) AddMulticlusterMeshWorkloadReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMulticlusterMeshWorkloadReconciler", reflect.TypeOf((*MockMulticlusterMeshWorkloadReconcileLoop)(nil).AddMulticlusterMeshWorkloadReconciler), varargs...)
}

// MockMulticlusterMeshReconciler is a mock of MulticlusterMeshReconciler interface.
type MockMulticlusterMeshReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterMeshReconcilerMockRecorder
}

// MockMulticlusterMeshReconcilerMockRecorder is the mock recorder for MockMulticlusterMeshReconciler.
type MockMulticlusterMeshReconcilerMockRecorder struct {
	mock *MockMulticlusterMeshReconciler
}

// NewMockMulticlusterMeshReconciler creates a new mock instance.
func NewMockMulticlusterMeshReconciler(ctrl *gomock.Controller) *MockMulticlusterMeshReconciler {
	mock := &MockMulticlusterMeshReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterMeshReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterMeshReconciler) EXPECT() *MockMulticlusterMeshReconcilerMockRecorder {
	return m.recorder
}

// ReconcileMesh mocks base method.
func (m *MockMulticlusterMeshReconciler) ReconcileMesh(clusterName string, obj *v1alpha1.Mesh) (reconcile.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReconcileMesh", clusterName, obj)
	ret0, _ := ret[0].(reconcile.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReconcileMesh indicates an expected call of ReconcileMesh.
func (mr *MockMulticlusterMeshReconcilerMockRecorder) ReconcileMesh(clusterName, obj interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileMesh", reflect.TypeOf((*MockMulticlusterMeshReconciler)(nil).ReconcileMesh), clusterName, obj)
}

// MockMulticlusterMeshDeletionReconciler is a mock of MulticlusterMeshDeletionReconciler interface.
type MockMulticlusterMeshDeletionReconciler struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterMeshDeletionReconcilerMockRecorder
}

// MockMulticlusterMeshDeletionReconcilerMockRecorder is the mock recorder for MockMulticlusterMeshDeletionReconciler.
type MockMulticlusterMeshDeletionReconcilerMockRecorder struct {
	mock *MockMulticlusterMeshDeletionReconciler
}

// NewMockMulticlusterMeshDeletionReconciler creates a new mock instance.
func NewMockMulticlusterMeshDeletionReconciler(ctrl *gomock.Controller) *MockMulticlusterMeshDeletionReconciler {
	mock := &MockMulticlusterMeshDeletionReconciler{ctrl: ctrl}
	mock.recorder = &MockMulticlusterMeshDeletionReconcilerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterMeshDeletionReconciler) EXPECT() *MockMulticlusterMeshDeletionReconcilerMockRecorder {
	return m.recorder
}

// ReconcileMeshDeletion mocks base method.
func (m *MockMulticlusterMeshDeletionReconciler) ReconcileMeshDeletion(clusterName string, req reconcile.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ReconcileMeshDeletion", clusterName, req)
}

// ReconcileMeshDeletion indicates an expected call of ReconcileMeshDeletion.
func (mr *MockMulticlusterMeshDeletionReconcilerMockRecorder) ReconcileMeshDeletion(clusterName, req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReconcileMeshDeletion", reflect.TypeOf((*MockMulticlusterMeshDeletionReconciler)(nil).ReconcileMeshDeletion), clusterName, req)
}

// MockMulticlusterMeshReconcileLoop is a mock of MulticlusterMeshReconcileLoop interface.
type MockMulticlusterMeshReconcileLoop struct {
	ctrl     *gomock.Controller
	recorder *MockMulticlusterMeshReconcileLoopMockRecorder
}

// MockMulticlusterMeshReconcileLoopMockRecorder is the mock recorder for MockMulticlusterMeshReconcileLoop.
type MockMulticlusterMeshReconcileLoopMockRecorder struct {
	mock *MockMulticlusterMeshReconcileLoop
}

// NewMockMulticlusterMeshReconcileLoop creates a new mock instance.
func NewMockMulticlusterMeshReconcileLoop(ctrl *gomock.Controller) *MockMulticlusterMeshReconcileLoop {
	mock := &MockMulticlusterMeshReconcileLoop{ctrl: ctrl}
	mock.recorder = &MockMulticlusterMeshReconcileLoopMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMulticlusterMeshReconcileLoop) EXPECT() *MockMulticlusterMeshReconcileLoopMockRecorder {
	return m.recorder
}

// AddMulticlusterMeshReconciler mocks base method.
func (m *MockMulticlusterMeshReconcileLoop) AddMulticlusterMeshReconciler(ctx context.Context, rec controller.MulticlusterMeshReconciler, predicates ...predicate.Predicate) {
	m.ctrl.T.Helper()
	varargs := []interface{}{ctx, rec}
	for _, a := range predicates {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "AddMulticlusterMeshReconciler", varargs...)
}

// AddMulticlusterMeshReconciler indicates an expected call of AddMulticlusterMeshReconciler.
func (mr *MockMulticlusterMeshReconcileLoopMockRecorder) AddMulticlusterMeshReconciler(ctx, rec interface{}, predicates ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{ctx, rec}, predicates...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMulticlusterMeshReconciler", reflect.TypeOf((*MockMulticlusterMeshReconcileLoop)(nil).AddMulticlusterMeshReconciler), varargs...)
}
