// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type UserInitParameters struct {

	// (Map of String) A map of arbitrary key value string pairs stored alongside of users.
	// A map of arbitrary key value string pairs stored alongside of users.
	// +mapType=granular
	Attributes map[string]*string `json:"attributes,omitempty" tf:"attributes,omitempty"`

	// (Set of String) A list of backend roles.
	// A list of backend roles.
	// +listType=set
	BackendRoles []*string `json:"backendRoles,omitempty" tf:"backend_roles,omitempty"`

	// (String) Description of the user.
	// Description of the user.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) The name of the security user.
	// The name of the security user.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type UserObservation struct {

	// (Map of String) A map of arbitrary key value string pairs stored alongside of users.
	// A map of arbitrary key value string pairs stored alongside of users.
	// +mapType=granular
	Attributes map[string]*string `json:"attributes,omitempty" tf:"attributes,omitempty"`

	// (Set of String) A list of backend roles.
	// A list of backend roles.
	// +listType=set
	BackendRoles []*string `json:"backendRoles,omitempty" tf:"backend_roles,omitempty"`

	// (String) Description of the user.
	// Description of the user.
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) The name of the security user.
	// The name of the security user.
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

type UserParameters struct {

	// (Map of String) A map of arbitrary key value string pairs stored alongside of users.
	// A map of arbitrary key value string pairs stored alongside of users.
	// +kubebuilder:validation:Optional
	// +mapType=granular
	Attributes map[string]*string `json:"attributes,omitempty" tf:"attributes,omitempty"`

	// (Set of String) A list of backend roles.
	// A list of backend roles.
	// +kubebuilder:validation:Optional
	// +listType=set
	BackendRoles []*string `json:"backendRoles,omitempty" tf:"backend_roles,omitempty"`

	// (String) Description of the user.
	// Description of the user.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// hashed password for the user, cannot be specified with password.
	// The pre-hashed password for the user, cannot be specified with `password`.
	// +kubebuilder:validation:Optional
	PasswordHashSecretRef *v1.SecretKeySelector `json:"passwordHashSecretRef,omitempty" tf:"-"`

	// descriptive HTTP 400 Bad Request error. For AWS OpenSearch domains "password must be at least 8 characters long and contain at least one uppercase letter, one lowercase letter, one digit, and one special character".
	// The plain text password for the user, cannot be specified with `password_hash`. Some implementations may enforce a password policy. Invalid passwords may cause a non-descriptive HTTP 400 Bad Request error. For AWS OpenSearch domains "password must be at least 8 characters long and contain at least one uppercase letter, one lowercase letter, one digit, and one special character".
	// +kubebuilder:validation:Optional
	PasswordSecretRef *v1.SecretKeySelector `json:"passwordSecretRef,omitempty" tf:"-"`

	// (String) The name of the security user.
	// The name of the security user.
	// +kubebuilder:validation:Optional
	Username *string `json:"username,omitempty" tf:"username,omitempty"`
}

// UserSpec defines the desired state of User
type UserSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     UserParameters `json:"forProvider"`
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
	InitProvider UserInitParameters `json:"initProvider,omitempty"`
}

// UserStatus defines the observed state of User.
type UserStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        UserObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// User is the Schema for the Users API. Provides an OpenSearch security user. Please refer to the OpenSearch Access Control documentation for details.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,opensearch}
type User struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.username) || (has(self.initProvider) && has(self.initProvider.username))",message="spec.forProvider.username is a required parameter"
	Spec   UserSpec   `json:"spec"`
	Status UserStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// UserList contains a list of Users
type UserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []User `json:"items"`
}

// Repository type metadata.
var (
	User_Kind             = "User"
	User_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: User_Kind}.String()
	User_KindAPIVersion   = User_Kind + "." + CRDGroupVersion.String()
	User_GroupVersionKind = CRDGroupVersion.WithKind(User_Kind)
)

func init() {
	SchemeBuilder.Register(&User{}, &UserList{})
}
