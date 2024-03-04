// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type MonitorInitParameters struct {

	// (String) The monitor document
	// The monitor document
	Body *string `json:"body,omitempty" tf:"body,omitempty"`
}

type MonitorObservation struct {

	// (String) The monitor document
	// The monitor document
	Body *string `json:"body,omitempty" tf:"body,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`
}

type MonitorParameters struct {

	// (String) The monitor document
	// The monitor document
	// +kubebuilder:validation:Optional
	Body *string `json:"body,omitempty" tf:"body,omitempty"`
}

// MonitorSpec defines the desired state of Monitor
type MonitorSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MonitorParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider MonitorInitParameters `json:"initProvider,omitempty"`
}

// MonitorStatus defines the observed state of Monitor.
type MonitorStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MonitorObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Monitor is the Schema for the Monitors API. Provides an OpenSearch monitor. Please refer to the OpenSearch monitor documentation for details.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,opensearch}
type Monitor struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.body) || (has(self.initProvider) && has(self.initProvider.body))",message="spec.forProvider.body is a required parameter"
	Spec   MonitorSpec   `json:"spec"`
	Status MonitorStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MonitorList contains a list of Monitors
type MonitorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Monitor `json:"items"`
}

// Repository type metadata.
var (
	Monitor_Kind             = "Monitor"
	Monitor_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Monitor_Kind}.String()
	Monitor_KindAPIVersion   = Monitor_Kind + "." + CRDGroupVersion.String()
	Monitor_GroupVersionKind = CRDGroupVersion.WithKind(Monitor_Kind)
)

func init() {
	SchemeBuilder.Register(&Monitor{}, &MonitorList{})
}
