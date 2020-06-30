// Code generated by skv2. DO NOT EDIT.

//go:generate mockgen -source ./multicluster_reconcilers.go -destination mocks/multicluster_reconcilers.go

// Definitions for the multicluster Kubernetes Controllers
package controller

import (
	"context"

	networking_smh_solo_io_v1alpha1 "github.com/solo-io/service-mesh-hub/pkg/api/networking.smh.solo.io/v1alpha1"

	"github.com/pkg/errors"
	"github.com/solo-io/skv2/pkg/ezkube"
	"github.com/solo-io/skv2/pkg/multicluster"
	mc_reconcile "github.com/solo-io/skv2/pkg/multicluster/reconcile"
	"github.com/solo-io/skv2/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

// Reconcile Upsert events for the TrafficPolicy Resource across clusters.
// implemented by the user
type MulticlusterTrafficPolicyReconciler interface {
	ReconcileTrafficPolicy(clusterName string, obj *networking_smh_solo_io_v1alpha1.TrafficPolicy) (reconcile.Result, error)
}

// Reconcile deletion events for the TrafficPolicy Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterTrafficPolicyDeletionReconciler interface {
	ReconcileTrafficPolicyDeletion(clusterName string, req reconcile.Request)
}

type MulticlusterTrafficPolicyReconcilerFuncs struct {
	OnReconcileTrafficPolicy         func(clusterName string, obj *networking_smh_solo_io_v1alpha1.TrafficPolicy) (reconcile.Result, error)
	OnReconcileTrafficPolicyDeletion func(clusterName string, req reconcile.Request)
}

func (f *MulticlusterTrafficPolicyReconcilerFuncs) ReconcileTrafficPolicy(clusterName string, obj *networking_smh_solo_io_v1alpha1.TrafficPolicy) (reconcile.Result, error) {
	if f.OnReconcileTrafficPolicy == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileTrafficPolicy(clusterName, obj)
}

func (f *MulticlusterTrafficPolicyReconcilerFuncs) ReconcileTrafficPolicyDeletion(clusterName string, req reconcile.Request) {
	if f.OnReconcileTrafficPolicyDeletion == nil {
		return
	}
	f.OnReconcileTrafficPolicyDeletion(clusterName, req)
}

type MulticlusterTrafficPolicyReconcileLoop interface {
	// AddMulticlusterTrafficPolicyReconciler adds a MulticlusterTrafficPolicyReconciler to the MulticlusterTrafficPolicyReconcileLoop.
	AddMulticlusterTrafficPolicyReconciler(ctx context.Context, rec MulticlusterTrafficPolicyReconciler, predicates ...predicate.Predicate)
}

type multiclusterTrafficPolicyReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterTrafficPolicyReconcileLoop) AddMulticlusterTrafficPolicyReconciler(ctx context.Context, rec MulticlusterTrafficPolicyReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericTrafficPolicyMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterTrafficPolicyReconcileLoop(name string, cw multicluster.ClusterWatcher) MulticlusterTrafficPolicyReconcileLoop {
	return &multiclusterTrafficPolicyReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &networking_smh_solo_io_v1alpha1.TrafficPolicy{})}
}

type genericTrafficPolicyMulticlusterReconciler struct {
	reconciler MulticlusterTrafficPolicyReconciler
}

func (g genericTrafficPolicyMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) {
	if deletionReconciler, ok := g.reconciler.(MulticlusterTrafficPolicyDeletionReconciler); ok {
		deletionReconciler.ReconcileTrafficPolicyDeletion(cluster, req)
	}
}

func (g genericTrafficPolicyMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*networking_smh_solo_io_v1alpha1.TrafficPolicy)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: TrafficPolicy handler received event for %T", object)
	}
	return g.reconciler.ReconcileTrafficPolicy(cluster, obj)
}

// Reconcile Upsert events for the AccessPolicy Resource across clusters.
// implemented by the user
type MulticlusterAccessPolicyReconciler interface {
	ReconcileAccessPolicy(clusterName string, obj *networking_smh_solo_io_v1alpha1.AccessPolicy) (reconcile.Result, error)
}

// Reconcile deletion events for the AccessPolicy Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterAccessPolicyDeletionReconciler interface {
	ReconcileAccessPolicyDeletion(clusterName string, req reconcile.Request)
}

type MulticlusterAccessPolicyReconcilerFuncs struct {
	OnReconcileAccessPolicy         func(clusterName string, obj *networking_smh_solo_io_v1alpha1.AccessPolicy) (reconcile.Result, error)
	OnReconcileAccessPolicyDeletion func(clusterName string, req reconcile.Request)
}

func (f *MulticlusterAccessPolicyReconcilerFuncs) ReconcileAccessPolicy(clusterName string, obj *networking_smh_solo_io_v1alpha1.AccessPolicy) (reconcile.Result, error) {
	if f.OnReconcileAccessPolicy == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileAccessPolicy(clusterName, obj)
}

func (f *MulticlusterAccessPolicyReconcilerFuncs) ReconcileAccessPolicyDeletion(clusterName string, req reconcile.Request) {
	if f.OnReconcileAccessPolicyDeletion == nil {
		return
	}
	f.OnReconcileAccessPolicyDeletion(clusterName, req)
}

type MulticlusterAccessPolicyReconcileLoop interface {
	// AddMulticlusterAccessPolicyReconciler adds a MulticlusterAccessPolicyReconciler to the MulticlusterAccessPolicyReconcileLoop.
	AddMulticlusterAccessPolicyReconciler(ctx context.Context, rec MulticlusterAccessPolicyReconciler, predicates ...predicate.Predicate)
}

type multiclusterAccessPolicyReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterAccessPolicyReconcileLoop) AddMulticlusterAccessPolicyReconciler(ctx context.Context, rec MulticlusterAccessPolicyReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericAccessPolicyMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterAccessPolicyReconcileLoop(name string, cw multicluster.ClusterWatcher) MulticlusterAccessPolicyReconcileLoop {
	return &multiclusterAccessPolicyReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &networking_smh_solo_io_v1alpha1.AccessPolicy{})}
}

type genericAccessPolicyMulticlusterReconciler struct {
	reconciler MulticlusterAccessPolicyReconciler
}

func (g genericAccessPolicyMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) {
	if deletionReconciler, ok := g.reconciler.(MulticlusterAccessPolicyDeletionReconciler); ok {
		deletionReconciler.ReconcileAccessPolicyDeletion(cluster, req)
	}
}

func (g genericAccessPolicyMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*networking_smh_solo_io_v1alpha1.AccessPolicy)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: AccessPolicy handler received event for %T", object)
	}
	return g.reconciler.ReconcileAccessPolicy(cluster, obj)
}

// Reconcile Upsert events for the VirtualMesh Resource across clusters.
// implemented by the user
type MulticlusterVirtualMeshReconciler interface {
	ReconcileVirtualMesh(clusterName string, obj *networking_smh_solo_io_v1alpha1.VirtualMesh) (reconcile.Result, error)
}

// Reconcile deletion events for the VirtualMesh Resource across clusters.
// Deletion receives a reconcile.Request as we cannot guarantee the last state of the object
// before being deleted.
// implemented by the user
type MulticlusterVirtualMeshDeletionReconciler interface {
	ReconcileVirtualMeshDeletion(clusterName string, req reconcile.Request)
}

type MulticlusterVirtualMeshReconcilerFuncs struct {
	OnReconcileVirtualMesh         func(clusterName string, obj *networking_smh_solo_io_v1alpha1.VirtualMesh) (reconcile.Result, error)
	OnReconcileVirtualMeshDeletion func(clusterName string, req reconcile.Request)
}

func (f *MulticlusterVirtualMeshReconcilerFuncs) ReconcileVirtualMesh(clusterName string, obj *networking_smh_solo_io_v1alpha1.VirtualMesh) (reconcile.Result, error) {
	if f.OnReconcileVirtualMesh == nil {
		return reconcile.Result{}, nil
	}
	return f.OnReconcileVirtualMesh(clusterName, obj)
}

func (f *MulticlusterVirtualMeshReconcilerFuncs) ReconcileVirtualMeshDeletion(clusterName string, req reconcile.Request) {
	if f.OnReconcileVirtualMeshDeletion == nil {
		return
	}
	f.OnReconcileVirtualMeshDeletion(clusterName, req)
}

type MulticlusterVirtualMeshReconcileLoop interface {
	// AddMulticlusterVirtualMeshReconciler adds a MulticlusterVirtualMeshReconciler to the MulticlusterVirtualMeshReconcileLoop.
	AddMulticlusterVirtualMeshReconciler(ctx context.Context, rec MulticlusterVirtualMeshReconciler, predicates ...predicate.Predicate)
}

type multiclusterVirtualMeshReconcileLoop struct {
	loop multicluster.Loop
}

func (m *multiclusterVirtualMeshReconcileLoop) AddMulticlusterVirtualMeshReconciler(ctx context.Context, rec MulticlusterVirtualMeshReconciler, predicates ...predicate.Predicate) {
	genericReconciler := genericVirtualMeshMulticlusterReconciler{reconciler: rec}

	m.loop.AddReconciler(ctx, genericReconciler, predicates...)
}

func NewMulticlusterVirtualMeshReconcileLoop(name string, cw multicluster.ClusterWatcher) MulticlusterVirtualMeshReconcileLoop {
	return &multiclusterVirtualMeshReconcileLoop{loop: mc_reconcile.NewLoop(name, cw, &networking_smh_solo_io_v1alpha1.VirtualMesh{})}
}

type genericVirtualMeshMulticlusterReconciler struct {
	reconciler MulticlusterVirtualMeshReconciler
}

func (g genericVirtualMeshMulticlusterReconciler) ReconcileDeletion(cluster string, req reconcile.Request) {
	if deletionReconciler, ok := g.reconciler.(MulticlusterVirtualMeshDeletionReconciler); ok {
		deletionReconciler.ReconcileVirtualMeshDeletion(cluster, req)
	}
}

func (g genericVirtualMeshMulticlusterReconciler) Reconcile(cluster string, object ezkube.Object) (reconcile.Result, error) {
	obj, ok := object.(*networking_smh_solo_io_v1alpha1.VirtualMesh)
	if !ok {
		return reconcile.Result{}, errors.Errorf("internal error: VirtualMesh handler received event for %T", object)
	}
	return g.reconciler.ReconcileVirtualMesh(cluster, obj)
}
