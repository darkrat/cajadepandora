package v1beta1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// HelloType is a top-level type
type RBACDefinition struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`
	RbacBindings      []RoleBinding `json:"rbacBindings"`
}

type Subject struct {
	Kind string `json:"kind"`
	Name string `json:"name"`
}

type RoleBinding struct {
	Namespace   string `json:"namespace"`
	ClusterRole string `json:"clusterRole"`
}

type RbacBinding struct {
	Name         string        `json:"name"`
	Subjects     []Subject     `json:"subjects"`
	RoleBindings []RoleBinding `json:"roleBindings"`
}
