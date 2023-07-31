//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The cert-manager Authors.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
	v1 "k8s.io/component-base/logs/api/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerConfiguration) DeepCopyInto(out *ControllerConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.KubernetesAPIQPS != nil {
		in, out := &in.KubernetesAPIQPS, &out.KubernetesAPIQPS
		*out = new(float32)
		**out = **in
	}
	if in.KubernetesAPIBurst != nil {
		in, out := &in.KubernetesAPIBurst, &out.KubernetesAPIBurst
		*out = new(int32)
		**out = **in
	}
	if in.LeaderElect != nil {
		in, out := &in.LeaderElect, &out.LeaderElect
		*out = new(bool)
		**out = **in
	}
	if in.Controllers != nil {
		in, out := &in.Controllers, &out.Controllers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ACMEHTTP01SolverRunAsNonRoot != nil {
		in, out := &in.ACMEHTTP01SolverRunAsNonRoot, &out.ACMEHTTP01SolverRunAsNonRoot
		*out = new(bool)
		**out = **in
	}
	if in.ACMEHTTP01SolverNameservers != nil {
		in, out := &in.ACMEHTTP01SolverNameservers, &out.ACMEHTTP01SolverNameservers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ClusterIssuerAmbientCredentials != nil {
		in, out := &in.ClusterIssuerAmbientCredentials, &out.ClusterIssuerAmbientCredentials
		*out = new(bool)
		**out = **in
	}
	if in.IssuerAmbientCredentials != nil {
		in, out := &in.IssuerAmbientCredentials, &out.IssuerAmbientCredentials
		*out = new(bool)
		**out = **in
	}
	if in.DefaultAutoCertificateAnnotations != nil {
		in, out := &in.DefaultAutoCertificateAnnotations, &out.DefaultAutoCertificateAnnotations
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.DNS01RecursiveNameservers != nil {
		in, out := &in.DNS01RecursiveNameservers, &out.DNS01RecursiveNameservers
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.DNS01RecursiveNameserversOnly != nil {
		in, out := &in.DNS01RecursiveNameserversOnly, &out.DNS01RecursiveNameserversOnly
		*out = new(bool)
		**out = **in
	}
	if in.EnableCertificateOwnerRef != nil {
		in, out := &in.EnableCertificateOwnerRef, &out.EnableCertificateOwnerRef
		*out = new(bool)
		**out = **in
	}
	if in.NumberOfConcurrentWorkers != nil {
		in, out := &in.NumberOfConcurrentWorkers, &out.NumberOfConcurrentWorkers
		*out = new(int32)
		**out = **in
	}
	if in.MaxConcurrentChallenges != nil {
		in, out := &in.MaxConcurrentChallenges, &out.MaxConcurrentChallenges
		*out = new(int32)
		**out = **in
	}
	if in.EnablePprof != nil {
		in, out := &in.EnablePprof, &out.EnablePprof
		*out = new(bool)
		**out = **in
	}
	if in.Logging != nil {
		in, out := &in.Logging, &out.Logging
		*out = new(v1.LoggingConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.CopiedAnnotationPrefixes != nil {
		in, out := &in.CopiedAnnotationPrefixes, &out.CopiedAnnotationPrefixes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.FeatureGates != nil {
		in, out := &in.FeatureGates, &out.FeatureGates
		*out = make(map[string]bool, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerConfiguration.
func (in *ControllerConfiguration) DeepCopy() *ControllerConfiguration {
	if in == nil {
		return nil
	}
	out := new(ControllerConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ControllerConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}
