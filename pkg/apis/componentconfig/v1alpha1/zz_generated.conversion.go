//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2024 The Kubernetes Authors.

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

// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	componentconfig "sigs.k8s.io/descheduler/pkg/apis/componentconfig"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*DeschedulerConfiguration)(nil), (*componentconfig.DeschedulerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_DeschedulerConfiguration_To_componentconfig_DeschedulerConfiguration(a.(*DeschedulerConfiguration), b.(*componentconfig.DeschedulerConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*componentconfig.DeschedulerConfiguration)(nil), (*DeschedulerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_componentconfig_DeschedulerConfiguration_To_v1alpha1_DeschedulerConfiguration(a.(*componentconfig.DeschedulerConfiguration), b.(*DeschedulerConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*TracingConfiguration)(nil), (*componentconfig.TracingConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_TracingConfiguration_To_componentconfig_TracingConfiguration(a.(*TracingConfiguration), b.(*componentconfig.TracingConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*componentconfig.TracingConfiguration)(nil), (*TracingConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_componentconfig_TracingConfiguration_To_v1alpha1_TracingConfiguration(a.(*componentconfig.TracingConfiguration), b.(*TracingConfiguration), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha1_DeschedulerConfiguration_To_componentconfig_DeschedulerConfiguration(in *DeschedulerConfiguration, out *componentconfig.DeschedulerConfiguration, s conversion.Scope) error {
	out.DeschedulingInterval = time.Duration(in.DeschedulingInterval)
	out.KubeconfigFile = in.KubeconfigFile
	out.PolicyConfigFile = in.PolicyConfigFile
	out.DryRun = in.DryRun
	out.NodeSelector = in.NodeSelector
	out.MaxNoOfPodsToEvictPerNode = in.MaxNoOfPodsToEvictPerNode
	out.EvictLocalStoragePods = in.EvictLocalStoragePods
	out.IgnorePVCPods = in.IgnorePVCPods
	if err := Convert_v1alpha1_TracingConfiguration_To_componentconfig_TracingConfiguration(&in.Tracing, &out.Tracing, s); err != nil {
		return err
	}
	out.LeaderElection = in.LeaderElection
	out.ClientConnection = in.ClientConnection
	return nil
}

// Convert_v1alpha1_DeschedulerConfiguration_To_componentconfig_DeschedulerConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_DeschedulerConfiguration_To_componentconfig_DeschedulerConfiguration(in *DeschedulerConfiguration, out *componentconfig.DeschedulerConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_DeschedulerConfiguration_To_componentconfig_DeschedulerConfiguration(in, out, s)
}

func autoConvert_componentconfig_DeschedulerConfiguration_To_v1alpha1_DeschedulerConfiguration(in *componentconfig.DeschedulerConfiguration, out *DeschedulerConfiguration, s conversion.Scope) error {
	out.DeschedulingInterval = time.Duration(in.DeschedulingInterval)
	out.KubeconfigFile = in.KubeconfigFile
	out.PolicyConfigFile = in.PolicyConfigFile
	out.DryRun = in.DryRun
	out.NodeSelector = in.NodeSelector
	out.MaxNoOfPodsToEvictPerNode = in.MaxNoOfPodsToEvictPerNode
	out.EvictLocalStoragePods = in.EvictLocalStoragePods
	out.IgnorePVCPods = in.IgnorePVCPods
	if err := Convert_componentconfig_TracingConfiguration_To_v1alpha1_TracingConfiguration(&in.Tracing, &out.Tracing, s); err != nil {
		return err
	}
	out.LeaderElection = in.LeaderElection
	out.ClientConnection = in.ClientConnection
	return nil
}

// Convert_componentconfig_DeschedulerConfiguration_To_v1alpha1_DeschedulerConfiguration is an autogenerated conversion function.
func Convert_componentconfig_DeschedulerConfiguration_To_v1alpha1_DeschedulerConfiguration(in *componentconfig.DeschedulerConfiguration, out *DeschedulerConfiguration, s conversion.Scope) error {
	return autoConvert_componentconfig_DeschedulerConfiguration_To_v1alpha1_DeschedulerConfiguration(in, out, s)
}

func autoConvert_v1alpha1_TracingConfiguration_To_componentconfig_TracingConfiguration(in *TracingConfiguration, out *componentconfig.TracingConfiguration, s conversion.Scope) error {
	out.CollectorEndpoint = in.CollectorEndpoint
	out.TransportCert = in.TransportCert
	out.ServiceName = in.ServiceName
	out.ServiceNamespace = in.ServiceNamespace
	out.SampleRate = in.SampleRate
	out.FallbackToNoOpProviderOnError = in.FallbackToNoOpProviderOnError
	return nil
}

// Convert_v1alpha1_TracingConfiguration_To_componentconfig_TracingConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_TracingConfiguration_To_componentconfig_TracingConfiguration(in *TracingConfiguration, out *componentconfig.TracingConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_TracingConfiguration_To_componentconfig_TracingConfiguration(in, out, s)
}

func autoConvert_componentconfig_TracingConfiguration_To_v1alpha1_TracingConfiguration(in *componentconfig.TracingConfiguration, out *TracingConfiguration, s conversion.Scope) error {
	out.CollectorEndpoint = in.CollectorEndpoint
	out.TransportCert = in.TransportCert
	out.ServiceName = in.ServiceName
	out.ServiceNamespace = in.ServiceNamespace
	out.SampleRate = in.SampleRate
	out.FallbackToNoOpProviderOnError = in.FallbackToNoOpProviderOnError
	return nil
}

// Convert_componentconfig_TracingConfiguration_To_v1alpha1_TracingConfiguration is an autogenerated conversion function.
func Convert_componentconfig_TracingConfiguration_To_v1alpha1_TracingConfiguration(in *componentconfig.TracingConfiguration, out *TracingConfiguration, s conversion.Scope) error {
	return autoConvert_componentconfig_TracingConfiguration_To_v1alpha1_TracingConfiguration(in, out, s)
}