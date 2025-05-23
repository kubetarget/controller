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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// VirtualTargetProviderSpec defines the desired state of VirtualTargetProvider
type VirtualTargetProviderSpec struct {
	// Image specifies the container image for the provider
	Image string `json:"image"`

	// ProviderType specifies the type of provider (e.g., qemu, corellium, avd)
	ProviderType string `json:"providerType"`

	// Configuration contains provider-specific configuration
	Configuration ProviderConfiguration `json:"configuration,omitempty"`
}

// ProviderConfiguration defines the configuration for the provider
type ProviderConfiguration struct {
	// Credentials contains provider-specific credentials
	Credentials *CredentialsConfig `json:"credentials,omitempty"`

	// Resources defines resource limits and requests
	Resources *ResourceConfig `json:"resources,omitempty"`

	// Additional provider-specific configuration
	AdditionalConfig map[string]string `json:"additionalConfig,omitempty"`
}

// CredentialsConfig defines credentials configuration
type CredentialsConfig struct {
	// SecretRef is a reference to a Secret containing credentials
	SecretRef string `json:"secretRef,omitempty"`
}

// ResourceConfig defines resource configuration
type ResourceConfig struct {
	// CPU resource configuration
	CPU CPUConfig `json:"cpu,omitempty"`

	// Memory resource configuration
	Memory MemoryConfig `json:"memory,omitempty"`
}

// CPUConfig defines CPU resource configuration
type CPUConfig struct {
	// Request specifies the CPU request
	Request string `json:"request,omitempty"`

	// Limit specifies the CPU limit
	Limit string `json:"limit,omitempty"`
}

// MemoryConfig defines memory resource configuration
type MemoryConfig struct {
	// Request specifies the memory request
	Request string `json:"request,omitempty"`

	// Limit specifies the memory limit
	Limit string `json:"limit,omitempty"`
}

// VirtualTargetProviderStatus defines the observed state of VirtualTargetProvider
type VirtualTargetProviderStatus struct {
	// Phase represents the current phase of the provider (Pending, Running, Failed)
	Phase string `json:"phase,omitempty"`

	// Conditions represent the latest available observations of the provider's current state
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// SupportedHardwareTypes lists the hardware types supported by this provider
	SupportedHardwareTypes []string `json:"supportedHardwareTypes,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// VirtualTargetProvider is the Schema for the virtualtargetproviders API.
type VirtualTargetProvider struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VirtualTargetProviderSpec   `json:"spec,omitempty"`
	Status VirtualTargetProviderStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualTargetProviderList contains a list of VirtualTargetProvider.
type VirtualTargetProviderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualTargetProvider `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VirtualTargetProvider{}, &VirtualTargetProviderList{})
}
