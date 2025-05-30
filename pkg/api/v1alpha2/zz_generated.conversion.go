//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2025 The Kubernetes Authors.

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

package v1alpha2

import (
	unsafe "unsafe"

	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	api "sigs.k8s.io/descheduler/pkg/api"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*AuthToken)(nil), (*api.AuthToken)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_AuthToken_To_api_AuthToken(a.(*AuthToken), b.(*api.AuthToken), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*api.AuthToken)(nil), (*AuthToken)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_api_AuthToken_To_v1alpha2_AuthToken(a.(*api.AuthToken), b.(*AuthToken), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*DeschedulerProfile)(nil), (*api.DeschedulerProfile)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_DeschedulerProfile_To_api_DeschedulerProfile(a.(*DeschedulerProfile), b.(*api.DeschedulerProfile), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*api.DeschedulerProfile)(nil), (*DeschedulerProfile)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_api_DeschedulerProfile_To_v1alpha2_DeschedulerProfile(a.(*api.DeschedulerProfile), b.(*DeschedulerProfile), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*MetricsCollector)(nil), (*api.MetricsCollector)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_MetricsCollector_To_api_MetricsCollector(a.(*MetricsCollector), b.(*api.MetricsCollector), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*api.MetricsCollector)(nil), (*MetricsCollector)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_api_MetricsCollector_To_v1alpha2_MetricsCollector(a.(*api.MetricsCollector), b.(*MetricsCollector), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*MetricsProvider)(nil), (*api.MetricsProvider)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_MetricsProvider_To_api_MetricsProvider(a.(*MetricsProvider), b.(*api.MetricsProvider), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*api.MetricsProvider)(nil), (*MetricsProvider)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_api_MetricsProvider_To_v1alpha2_MetricsProvider(a.(*api.MetricsProvider), b.(*MetricsProvider), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*api.PluginConfig)(nil), (*PluginConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_api_PluginConfig_To_v1alpha2_PluginConfig(a.(*api.PluginConfig), b.(*PluginConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*PluginSet)(nil), (*api.PluginSet)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_PluginSet_To_api_PluginSet(a.(*PluginSet), b.(*api.PluginSet), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*api.PluginSet)(nil), (*PluginSet)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_api_PluginSet_To_v1alpha2_PluginSet(a.(*api.PluginSet), b.(*PluginSet), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Plugins)(nil), (*api.Plugins)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_Plugins_To_api_Plugins(a.(*Plugins), b.(*api.Plugins), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*api.Plugins)(nil), (*Plugins)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_api_Plugins_To_v1alpha2_Plugins(a.(*api.Plugins), b.(*Plugins), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Prometheus)(nil), (*api.Prometheus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_Prometheus_To_api_Prometheus(a.(*Prometheus), b.(*api.Prometheus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*api.Prometheus)(nil), (*Prometheus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_api_Prometheus_To_v1alpha2_Prometheus(a.(*api.Prometheus), b.(*Prometheus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*SecretReference)(nil), (*api.SecretReference)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_SecretReference_To_api_SecretReference(a.(*SecretReference), b.(*api.SecretReference), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*api.SecretReference)(nil), (*SecretReference)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_api_SecretReference_To_v1alpha2_SecretReference(a.(*api.SecretReference), b.(*SecretReference), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*api.DeschedulerPolicy)(nil), (*DeschedulerPolicy)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_api_DeschedulerPolicy_To_v1alpha2_DeschedulerPolicy(a.(*api.DeschedulerPolicy), b.(*DeschedulerPolicy), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*DeschedulerPolicy)(nil), (*api.DeschedulerPolicy)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_DeschedulerPolicy_To_api_DeschedulerPolicy(a.(*DeschedulerPolicy), b.(*api.DeschedulerPolicy), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*PluginConfig)(nil), (*api.PluginConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_PluginConfig_To_api_PluginConfig(a.(*PluginConfig), b.(*api.PluginConfig), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha2_AuthToken_To_api_AuthToken(in *AuthToken, out *api.AuthToken, s conversion.Scope) error {
	out.SecretReference = (*api.SecretReference)(unsafe.Pointer(in.SecretReference))
	return nil
}

// Convert_v1alpha2_AuthToken_To_api_AuthToken is an autogenerated conversion function.
func Convert_v1alpha2_AuthToken_To_api_AuthToken(in *AuthToken, out *api.AuthToken, s conversion.Scope) error {
	return autoConvert_v1alpha2_AuthToken_To_api_AuthToken(in, out, s)
}

func autoConvert_api_AuthToken_To_v1alpha2_AuthToken(in *api.AuthToken, out *AuthToken, s conversion.Scope) error {
	out.SecretReference = (*SecretReference)(unsafe.Pointer(in.SecretReference))
	return nil
}

// Convert_api_AuthToken_To_v1alpha2_AuthToken is an autogenerated conversion function.
func Convert_api_AuthToken_To_v1alpha2_AuthToken(in *api.AuthToken, out *AuthToken, s conversion.Scope) error {
	return autoConvert_api_AuthToken_To_v1alpha2_AuthToken(in, out, s)
}

func autoConvert_v1alpha2_DeschedulerPolicy_To_api_DeschedulerPolicy(in *DeschedulerPolicy, out *api.DeschedulerPolicy, s conversion.Scope) error {
	if in.Profiles != nil {
		in, out := &in.Profiles, &out.Profiles
		*out = make([]api.DeschedulerProfile, len(*in))
		for i := range *in {
			if err := Convert_v1alpha2_DeschedulerProfile_To_api_DeschedulerProfile(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Profiles = nil
	}
	out.NodeSelector = (*string)(unsafe.Pointer(in.NodeSelector))
	out.MaxNoOfPodsToEvictPerNode = (*uint)(unsafe.Pointer(in.MaxNoOfPodsToEvictPerNode))
	out.MaxNoOfPodsToEvictPerNamespace = (*uint)(unsafe.Pointer(in.MaxNoOfPodsToEvictPerNamespace))
	out.MaxNoOfPodsToEvictTotal = (*uint)(unsafe.Pointer(in.MaxNoOfPodsToEvictTotal))
	out.EvictionFailureEventNotification = (*bool)(unsafe.Pointer(in.EvictionFailureEventNotification))
	out.MetricsCollector = (*api.MetricsCollector)(unsafe.Pointer(in.MetricsCollector))
	out.MetricsProviders = *(*[]api.MetricsProvider)(unsafe.Pointer(&in.MetricsProviders))
	out.GracePeriodSeconds = (*int64)(unsafe.Pointer(in.GracePeriodSeconds))
	return nil
}

func autoConvert_api_DeschedulerPolicy_To_v1alpha2_DeschedulerPolicy(in *api.DeschedulerPolicy, out *DeschedulerPolicy, s conversion.Scope) error {
	if in.Profiles != nil {
		in, out := &in.Profiles, &out.Profiles
		*out = make([]DeschedulerProfile, len(*in))
		for i := range *in {
			if err := Convert_api_DeschedulerProfile_To_v1alpha2_DeschedulerProfile(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.Profiles = nil
	}
	out.NodeSelector = (*string)(unsafe.Pointer(in.NodeSelector))
	out.MaxNoOfPodsToEvictPerNode = (*uint)(unsafe.Pointer(in.MaxNoOfPodsToEvictPerNode))
	out.MaxNoOfPodsToEvictPerNamespace = (*uint)(unsafe.Pointer(in.MaxNoOfPodsToEvictPerNamespace))
	out.MaxNoOfPodsToEvictTotal = (*uint)(unsafe.Pointer(in.MaxNoOfPodsToEvictTotal))
	out.EvictionFailureEventNotification = (*bool)(unsafe.Pointer(in.EvictionFailureEventNotification))
	out.MetricsCollector = (*MetricsCollector)(unsafe.Pointer(in.MetricsCollector))
	out.MetricsProviders = *(*[]MetricsProvider)(unsafe.Pointer(&in.MetricsProviders))
	out.GracePeriodSeconds = (*int64)(unsafe.Pointer(in.GracePeriodSeconds))
	return nil
}

func autoConvert_v1alpha2_DeschedulerProfile_To_api_DeschedulerProfile(in *DeschedulerProfile, out *api.DeschedulerProfile, s conversion.Scope) error {
	out.Name = in.Name
	if in.PluginConfigs != nil {
		in, out := &in.PluginConfigs, &out.PluginConfigs
		*out = make([]api.PluginConfig, len(*in))
		for i := range *in {
			if err := Convert_v1alpha2_PluginConfig_To_api_PluginConfig(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.PluginConfigs = nil
	}
	if err := Convert_v1alpha2_Plugins_To_api_Plugins(&in.Plugins, &out.Plugins, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha2_DeschedulerProfile_To_api_DeschedulerProfile is an autogenerated conversion function.
func Convert_v1alpha2_DeschedulerProfile_To_api_DeschedulerProfile(in *DeschedulerProfile, out *api.DeschedulerProfile, s conversion.Scope) error {
	return autoConvert_v1alpha2_DeschedulerProfile_To_api_DeschedulerProfile(in, out, s)
}

func autoConvert_api_DeschedulerProfile_To_v1alpha2_DeschedulerProfile(in *api.DeschedulerProfile, out *DeschedulerProfile, s conversion.Scope) error {
	out.Name = in.Name
	if in.PluginConfigs != nil {
		in, out := &in.PluginConfigs, &out.PluginConfigs
		*out = make([]PluginConfig, len(*in))
		for i := range *in {
			if err := Convert_api_PluginConfig_To_v1alpha2_PluginConfig(&(*in)[i], &(*out)[i], s); err != nil {
				return err
			}
		}
	} else {
		out.PluginConfigs = nil
	}
	if err := Convert_api_Plugins_To_v1alpha2_Plugins(&in.Plugins, &out.Plugins, s); err != nil {
		return err
	}
	return nil
}

// Convert_api_DeschedulerProfile_To_v1alpha2_DeschedulerProfile is an autogenerated conversion function.
func Convert_api_DeschedulerProfile_To_v1alpha2_DeschedulerProfile(in *api.DeschedulerProfile, out *DeschedulerProfile, s conversion.Scope) error {
	return autoConvert_api_DeschedulerProfile_To_v1alpha2_DeschedulerProfile(in, out, s)
}

func autoConvert_v1alpha2_MetricsCollector_To_api_MetricsCollector(in *MetricsCollector, out *api.MetricsCollector, s conversion.Scope) error {
	out.Enabled = in.Enabled
	return nil
}

// Convert_v1alpha2_MetricsCollector_To_api_MetricsCollector is an autogenerated conversion function.
func Convert_v1alpha2_MetricsCollector_To_api_MetricsCollector(in *MetricsCollector, out *api.MetricsCollector, s conversion.Scope) error {
	return autoConvert_v1alpha2_MetricsCollector_To_api_MetricsCollector(in, out, s)
}

func autoConvert_api_MetricsCollector_To_v1alpha2_MetricsCollector(in *api.MetricsCollector, out *MetricsCollector, s conversion.Scope) error {
	out.Enabled = in.Enabled
	return nil
}

// Convert_api_MetricsCollector_To_v1alpha2_MetricsCollector is an autogenerated conversion function.
func Convert_api_MetricsCollector_To_v1alpha2_MetricsCollector(in *api.MetricsCollector, out *MetricsCollector, s conversion.Scope) error {
	return autoConvert_api_MetricsCollector_To_v1alpha2_MetricsCollector(in, out, s)
}

func autoConvert_v1alpha2_MetricsProvider_To_api_MetricsProvider(in *MetricsProvider, out *api.MetricsProvider, s conversion.Scope) error {
	out.Source = api.MetricsSource(in.Source)
	out.Prometheus = (*api.Prometheus)(unsafe.Pointer(in.Prometheus))
	return nil
}

// Convert_v1alpha2_MetricsProvider_To_api_MetricsProvider is an autogenerated conversion function.
func Convert_v1alpha2_MetricsProvider_To_api_MetricsProvider(in *MetricsProvider, out *api.MetricsProvider, s conversion.Scope) error {
	return autoConvert_v1alpha2_MetricsProvider_To_api_MetricsProvider(in, out, s)
}

func autoConvert_api_MetricsProvider_To_v1alpha2_MetricsProvider(in *api.MetricsProvider, out *MetricsProvider, s conversion.Scope) error {
	out.Source = MetricsSource(in.Source)
	out.Prometheus = (*Prometheus)(unsafe.Pointer(in.Prometheus))
	return nil
}

// Convert_api_MetricsProvider_To_v1alpha2_MetricsProvider is an autogenerated conversion function.
func Convert_api_MetricsProvider_To_v1alpha2_MetricsProvider(in *api.MetricsProvider, out *MetricsProvider, s conversion.Scope) error {
	return autoConvert_api_MetricsProvider_To_v1alpha2_MetricsProvider(in, out, s)
}

func autoConvert_v1alpha2_PluginConfig_To_api_PluginConfig(in *PluginConfig, out *api.PluginConfig, s conversion.Scope) error {
	out.Name = in.Name
	if err := runtime.Convert_runtime_RawExtension_To_runtime_Object(&in.Args, &out.Args, s); err != nil {
		return err
	}
	return nil
}

func autoConvert_api_PluginConfig_To_v1alpha2_PluginConfig(in *api.PluginConfig, out *PluginConfig, s conversion.Scope) error {
	out.Name = in.Name
	if err := runtime.Convert_runtime_Object_To_runtime_RawExtension(&in.Args, &out.Args, s); err != nil {
		return err
	}
	return nil
}

// Convert_api_PluginConfig_To_v1alpha2_PluginConfig is an autogenerated conversion function.
func Convert_api_PluginConfig_To_v1alpha2_PluginConfig(in *api.PluginConfig, out *PluginConfig, s conversion.Scope) error {
	return autoConvert_api_PluginConfig_To_v1alpha2_PluginConfig(in, out, s)
}

func autoConvert_v1alpha2_PluginSet_To_api_PluginSet(in *PluginSet, out *api.PluginSet, s conversion.Scope) error {
	out.Enabled = *(*[]string)(unsafe.Pointer(&in.Enabled))
	out.Disabled = *(*[]string)(unsafe.Pointer(&in.Disabled))
	return nil
}

// Convert_v1alpha2_PluginSet_To_api_PluginSet is an autogenerated conversion function.
func Convert_v1alpha2_PluginSet_To_api_PluginSet(in *PluginSet, out *api.PluginSet, s conversion.Scope) error {
	return autoConvert_v1alpha2_PluginSet_To_api_PluginSet(in, out, s)
}

func autoConvert_api_PluginSet_To_v1alpha2_PluginSet(in *api.PluginSet, out *PluginSet, s conversion.Scope) error {
	out.Enabled = *(*[]string)(unsafe.Pointer(&in.Enabled))
	out.Disabled = *(*[]string)(unsafe.Pointer(&in.Disabled))
	return nil
}

// Convert_api_PluginSet_To_v1alpha2_PluginSet is an autogenerated conversion function.
func Convert_api_PluginSet_To_v1alpha2_PluginSet(in *api.PluginSet, out *PluginSet, s conversion.Scope) error {
	return autoConvert_api_PluginSet_To_v1alpha2_PluginSet(in, out, s)
}

func autoConvert_v1alpha2_Plugins_To_api_Plugins(in *Plugins, out *api.Plugins, s conversion.Scope) error {
	if err := Convert_v1alpha2_PluginSet_To_api_PluginSet(&in.PreSort, &out.PreSort, s); err != nil {
		return err
	}
	if err := Convert_v1alpha2_PluginSet_To_api_PluginSet(&in.Sort, &out.Sort, s); err != nil {
		return err
	}
	if err := Convert_v1alpha2_PluginSet_To_api_PluginSet(&in.Deschedule, &out.Deschedule, s); err != nil {
		return err
	}
	if err := Convert_v1alpha2_PluginSet_To_api_PluginSet(&in.Balance, &out.Balance, s); err != nil {
		return err
	}
	if err := Convert_v1alpha2_PluginSet_To_api_PluginSet(&in.Filter, &out.Filter, s); err != nil {
		return err
	}
	if err := Convert_v1alpha2_PluginSet_To_api_PluginSet(&in.PreEvictionFilter, &out.PreEvictionFilter, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha2_Plugins_To_api_Plugins is an autogenerated conversion function.
func Convert_v1alpha2_Plugins_To_api_Plugins(in *Plugins, out *api.Plugins, s conversion.Scope) error {
	return autoConvert_v1alpha2_Plugins_To_api_Plugins(in, out, s)
}

func autoConvert_api_Plugins_To_v1alpha2_Plugins(in *api.Plugins, out *Plugins, s conversion.Scope) error {
	if err := Convert_api_PluginSet_To_v1alpha2_PluginSet(&in.PreSort, &out.PreSort, s); err != nil {
		return err
	}
	if err := Convert_api_PluginSet_To_v1alpha2_PluginSet(&in.Sort, &out.Sort, s); err != nil {
		return err
	}
	if err := Convert_api_PluginSet_To_v1alpha2_PluginSet(&in.Deschedule, &out.Deschedule, s); err != nil {
		return err
	}
	if err := Convert_api_PluginSet_To_v1alpha2_PluginSet(&in.Balance, &out.Balance, s); err != nil {
		return err
	}
	if err := Convert_api_PluginSet_To_v1alpha2_PluginSet(&in.Filter, &out.Filter, s); err != nil {
		return err
	}
	if err := Convert_api_PluginSet_To_v1alpha2_PluginSet(&in.PreEvictionFilter, &out.PreEvictionFilter, s); err != nil {
		return err
	}
	return nil
}

// Convert_api_Plugins_To_v1alpha2_Plugins is an autogenerated conversion function.
func Convert_api_Plugins_To_v1alpha2_Plugins(in *api.Plugins, out *Plugins, s conversion.Scope) error {
	return autoConvert_api_Plugins_To_v1alpha2_Plugins(in, out, s)
}

func autoConvert_v1alpha2_Prometheus_To_api_Prometheus(in *Prometheus, out *api.Prometheus, s conversion.Scope) error {
	out.URL = in.URL
	out.AuthToken = (*api.AuthToken)(unsafe.Pointer(in.AuthToken))
	return nil
}

// Convert_v1alpha2_Prometheus_To_api_Prometheus is an autogenerated conversion function.
func Convert_v1alpha2_Prometheus_To_api_Prometheus(in *Prometheus, out *api.Prometheus, s conversion.Scope) error {
	return autoConvert_v1alpha2_Prometheus_To_api_Prometheus(in, out, s)
}

func autoConvert_api_Prometheus_To_v1alpha2_Prometheus(in *api.Prometheus, out *Prometheus, s conversion.Scope) error {
	out.URL = in.URL
	out.AuthToken = (*AuthToken)(unsafe.Pointer(in.AuthToken))
	return nil
}

// Convert_api_Prometheus_To_v1alpha2_Prometheus is an autogenerated conversion function.
func Convert_api_Prometheus_To_v1alpha2_Prometheus(in *api.Prometheus, out *Prometheus, s conversion.Scope) error {
	return autoConvert_api_Prometheus_To_v1alpha2_Prometheus(in, out, s)
}

func autoConvert_v1alpha2_SecretReference_To_api_SecretReference(in *SecretReference, out *api.SecretReference, s conversion.Scope) error {
	out.Namespace = in.Namespace
	out.Name = in.Name
	return nil
}

// Convert_v1alpha2_SecretReference_To_api_SecretReference is an autogenerated conversion function.
func Convert_v1alpha2_SecretReference_To_api_SecretReference(in *SecretReference, out *api.SecretReference, s conversion.Scope) error {
	return autoConvert_v1alpha2_SecretReference_To_api_SecretReference(in, out, s)
}

func autoConvert_api_SecretReference_To_v1alpha2_SecretReference(in *api.SecretReference, out *SecretReference, s conversion.Scope) error {
	out.Namespace = in.Namespace
	out.Name = in.Name
	return nil
}

// Convert_api_SecretReference_To_v1alpha2_SecretReference is an autogenerated conversion function.
func Convert_api_SecretReference_To_v1alpha2_SecretReference(in *api.SecretReference, out *SecretReference, s conversion.Scope) error {
	return autoConvert_api_SecretReference_To_v1alpha2_SecretReference(in, out, s)
}
