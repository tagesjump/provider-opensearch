// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type TemplateInitParameters struct {

	// (String) The JSON body of the index template.
	// The JSON body of the index template.
	Body *string `json:"body,omitempty" tf:"body,omitempty"`

	// (String) The name of the index template.
	// The name of the index template.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type TemplateObservation struct {

	// (String) The JSON body of the index template.
	// The JSON body of the index template.
	Body *string `json:"body,omitempty" tf:"body,omitempty"`

	// (String) The ID of this resource.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) The name of the index template.
	// The name of the index template.
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

type TemplateParameters struct {

	// (String) The JSON body of the index template.
	// The JSON body of the index template.
	// +kubebuilder:validation:Optional
	Body *string `json:"body,omitempty" tf:"body,omitempty"`

	// (String) The name of the index template.
	// The name of the index template.
	// +kubebuilder:validation:Optional
	Name *string `json:"name,omitempty" tf:"name,omitempty"`
}

// TemplateSpec defines the desired state of Template
type TemplateSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TemplateParameters `json:"forProvider"`
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
	InitProvider TemplateInitParameters `json:"initProvider,omitempty"`
}

// TemplateStatus defines the observed state of Template.
type TemplateStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TemplateObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// Template is the Schema for the Templates API. Provides an OpenSearch index template resource.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,opensearch}
type Template struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.body) || (has(self.initProvider) && has(self.initProvider.body))",message="spec.forProvider.body is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.name) || (has(self.initProvider) && has(self.initProvider.name))",message="spec.forProvider.name is a required parameter"
	Spec   TemplateSpec   `json:"spec"`
	Status TemplateStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TemplateList contains a list of Templates
type TemplateList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Template `json:"items"`
}

// Repository type metadata.
var (
	Template_Kind             = "Template"
	Template_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Template_Kind}.String()
	Template_KindAPIVersion   = Template_Kind + "." + CRDGroupVersion.String()
	Template_GroupVersionKind = CRDGroupVersion.WithKind(Template_Kind)
)

func init() {
	SchemeBuilder.Register(&Template{}, &TemplateList{})
}
