// internal/controller/smartscaler_controller.go

package controller

import (
    "context"
    "math"

    appsv1 "k8s.io/api/apps/v1"
    "k8s.io/apimachinery/pkg/types"
    ctrl "sigs.k8s.io/controller-runtime"
    "sigs.k8s.io/controller-runtime/pkg/client"

    aiopsv1 "github.com/xcentralnn/kubecurator/api/v1alpha1"
)

type SmartScalerReconciler struct {
    client.Client
}

func (r *SmartScalerReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {

    var scaler aiopsv1.SmartScaler
    if err := r.Get(ctx, req.NamespacedName, &scaler); err != nil {
        return ctrl.Result{}, client.IgnoreNotFound(err)
    }

    // 1. Fetch target deployment
    var deploy appsv1.Deployment
    key := types.NamespacedName{
        Name:      scaler.Spec.TargetRef.Name,
        Namespace: req.Namespace,
    }

    if err := r.Get(ctx, key, &deploy); err != nil {
        return ctrl.Result{}, err
    }

    // 2. Mock metrics (replace with Prometheus later)
    cpuUsage := getMockCPU()

    // 3. Predict (placeholder ML)
    predicted := predict(cpuUsage)

    // 4. Compute replicas
    targetCPU := 0.6
    replicas := int32(math.Ceil(predicted / targetCPU))

    if replicas < scaler.Spec.MinReplicas {
        replicas = scaler.Spec.MinReplicas
    }
    if replicas > scaler.Spec.MaxReplicas {
        replicas = scaler.Spec.MaxReplicas
    }

    // 5. Apply scaling
    deploy.Spec.Replicas = &replicas
    if err := r.Update(ctx, &deploy); err != nil {
        return ctrl.Result{}, err
    }

    // 6. Update status
    scaler.Status.CurrentReplicas = replicas
    scaler.Status.PredictedLoad = predicted
    _ = r.Status().Update(ctx, &scaler)

    return ctrl.Result{}, nil
}

// ---- Mock ML ----

func getMockCPU() float64 {
    return 0.75
}

func predict(x float64) float64 {
    // placeholder for ML model
    return x * 1.1
}

// ---- Setup ----

func (r *SmartScalerReconciler) SetupWithManager(mgr ctrl.Manager) error {
    return ctrl.NewControllerManagedBy(mgr).
        For(&aiopsv1.SmartScaler{}).
        Complete(r)
}
