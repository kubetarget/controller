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

// VirtualTargetDeviceSpec defines the desired state of VirtualTargetDevice.
type VirtualTargetDeviceSpec struct {
	// ProviderRef is the name of the VirtualTargetProvider to use for this device
	ProviderRef string `json:"providerRef"`

	// HardwareType specifies the type of hardware to emulate (e.g., stm32f4)
	HardwareType string `json:"hardwareType"`

	// Configuration contains the device-specific configuration
	Configuration DeviceConfiguration `json:"configuration"`
}

// DeviceConfiguration defines the configuration for the virtual device
type DeviceConfiguration struct {
	// Peripherals defines the available peripherals for the device
	Peripherals PeripheralsConfig `json:"peripherals,omitempty"`

	// Firmware configuration for the device
	Firmware FirmwareConfig `json:"firmware,omitempty"`
}

// PeripheralsConfig defines the peripheral configuration
type PeripheralsConfig struct {
	// CAN bus configuration
	CAN CANConfig `json:"can,omitempty"`

	// UART configuration
	UART UARTConfig `json:"uart,omitempty"`
}

// CANConfig defines CAN bus configuration
type CANConfig struct {
	// Count specifies the number of CAN interfaces
	Count int `json:"count,omitempty"`
}

// UARTConfig defines UART configuration
type UARTConfig struct {
	// Count specifies the number of UART interfaces
	Count int `json:"count,omitempty"`
}

// FirmwareConfig defines firmware configuration
type FirmwareConfig struct {
	// ConfigMap reference containing the firmware
	ConfigMap string `json:"configMap,omitempty"`
}

// VirtualTargetDeviceStatus defines the observed state of VirtualTargetDevice.
type VirtualTargetDeviceStatus struct {
	// Phase represents the current phase of the device (Pending, Running, Failed)
	Phase string `json:"phase,omitempty"`

	// Conditions represent the latest available observations of the device's current state
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// ProviderStatus contains provider-specific status information
	ProviderStatus map[string]string `json:"providerStatus,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// VirtualTargetDevice is the Schema for the virtualtargetdevices API.
type VirtualTargetDevice struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VirtualTargetDeviceSpec   `json:"spec,omitempty"`
	Status VirtualTargetDeviceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualTargetDeviceList contains a list of VirtualTargetDevice.
type VirtualTargetDeviceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualTargetDevice `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VirtualTargetDevice{}, &VirtualTargetDeviceList{})
}
