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
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// TargetProviderSpec defines the desired state of TargetProvider.
type TargetProviderSpec struct {
	// Image specifies the container image to use for the target provider
	Image string `json:"image"`

	// Config contains the configuration for the target provider
	Config TargetProviderConfig `json:"config"`

	// Template is the pod template specification
	Template corev1.PodTemplateSpec `json:"template,omitempty"`
}

// TargetProviderConfig defines the configuration for the target provider
type TargetProviderConfig struct {
	// Machine specifies the target machine type
	Machine string `json:"machine"`

	// Memory specifies the memory allocation for the target
	Memory string `json:"memory"`
}

// TargetProviderStatus defines the observed state of TargetProvider.
type TargetProviderStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// TargetProvider is the Schema for the targetproviders API.
type TargetProvider struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   TargetProviderSpec   `json:"spec,omitempty"`
	Status TargetProviderStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TargetProviderList contains a list of TargetProvider.
type TargetProviderList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TargetProvider `json:"items"`
}

func init() {
	SchemeBuilder.Register(&TargetProvider{}, &TargetProviderList{})
}
