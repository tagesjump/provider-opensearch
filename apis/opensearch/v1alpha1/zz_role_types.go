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

type IndexPermissionsInitParameters struct {

	// (Set of String) A list of allowed actions.
	// A list of allowed actions.
	// +listType=set
	AllowedActions []*string `json:"allowedActions,omitempty" tf:"allowed_actions,omitempty"`

	// level security (json formatted using jsonencode).
	// A selector for document-level security (json formatted using jsonencode).
	DocumentLevelSecurity *string `json:"documentLevelSecurity,omitempty" tf:"document_level_security,omitempty"`

	// level security.
	// A list of selectors for field-level security.
	// +listType=set
	FieldLevelSecurity []*string `json:"fieldLevelSecurity,omitempty" tf:"field_level_security,omitempty"`

	// (Set of String) A list of glob patterns for the index names.
	// A list of glob patterns for the index names.
	// +listType=set
	IndexPatterns []*string `json:"indexPatterns,omitempty" tf:"index_patterns,omitempty"`

	// (Set of String) A list of masked fields
	// A list of masked fields
	// +listType=set
	MaskedFields []*string `json:"maskedFields,omitempty" tf:"masked_fields,omitempty"`
}

type IndexPermissionsObservation struct {

	// (Set of String) A list of allowed actions.
	// A list of allowed actions.
	// +listType=set
	AllowedActions []*string `json:"allowedActions,omitempty" tf:"allowed_actions,omitempty"`

	// level security (json formatted using jsonencode).
	// A selector for document-level security (json formatted using jsonencode).
	DocumentLevelSecurity *string `json:"documentLevelSecurity,omitempty" tf:"document_level_security,omitempty"`

	// level security.
	// A list of selectors for field-level security.
	// +listType=set
	FieldLevelSecurity []*string `json:"fieldLevelSecurity,omitempty" tf:"field_level_security,omitempty"`

	// (Set of String) A list of glob patterns for the index names.
	// A list of glob patterns for the index names.
	// +listType=set
	IndexPatterns []*string `json:"indexPatterns,omitempty" tf:"index_patterns,omitempty"`

	// (Set of String) A list of masked fields
	// A list of masked fields
	// +listType=set
	MaskedFields []*string `json:"maskedFields,omitempty" tf:"masked_fields,omitempty"`
}

type IndexPermissionsParameters struct {

	// (Set of String) A list of allowed actions.
	// A list of allowed actions.
	// +kubebuilder:validation:Optional
	// +listType=set
	AllowedActions []*string `json:"allowedActions,omitempty" tf:"allowed_actions,omitempty"`

	// level security (json formatted using jsonencode).
	// A selector for document-level security (json formatted using jsonencode).
	// +kubebuilder:validation:Optional
	DocumentLevelSecurity *string `json:"documentLevelSecurity,omitempty" tf:"document_level_security,omitempty"`

	// level security.
	// A list of selectors for field-level security.
	// +kubebuilder:validation:Optional
	// +listType=set
	FieldLevelSecurity []*string `json:"fieldLevelSecurity,omitempty" tf:"field_level_security,omitempty"`

	// (Set of String) A list of glob patterns for the index names.
	// A list of glob patterns for the index names.
	// +kubebuilder:validation:Optional
	// +listType=set
	IndexPatterns []*string `json:"indexPatterns,omitempty" tf:"index_patterns,omitempty"`

	// (Set of String) A list of masked fields
	// A list of masked fields
	// +kubebuilder:validation:Optional
	// +listType=set
	MaskedFields []*string `json:"maskedFields,omitempty" tf:"masked_fields,omitempty"`
}

type RoleInitParameters struct {

	// (Set of String) A list of cluster permissions.
	// A list of cluster permissions.
	// +listType=set
	ClusterPermissions []*string `json:"clusterPermissions,omitempty" tf:"cluster_permissions,omitempty"`

	// (String) Description of the role.
	// Description of the role.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (Block Set) A configuration of index permissions (see below for nested schema)
	// A configuration of index permissions
	IndexPermissions []IndexPermissionsInitParameters `json:"indexPermissions,omitempty" tf:"index_permissions,omitempty"`

	// (String) The name of the security role.
	// The name of the security role.
	RoleName *string `json:"roleName,omitempty" tf:"role_name,omitempty"`

	// (Block Set) A configuration of tenant permissions (see below for nested schema)
	// A configuration of tenant permissions
	TenantPermissions []TenantPermissionsInitParameters `json:"tenantPermissions,omitempty" tf:"tenant_permissions,omitempty"`
}

type RoleObservation struct {

	// (Set of String) A list of cluster permissions.
	// A list of cluster permissions.
	// +listType=set
	ClusterPermissions []*string `json:"clusterPermissions,omitempty" tf:"cluster_permissions,omitempty"`

	// (String) Description of the role.
	// Description of the role.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Block Set) A configuration of index permissions (see below for nested schema)
	// A configuration of index permissions
	IndexPermissions []IndexPermissionsObservation `json:"indexPermissions,omitempty" tf:"index_permissions,omitempty"`

	// (String) The name of the security role.
	// The name of the security role.
	RoleName *string `json:"roleName,omitempty" tf:"role_name,omitempty"`

	// (Block Set) A configuration of tenant permissions (see below for nested schema)
	// A configuration of tenant permissions
	TenantPermissions []TenantPermissionsObservation `json:"tenantPermissions,omitempty" tf:"tenant_permissions,omitempty"`
}

type RoleParameters struct {

	// (Set of String) A list of cluster permissions.
	// A list of cluster permissions.
	// +kubebuilder:validation:Optional
	// +listType=set
	ClusterPermissions []*string `json:"clusterPermissions,omitempty" tf:"cluster_permissions,omitempty"`

	// (String) Description of the role.
	// Description of the role.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (Block Set) A configuration of index permissions (see below for nested schema)
	// A configuration of index permissions
	// +kubebuilder:validation:Optional
	IndexPermissions []IndexPermissionsParameters `json:"indexPermissions,omitempty" tf:"index_permissions,omitempty"`

	// (String) The name of the security role.
	// The name of the security role.
	// +kubebuilder:validation:Optional
	RoleName *string `json:"roleName,omitempty" tf:"role_name,omitempty"`

	// (Block Set) A configuration of tenant permissions (see below for nested schema)
	// A configuration of tenant permissions
	// +kubebuilder:validation:Optional
	TenantPermissions []TenantPermissionsParameters `json:"tenantPermissions,omitempty" tf:"tenant_permissions,omitempty"`
}

type TenantPermissionsInitParameters struct {

	// (Set of String) A list of allowed actions.
	// A list of allowed actions.
	// +listType=set
	AllowedActions []*string `json:"allowedActions,omitempty" tf:"allowed_actions,omitempty"`

	// (Set of String) A list of glob patterns for the tenant names
	// A list of glob patterns for the tenant names
	// +listType=set
	TenantPatterns []*string `json:"tenantPatterns,omitempty" tf:"tenant_patterns,omitempty"`
}

type TenantPermissionsObservation struct {

	// (Set of String) A list of allowed actions.
	// A list of allowed actions.
	// +listType=set
	AllowedActions []*string `json:"allowedActions,omitempty" tf:"allowed_actions,omitempty"`

	// (Set of String) A list of glob patterns for the tenant names
	// A list of glob patterns for the tenant names
	// +listType=set
	TenantPatterns []*string `json:"tenantPatterns,omitempty" tf:"tenant_patterns,omitempty"`
}

type TenantPermissionsParameters struct {

	// (Set of String) A list of allowed actions.
	// A list of allowed actions.
	// +kubebuilder:validation:Optional
	// +listType=set
	AllowedActions []*string `json:"allowedActions,omitempty" tf:"allowed_actions,omitempty"`

	// (Set of String) A list of glob patterns for the tenant names
	// A list of glob patterns for the tenant names
	// +kubebuilder:validation:Optional
	// +listType=set
	TenantPatterns []*string `json:"tenantPatterns,omitempty" tf:"tenant_patterns,omitempty"`
}

// RoleSpec defines the desired state of Role
type RoleSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RoleParameters `json:"forProvider"`
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
	InitProvider RoleInitParameters `json:"initProvider,omitempty"`
}

// RoleStatus defines the observed state of Role.
type RoleStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RoleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Role is the Schema for the Roles API. Provides an OpenSearch security role resource. Please refer to the OpenSearch Access Control documentation for details.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,opensearch}
type Role struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.roleName) || (has(self.initProvider) && has(self.initProvider.roleName))",message="spec.forProvider.roleName is a required parameter"
	Spec   RoleSpec   `json:"spec"`
	Status RoleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RoleList contains a list of Roles
type RoleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Role `json:"items"`
}

// Repository type metadata.
var (
	Role_Kind             = "Role"
	Role_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Role_Kind}.String()
	Role_KindAPIVersion   = Role_Kind + "." + CRDGroupVersion.String()
	Role_GroupVersionKind = CRDGroupVersion.WithKind(Role_Kind)
)

func init() {
	SchemeBuilder.Register(&Role{}, &RoleList{})
}