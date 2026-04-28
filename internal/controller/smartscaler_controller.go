package controller

import (
	"context"
	"math"
	"time"

	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"

	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	curatorv1 "github.com/xcentralnn/curator/api/v1alpha1"
)

// +kubebuilder:rbac:groups=apps,resources=deployments,verbs=get;list;watch;update;patch
// +kubebuilder:rbac:groups=curator.dev,resources=smartscalers,verbs=get;list;watch;update;patch
// +kubebuilder:rbac:groups=curator.dev,resources=smartscalers/status,verbs=get;update;patch

type SmartScalerReconciler struct {
	client.Client
}

func (r *SmartScalerReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := ctrl.LoggerFrom(ctx)

	var scaler curatorv1.SmartScaler
	if err := r.Get(ctx, req.NamespacedName, &scaler); err != nil {
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// Resolve namespace
	ns := scaler.Spec.TargetRef.Namespace
	if ns == "" {
		ns = req.Namespace
	}

	// Fetch Deployment
	var deploy appsv1.Deployment
	key := types.NamespacedName{
		Name:      scaler.Spec.TargetRef.Name,
		Namespace: ns,
	}

	if err := r.Get(ctx, key, &deploy); err != nil {
		log.Error(err, "failed to get target deployment")
		return ctrl.Result{RequeueAfter: 30 * time.Second}, nil
	}

	// Current replicas
	var current int32 = 1
	if deploy.Spec.Replicas != nil {
		current = *deploy.Spec.Replicas
	}

	// ---- Metrics + Prediction ----
	cpu := getMockCPU()
	predicted := predict(cpu)

	// ---- Compute desired ----
	target := scaler.Spec.TargetCPUUtilization
	if target == 0 {
		target = 0.6
	}

	desired := int32(math.Ceil(predicted / target))

	// Clamp
	if desired < scaler.Spec.MinReplicas {
		desired = scaler.Spec.MinReplicas
	}
	if desired > scaler.Spec.MaxReplicas {
		desired = scaler.Spec.MaxReplicas
	}

	// ---- Idempotency ----
	if current == desired {
		log.Info("no scaling required", "replicas", current)

		r.updateStatus(ctx, &scaler, current, desired, predicted)

		return ctrl.Result{RequeueAfter: 30 * time.Second}, nil
	}

	// ---- Cooldown ----
	if !scaler.Status.LastScaleTime.IsZero() && scaler.Spec.CooldownSeconds > 0 {
		elapsed := time.Since(scaler.Status.LastScaleTime.Time)
		if elapsed < time.Duration(scaler.Spec.CooldownSeconds)*time.Second {
			log.Info("cooldown active")
			return ctrl.Result{RequeueAfter: 10 * time.Second}, nil
		}
	}

	// ---- Mode: recommend only ----
	if scaler.Spec.Mode == "recommend" {
		log.Info("recommend mode, skipping scale")

		r.updateStatus(ctx, &scaler, current, desired, predicted)

		return ctrl.Result{RequeueAfter: 30 * time.Second}, nil
	}

	// ---- Apply scaling (safe patch) ----
	patch := client.MergeFrom(deploy.DeepCopy())
	deploy.Spec.Replicas = &desired

	if err := r.Patch(ctx, &deploy, patch); err != nil {
		log.Error(err, "failed to patch deployment")
		return ctrl.Result{}, err
	}

	log.Info("scaled deployment", "from", current, "to", desired)

	// ---- Update status ----
	scaler.Status.LastScaleTime = metav1.Now()
	r.updateStatus(ctx, &scaler, current, desired, predicted)

	return ctrl.Result{RequeueAfter: 30 * time.Second}, nil
}

func (r *SmartScalerReconciler) updateStatus(
	ctx context.Context,
	scaler *curatorv1.SmartScaler,
	current, desired int32,
	predicted float64,
) {
	scaler.Status.CurrentReplicas = current
	scaler.Status.DesiredReplicas = desired
	scaler.Status.PredictedLoad = predicted

	_ = r.Status().Update(ctx, scaler)
}

// ---- Mock ML ----

func getMockCPU() float64 {
	return 0.75
}

func predict(x float64) float64 {
	return x * 1.1
}

// ---- Setup ----

func (r *SmartScalerReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&curatorv1.SmartScaler{}).
		Complete(r)
}
