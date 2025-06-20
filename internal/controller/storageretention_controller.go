/*
Copyright 2025.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/log"

	infraalphav1 "github.com/colinbruner/gcs-storage-controller/api/alphav1"
)

// StorageRetentionReconciler reconciles a StorageRetention object
type StorageRetentionReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=infra.colinbruner.com,resources=storageretentions,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=infra.colinbruner.com,resources=storageretentions/status,verbs=get;update;patch
// +kubebuilder:rbac:groups=infra.colinbruner.com,resources=storageretentions/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the StorageRetention object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.21.0/pkg/reconcile
func (r *StorageRetentionReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := logf.FromContext(ctx)

	var st infraalphav1.StorageRetention
	if err := r.Get(ctx, req.NamespacedName, &st); err != nil {
		log.Error(err, "unable to fetch StorageRetention")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	} else {
		log.Info("Reconciled StorageRetention", "name", st.Name, "namespace", st.Namespace)
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *StorageRetentionReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&infraalphav1.StorageRetention{}).
		Named("storageretention").
		Complete(r)
}
