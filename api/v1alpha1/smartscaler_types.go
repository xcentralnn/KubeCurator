// api/v1alpha1/smartscaler_types.go

package v1alpha1

import (
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type TargetRef struct {
    APIVersion string `json:"apiVersion"`
    Kind       string `json:"kind"`
    Name       string `json:"name"`
}

type SmartScalerSpec struct {
    TargetRef   TargetRef `json:"targetRef"`
    MinReplicas int32     `json:"minReplicas"`
    MaxReplicas int32     `json:"maxReplicas"`
}

type SmartScalerStatus struct {
    CurrentReplicas int32   `json:"currentReplicas,omitempty"`
    PredictedLoad   float64 `json:"predictedLoad,omitempty"`
}

type SmartScaler struct {
    metav1.TypeMeta   `json:",inline"`
    metav1.ObjectMeta `json:"metadata,omitempty"`

    Spec   SmartScalerSpec   `json:"spec,omitempty"`
    Status SmartScalerStatus `json:"status,omitempty"`
}
