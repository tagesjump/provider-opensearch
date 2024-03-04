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

type ScriptInitParameters struct {

	// (String) Specifies the language the script is written in. Defaults to painless.
	// Specifies the language the script is written in. Defaults to painless.
	Lang *string `json:"lang,omitempty" tf:"lang,omitempty"`

	// (String) Identifier for the stored script. Must be unique within the cluster.
	// Identifier for the stored script. Must be unique within the cluster.
	ScriptID *string `json:"scriptId,omitempty" tf:"script_id,omitempty"`

	// (String) The source of the stored script
	// The source of the stored script
	Source *string `json:"source,omitempty" tf:"source,omitempty"`
}

type ScriptObservation struct {

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) Specifies the language the script is written in. Defaults to painless.
	// Specifies the language the script is written in. Defaults to painless.
	Lang *string `json:"lang,omitempty" tf:"lang,omitempty"`

	// (String) Identifier for the stored script. Must be unique within the cluster.
	// Identifier for the stored script. Must be unique within the cluster.
	ScriptID *string `json:"scriptId,omitempty" tf:"script_id,omitempty"`

	// (String) The source of the stored script
	// The source of the stored script
	Source *string `json:"source,omitempty" tf:"source,omitempty"`
}

type ScriptParameters struct {

	// (String) Specifies the language the script is written in. Defaults to painless.
	// Specifies the language the script is written in. Defaults to painless.
	// +kubebuilder:validation:Optional
	Lang *string `json:"lang,omitempty" tf:"lang,omitempty"`

	// (String) Identifier for the stored script. Must be unique within the cluster.
	// Identifier for the stored script. Must be unique within the cluster.
	// +kubebuilder:validation:Optional
	ScriptID *string `json:"scriptId,omitempty" tf:"script_id,omitempty"`

	// (String) The source of the stored script
	// The source of the stored script
	// +kubebuilder:validation:Optional
	Source *string `json:"source,omitempty" tf:"source,omitempty"`
}

// ScriptSpec defines the desired state of Script
type ScriptSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ScriptParameters `json:"forProvider"`
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
	InitProvider ScriptInitParameters `json:"initProvider,omitempty"`
}

// ScriptStatus defines the observed state of Script.
type ScriptStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ScriptObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Script is the Schema for the Scripts API. Provides an OpenSearch script resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,opensearch}
type Script struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.scriptId) || (has(self.initProvider) && has(self.initProvider.scriptId))",message="spec.forProvider.scriptId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.source) || (has(self.initProvider) && has(self.initProvider.source))",message="spec.forProvider.source is a required parameter"
	Spec   ScriptSpec   `json:"spec"`
	Status ScriptStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ScriptList contains a list of Scripts
type ScriptList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Script `json:"items"`
}

// Repository type metadata.
var (
	Script_Kind             = "Script"
	Script_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Script_Kind}.String()
	Script_KindAPIVersion   = Script_Kind + "." + CRDGroupVersion.String()
	Script_GroupVersionKind = CRDGroupVersion.WithKind(Script_Kind)
)

func init() {
	SchemeBuilder.Register(&Script{}, &ScriptList{})
}