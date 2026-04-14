package controller

import (
	"context"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	appsv1alpha1 "github.com/mikal-zh/hello-operator/api/v1alpha1"
)

// HelloAppReconciler reconciles a HelloApp object
type HelloAppReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=apps.intern.dev,resources=helloapps,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=apps.intern.dev,resources=helloapps/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=apps.intern.dev,resources=helloapps/finalizers,verbs=update

func (r *HelloAppReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	logger := log.FromContext(ctx)

	var helloApp appsv1alpha1.HelloApp
	if err := r.Get(ctx, req.NamespacedName, &helloApp); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}
	logger.Info("Successfully fetched HelloApp", "name", helloApp.Name, "namespace", helloApp.Namespace)
	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *HelloAppReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&appsv1alpha1.HelloApp{}).
		Complete(r)
}
