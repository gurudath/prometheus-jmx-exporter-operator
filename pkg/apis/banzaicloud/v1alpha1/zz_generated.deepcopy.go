// +build !ignore_autogenerated

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrometheusJmxExporter) DeepCopyInto(out *PrometheusJmxExporter) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrometheusJmxExporter.
func (in *PrometheusJmxExporter) DeepCopy() *PrometheusJmxExporter {
	if in == nil {
		return nil
	}
	out := new(PrometheusJmxExporter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PrometheusJmxExporter) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrometheusJmxExporterList) DeepCopyInto(out *PrometheusJmxExporterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PrometheusJmxExporter, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrometheusJmxExporterList.
func (in *PrometheusJmxExporterList) DeepCopy() *PrometheusJmxExporterList {
	if in == nil {
		return nil
	}
	out := new(PrometheusJmxExporterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PrometheusJmxExporterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrometheusJmxExporterSpec) DeepCopyInto(out *PrometheusJmxExporterSpec) {
	*out = *in
	if in.LabelSelector != nil {
		in, out := &in.LabelSelector, &out.LabelSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrometheusJmxExporterSpec.
func (in *PrometheusJmxExporterSpec) DeepCopy() *PrometheusJmxExporterSpec {
	if in == nil {
		return nil
	}
	out := new(PrometheusJmxExporterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrometheusJmxExporterStatus) DeepCopyInto(out *PrometheusJmxExporterStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrometheusJmxExporterStatus.
func (in *PrometheusJmxExporterStatus) DeepCopy() *PrometheusJmxExporterStatus {
	if in == nil {
		return nil
	}
	out := new(PrometheusJmxExporterStatus)
	in.DeepCopyInto(out)
	return out
}
