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
	"k8s.io/apimachinery/pkg/runtime"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// VirtualTargetSpec defines the desired state of VirtualTarget.
type VirtualTargetSpec struct {
	// TargetProviderRef specifies the target provider to use for management of this virtual target
	TargetProviderRef corev1.ObjectReference `json:"targetProviderRef"`

	// Type specifies the type of virtual target as understood by the provider
	// +kubebuilder:validation:Required
	Type string `json:"type"`

	// Config contains provider-specific configuration for the virtual target
	// +kubebuilder:validation:Required
	Config runtime.RawExtension `json:"config"`
}

// VirtualTargetStatus defines the observed state of VirtualTarget.
type VirtualTargetStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// VirtualTarget is the Schema for the virtualtargets API.
type VirtualTarget struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VirtualTargetSpec   `json:"spec,omitempty"`
	Status VirtualTargetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualTargetList contains a list of VirtualTarget.
type VirtualTargetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualTarget `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VirtualTarget{}, &VirtualTargetList{})
}
