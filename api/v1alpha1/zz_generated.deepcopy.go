//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2021.

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
func (in *ResourcesLimit) DeepCopyInto(out *ResourcesLimit) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourcesLimit.
func (in *ResourcesLimit) DeepCopy() *ResourcesLimit {
	if in == nil {
		return nil
	}
	out := new(ResourcesLimit)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourcesRequest) DeepCopyInto(out *ResourcesRequest) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourcesRequest.
func (in *ResourcesRequest) DeepCopy() *ResourcesRequest {
	if in == nil {
		return nil
	}
	out := new(ResourcesRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TangServer) DeepCopyInto(out *TangServer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TangServer.
func (in *TangServer) DeepCopy() *TangServer {
	if in == nil {
		return nil
	}
	out := new(TangServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TangServer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TangServerActiveKeys) DeepCopyInto(out *TangServerActiveKeys) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TangServerActiveKeys.
func (in *TangServerActiveKeys) DeepCopy() *TangServerActiveKeys {
	if in == nil {
		return nil
	}
	out := new(TangServerActiveKeys)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TangServerHiddenKeys) DeepCopyInto(out *TangServerHiddenKeys) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TangServerHiddenKeys.
func (in *TangServerHiddenKeys) DeepCopy() *TangServerHiddenKeys {
	if in == nil {
		return nil
	}
	out := new(TangServerHiddenKeys)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TangServerList) DeepCopyInto(out *TangServerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]TangServer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TangServerList.
func (in *TangServerList) DeepCopy() *TangServerList {
	if in == nil {
		return nil
	}
	out := new(TangServerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *TangServerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TangServerSpec) DeepCopyInto(out *TangServerSpec) {
	*out = *in
	out.ResourcesRequest = in.ResourcesRequest
	out.ResourcesLimit = in.ResourcesLimit
	if in.HiddenKeys != nil {
		in, out := &in.HiddenKeys, &out.HiddenKeys
		*out = make([]TangServerHiddenKeys, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TangServerSpec.
func (in *TangServerSpec) DeepCopy() *TangServerSpec {
	if in == nil {
		return nil
	}
	out := new(TangServerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TangServerStatus) DeepCopyInto(out *TangServerStatus) {
	*out = *in
	if in.ActiveKeys != nil {
		in, out := &in.ActiveKeys, &out.ActiveKeys
		*out = make([]TangServerActiveKeys, len(*in))
		copy(*out, *in)
	}
	if in.HiddenKeys != nil {
		in, out := &in.HiddenKeys, &out.HiddenKeys
		*out = make([]TangServerHiddenKeys, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TangServerStatus.
func (in *TangServerStatus) DeepCopy() *TangServerStatus {
	if in == nil {
		return nil
	}
	out := new(TangServerStatus)
	in.DeepCopyInto(out)
	return out
}
