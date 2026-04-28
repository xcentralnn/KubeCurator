package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

func (in *SmartScaler) DeepCopyInto(out *SmartScaler) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

func (in *SmartScaler) DeepCopy() *SmartScaler {
	if in == nil {
		return nil
	}
	out := new(SmartScaler)
	in.DeepCopyInto(out)
	return out
}

func (in *SmartScaler) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *SmartScalerList) DeepCopyInto(out *SmartScalerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SmartScaler, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

func (in *SmartScalerList) DeepCopy() *SmartScalerList {
	if in == nil {
		return nil
	}
	out := new(SmartScalerList)
	in.DeepCopyInto(out)
	return out
}

func (in *SmartScalerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

func (in *SmartScalerStatus) DeepCopyInto(out *SmartScalerStatus) {
	*out = *in
	in.LastScaleTime.DeepCopyInto(&out.LastScaleTime)
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

func (in *SmartScalerStatus) DeepCopy() *SmartScalerStatus {
	if in == nil {
		return nil
	}
	out := new(SmartScalerStatus)
	in.DeepCopyInto(out)
	return out
}
