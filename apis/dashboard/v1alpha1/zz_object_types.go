// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ObjectInitParameters struct {

	// (String) The JSON body of the dashboard object.
	// The JSON body of the dashboard object.
	Body *string `json:"body,omitempty" tf:"body,omitempty"`

	// (String) The name of the index where dashboard data is stored.
	// The name of the index where dashboard data is stored.
	Index *string `json:"index,omitempty" tf:"index,omitempty"`
}

type ObjectObservation struct {

	// (String) The JSON body of the dashboard object.
	// The JSON body of the dashboard object.
	Body *string `json:"body,omitempty" tf:"body,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) The name of the index where dashboard data is stored.
	// The name of the index where dashboard data is stored.
	Index *string `json:"index,omitempty" tf:"index,omitempty"`
}

type ObjectParameters struct {

	// (String) The JSON body of the dashboard object.
	// The JSON body of the dashboard object.
	// +kubebuilder:validation:Optional
	Body *string `json:"body,omitempty" tf:"body,omitempty"`

	// (String) The name of the index where dashboard data is stored.
	// The name of the index where dashboard data is stored.
	// +kubebuilder:validation:Optional
	Index *string `json:"index,omitempty" tf:"index,omitempty"`
}

// ObjectSpec defines the desired state of Object
type ObjectSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ObjectParameters `json:"forProvider"`
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
	InitProvider ObjectInitParameters `json:"initProvider,omitempty"`
}

// ObjectStatus defines the observed state of Object.
type ObjectStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ObjectObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Object is the Schema for the Objects API. Provides an OpenSearch Dashboards object resource. This resource interacts directly with the underlying OpenSearch index backing Dashboards, so the format must match what Dashboards the version of Dashboards is expecting. Dashboards with older versions - directly pulling the JSON from a Dashboards index of the same version of OpenSearch targeted by the provider is a workaround.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,opensearch}
type Object struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.body) || (has(self.initProvider) && has(self.initProvider.body))",message="spec.forProvider.body is a required parameter"
	Spec   ObjectSpec   `json:"spec"`
	Status ObjectStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ObjectList contains a list of Objects
type ObjectList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Object `json:"items"`
}

// Repository type metadata.
var (
	Object_Kind             = "Object"
	Object_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Object_Kind}.String()
	Object_KindAPIVersion   = Object_Kind + "." + CRDGroupVersion.String()
	Object_GroupVersionKind = CRDGroupVersion.WithKind(Object_Kind)
)

func init() {
	SchemeBuilder.Register(&Object{}, &ObjectList{})
}
