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

type MappingInitParameters struct {

	// (Set of String) A list of backend roles.
	// A list of backend roles.
	// +listType=set
	AndBackendRoles []*string `json:"andBackendRoles,omitempty" tf:"and_backend_roles,omitempty"`

	// (Set of String) A list of backend roles.
	// A list of backend roles.
	// +listType=set
	BackendRoles []*string `json:"backendRoles,omitempty" tf:"backend_roles,omitempty"`

	// (String) Description of the role mapping.
	// Description of the role mapping.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (Set of String) A list of host names.
	// A list of host names.
	// +listType=set
	Hosts []*string `json:"hosts,omitempty" tf:"hosts,omitempty"`

	// (String) The name of the security role.
	// The name of the security role.
	// +crossplane:generate:reference:type=github.com/tagesjump/provider-opensearch/apis/opensearch/v1alpha1.Role
	RoleName *string `json:"roleName,omitempty" tf:"role_name,omitempty"`

	// Reference to a Role in opensearch to populate roleName.
	// +kubebuilder:validation:Optional
	RoleNameRef *v1.Reference `json:"roleNameRef,omitempty" tf:"-"`

	// Selector for a Role in opensearch to populate roleName.
	// +kubebuilder:validation:Optional
	RoleNameSelector *v1.Selector `json:"roleNameSelector,omitempty" tf:"-"`

	// (Set of String) A list of users.
	// A list of users.
	// +crossplane:generate:reference:type=github.com/tagesjump/provider-opensearch/apis/opensearch/v1alpha1.User
	// +listType=set
	Users []*string `json:"users,omitempty" tf:"users,omitempty"`

	// References to User in opensearch to populate users.
	// +kubebuilder:validation:Optional
	UsersRefs []v1.Reference `json:"usersRefs,omitempty" tf:"-"`

	// Selector for a list of User in opensearch to populate users.
	// +kubebuilder:validation:Optional
	UsersSelector *v1.Selector `json:"usersSelector,omitempty" tf:"-"`
}

type MappingObservation struct {

	// (Set of String) A list of backend roles.
	// A list of backend roles.
	// +listType=set
	AndBackendRoles []*string `json:"andBackendRoles,omitempty" tf:"and_backend_roles,omitempty"`

	// (Set of String) A list of backend roles.
	// A list of backend roles.
	// +listType=set
	BackendRoles []*string `json:"backendRoles,omitempty" tf:"backend_roles,omitempty"`

	// (String) Description of the role mapping.
	// Description of the role mapping.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (Set of String) A list of host names.
	// A list of host names.
	// +listType=set
	Hosts []*string `json:"hosts,omitempty" tf:"hosts,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) The name of the security role.
	// The name of the security role.
	RoleName *string `json:"roleName,omitempty" tf:"role_name,omitempty"`

	// (Set of String) A list of users.
	// A list of users.
	// +listType=set
	Users []*string `json:"users,omitempty" tf:"users,omitempty"`
}

type MappingParameters struct {

	// (Set of String) A list of backend roles.
	// A list of backend roles.
	// +kubebuilder:validation:Optional
	// +listType=set
	AndBackendRoles []*string `json:"andBackendRoles,omitempty" tf:"and_backend_roles,omitempty"`

	// (Set of String) A list of backend roles.
	// A list of backend roles.
	// +kubebuilder:validation:Optional
	// +listType=set
	BackendRoles []*string `json:"backendRoles,omitempty" tf:"backend_roles,omitempty"`

	// (String) Description of the role mapping.
	// Description of the role mapping.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (Set of String) A list of host names.
	// A list of host names.
	// +kubebuilder:validation:Optional
	// +listType=set
	Hosts []*string `json:"hosts,omitempty" tf:"hosts,omitempty"`

	// (String) The name of the security role.
	// The name of the security role.
	// +crossplane:generate:reference:type=github.com/tagesjump/provider-opensearch/apis/opensearch/v1alpha1.Role
	// +kubebuilder:validation:Optional
	RoleName *string `json:"roleName,omitempty" tf:"role_name,omitempty"`

	// Reference to a Role in opensearch to populate roleName.
	// +kubebuilder:validation:Optional
	RoleNameRef *v1.Reference `json:"roleNameRef,omitempty" tf:"-"`

	// Selector for a Role in opensearch to populate roleName.
	// +kubebuilder:validation:Optional
	RoleNameSelector *v1.Selector `json:"roleNameSelector,omitempty" tf:"-"`

	// (Set of String) A list of users.
	// A list of users.
	// +crossplane:generate:reference:type=github.com/tagesjump/provider-opensearch/apis/opensearch/v1alpha1.User
	// +kubebuilder:validation:Optional
	// +listType=set
	Users []*string `json:"users,omitempty" tf:"users,omitempty"`

	// References to User in opensearch to populate users.
	// +kubebuilder:validation:Optional
	UsersRefs []v1.Reference `json:"usersRefs,omitempty" tf:"-"`

	// Selector for a list of User in opensearch to populate users.
	// +kubebuilder:validation:Optional
	UsersSelector *v1.Selector `json:"usersSelector,omitempty" tf:"-"`
}

// MappingSpec defines the desired state of Mapping
type MappingSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     MappingParameters `json:"forProvider"`
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
	InitProvider MappingInitParameters `json:"initProvider,omitempty"`
}

// MappingStatus defines the observed state of Mapping.
type MappingStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        MappingObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Mapping is the Schema for the Mappings API. Provides an OpenSearch security role mapping. Please refer to the OpenSearch Access Control documentation for details.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,opensearch}
type Mapping struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MappingSpec   `json:"spec"`
	Status            MappingStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MappingList contains a list of Mappings
type MappingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Mapping `json:"items"`
}

// Repository type metadata.
var (
	Mapping_Kind             = "Mapping"
	Mapping_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Mapping_Kind}.String()
	Mapping_KindAPIVersion   = Mapping_Kind + "." + CRDGroupVersion.String()
	Mapping_GroupVersionKind = CRDGroupVersion.WithKind(Mapping_Kind)
)

func init() {
	SchemeBuilder.Register(&Mapping{}, &MappingList{})
}
