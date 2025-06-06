//go:build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QemuTargetProvider) DeepCopyInto(out *QemuTargetProvider) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QemuTargetProvider.
func (in *QemuTargetProvider) DeepCopy() *QemuTargetProvider {
	if in == nil {
		return nil
	}
	out := new(QemuTargetProvider)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *QemuTargetProvider) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QemuTargetProviderList) DeepCopyInto(out *QemuTargetProviderList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]QemuTargetProvider, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QemuTargetProviderList.
func (in *QemuTargetProviderList) DeepCopy() *QemuTargetProviderList {
	if in == nil {
		return nil
	}
	out := new(QemuTargetProviderList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *QemuTargetProviderList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QemuTargetProviderSpec) DeepCopyInto(out *QemuTargetProviderSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QemuTargetProviderSpec.
func (in *QemuTargetProviderSpec) DeepCopy() *QemuTargetProviderSpec {
	if in == nil {
		return nil
	}
	out := new(QemuTargetProviderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QemuTargetProviderStatus) DeepCopyInto(out *QemuTargetProviderStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QemuTargetProviderStatus.
func (in *QemuTargetProviderStatus) DeepCopy() *QemuTargetProviderStatus {
	if in == nil {
		return nil
	}
	out := new(QemuTargetProviderStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QemuVirtualTarget) DeepCopyInto(out *QemuVirtualTarget) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QemuVirtualTarget.
func (in *QemuVirtualTarget) DeepCopy() *QemuVirtualTarget {
	if in == nil {
		return nil
	}
	out := new(QemuVirtualTarget)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *QemuVirtualTarget) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QemuVirtualTargetList) DeepCopyInto(out *QemuVirtualTargetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]QemuVirtualTarget, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QemuVirtualTargetList.
func (in *QemuVirtualTargetList) DeepCopy() *QemuVirtualTargetList {
	if in == nil {
		return nil
	}
	out := new(QemuVirtualTargetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *QemuVirtualTargetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QemuVirtualTargetSpec) DeepCopyInto(out *QemuVirtualTargetSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QemuVirtualTargetSpec.
func (in *QemuVirtualTargetSpec) DeepCopy() *QemuVirtualTargetSpec {
	if in == nil {
		return nil
	}
	out := new(QemuVirtualTargetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *QemuVirtualTargetStatus) DeepCopyInto(out *QemuVirtualTargetStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new QemuVirtualTargetStatus.
func (in *QemuVirtualTargetStatus) DeepCopy() *QemuVirtualTargetStatus {
	if in == nil {
		return nil
	}
	out := new(QemuVirtualTargetStatus)
	in.DeepCopyInto(out)
	return out
}
