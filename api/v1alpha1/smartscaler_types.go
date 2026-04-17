package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type TargetRef struct {
	APIVersion string `json:"apiVersion"`
	Kind       string `json:"kind"`
	Name       string `json:"name"`
	Namespace  string `json:"namespace,omitempty"`
}

type SmartScalerSpec struct {
	TargetRef TargetRef `json:"targetRef"`

	MinReplicas int32 `json:"minReplicas"`
	MaxReplicas int32 `json:"maxReplicas"`

	TargetCPUUtilization float64 `json:"targetCPUUtilization,omitempty"`

	CooldownSeconds int32 `json:"cooldownSeconds,omitempty"`

	Mode string `json:"mode,omitempty"`
}

type SmartScalerStatus struct {
	CurrentReplicas int32   `json:"currentReplicas,omitempty"`
	DesiredReplicas int32   `json:"desiredReplicas,omitempty"`
	PredictedLoad   float64 `json:"predictedLoad,omitempty"`

	LastScaleTime metav1.Time `json:"lastScaleTime,omitempty"`

	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

type SmartScaler struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SmartScalerSpec   `json:"spec,omitempty"`
	Status SmartScalerStatus `json:"status,omitempty"`
}

type SmartScalerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SmartScaler `json:"items"`
}