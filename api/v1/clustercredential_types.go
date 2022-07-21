package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	platformv1 "tkestack.io/tke/api/platform/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope="Cluster",shortName=cc
// +kubebuilder:storageversion

// ClusterCredential is the Schema for the clustercredentials API
type ClusterCredential struct {
	metav1.TypeMeta              `json:",inline"`
	metav1.ObjectMeta            `json:"metadata,omitempty"`
	platformv1.ClusterCredential `json:",inline"`
}

//+kubebuilder:object:root=true

// ClusterCredentialList contains a list of ClusterCredential
type ClusterCredentialList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterCredential `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ClusterCredential{}, &ClusterCredentialList{})
}
